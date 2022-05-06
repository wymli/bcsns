package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/pkg/errors"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/utils"
	"google.golang.org/grpc"
)

type User struct {
	Uid       string `json:"uid,omitempty"`
	Phone     int    `json:"phone,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Sex       int    `json:"sex,omitempty"`
	Age       int    `json:"age,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
	Address   string `json:"address,omitempty"`
	PublicKey string `json:"public_key,omitempty"`
	Password  string `json:"password,omitempty"`
	Follows   []User `json:"follows,omitempty"`
}

func (u *User) Store(client *dgo.Dgraph) (uint64, error) {
	userJson, err := json.Marshal(u)
	if err != nil {
		return 0, errors.Wrapf(errx.ERROR_SERVER_COMMON, "failed to marshall json, err:%v", err)
	}

	qFormat := `
  query {
      user as var(func: eq(phone, "%d"))
  }`
	query := fmt.Sprintf(qFormat, u.Phone)

	mu := &api.Mutation{
		Cond: `@if(eq(len(user), 0))`,
		// SetNquads: []byte(`uid(user) <email> "correct_email@dgraph.io" .`),
		SetJson: userJson,
	}

	apiReq := &api.Request{
		Query:     query,
		Mutations: []*api.Mutation{mu},
		CommitNow: true,
	}

	res, err := client.NewTxn().Do(context.Background(), apiReq)
	if err != nil {
		return 0, errx.Wrapf(errx.ERROR_DB, "failed to insert user in dgraph, err:%v", err)
	}

	if len(res.Uids) == 0 {
		return 0, errx.Wrapf(errx.ERROR_USER_DUPLICATE, "duplicate register of user phone: %v", u.Phone)
	}

	_, v, ok := utils.ExtractOneFromStringMap(res.Uids)
	if !ok {
		return 0, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to extract uid from dgraph response, want 1 uid, get %d uid", len(res.Uids))
	}

	// strconv: base=0 则根据前缀判断
	uid, err := strconv.ParseUint(v, 0, 64)
	if err != nil {
		return 0, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to convert user uid from string to uint64: %v, err: %v", uid, err)
	}

	return uid, nil
}

type Group struct {
	Uid         string `json:"uid,omitempty"`
	GroupName   string `json:"group_name,omitempty"`
	GroupAvatar string `json:"group_avatar,omitempty"`
	Members     []User `json:"members,omitempty"`
}

func (g *Group) Store(client *dgo.Dgraph) (uint64, error) {
	groupJson, err := json.Marshal(g)
	fmt.Println(string(groupJson))
	if err != nil {
		return 0, errors.Wrapf(errx.ERROR_SERVER_COMMON, "failed to marshall json, err:%v", err)
	}

	m := api.Mutation{
		SetJson: groupJson,
	}

	apiReq := api.Request{
		Mutations: []*api.Mutation{&m},
		CommitNow: true,
	}

	res, err := client.NewTxn().Do(context.Background(), &apiReq)
	if err != nil {
		return 0, errx.Wrapf(errx.ERROR_DB, "failed to insert group in dgraph, err:%v", err)
	}

	if len(res.Uids) == 0 {
		return 0, errx.Wrapf(errx.ERROR_UNKNOWN, "len(res.uids) == 0 when inserting group ")
	}

	_, v, ok := utils.ExtractOneFromStringMap(res.Uids)
	if !ok {
		return 0, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to extract uid from dgraph response, want 1 uid, get %d uid", len(res.Uids))
	}

	uid, err := strconv.ParseUint(v, 0, 64)
	if err != nil {
		return 0, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to convert group uid from string to uint64: %v, err: %v", uid, err)
	}

	return uid, nil
}

var quser = `query all($uid: string) {
	all(func: uid($uid)) {
		uid
		nickname
		sex
		age
		phone
		avatar
		address
		public_key
		follows{
			uid
		}
		ingroups{
			uid
		}
	}
}`

var qfollows = `query all($uid: string) {
	all(func: uid($uid)) {
		uid
		follows {
			uid
		}
	}
}`

func QueryUser(client *dgo.Dgraph, uid uint64) (User, error) {
	res, err := client.NewTxn().QueryWithVars(context.Background(), quser, map[string]string{"$uid": "0x" + strconv.FormatUint(uid, 16)})
	if err != nil {
		return User{}, errx.Wrapf(errx.ERROR_DB, "failed to query uid:%v profile from dgraph", uid)
	}

	fmt.Println("stest:", string(res.Json))

	var decode struct {
		All []User
	}

	err = json.Unmarshal(res.Json, &decode)
	if err != nil {
		return User{}, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to unmarshall json: %v, err:%v", string(res.Json), err)
	}

	switch {
	case len(decode.All) == 0:
		return User{}, errx.Wrapf(errx.ERROR_USER_NOT_FOUND, "no user found of uid:%v", uid)
	case len(decode.All) > 1:
		return User{}, errx.Wrapf(errx.ERROR_USER_DUPLICATE, "duplicate user found of uid:%v", uid)
	}

	user := decode.All[0]
	return user, nil
}

var qgroup = `query all($uid: string) {
	all(func: uid($uid)) {
		uid
		group_name
		group_avatar
		members{
			uid
		}
	}
}`

func QueryGroup(client *dgo.Dgraph, uid uint64) (Group, error) {
	res, err := client.NewTxn().QueryWithVars(context.Background(), qgroup, map[string]string{"$uid": "0x" + strconv.FormatUint(uid, 16)})
	if err != nil {
		return Group{}, errx.Wrapf(errx.ERROR_DB, "failed to query uid:%v profile from dgraph", uid)
	}

	var decode struct {
		All []Group
	}
	fmt.Println(res)
	fmt.Println(string(res.Json))

	err = json.Unmarshal(res.Json, &decode)
	if err != nil {
		return Group{}, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to unmarshall json: %v, err:%v", string(res.Json), err)
	}

	switch {
	case len(decode.All) == 0:
		return Group{}, errx.Wrapf(errx.ERROR_USER_NOT_FOUND, "no group found of uid:%v", uid)
	case len(decode.All) > 1:
		return Group{}, errx.Wrapf(errx.ERROR_USER_DUPLICATE, "duplicate group found of uid:%v", uid)
	}

	group := decode.All[0]
	fmt.Printf("group decode: %v\n", decode)
	return group, nil
}

func AddFriend(client *dgo.Dgraph, from uint64, to uint64) error {
	f := "0x" + strconv.FormatUint(from, 16)
	t := "0x" + strconv.FormatUint(to, 16)
	stmt := fmt.Sprintf("<%s> <follows> <%s> .", f, t)
	fmt.Println(from, f, to, t)

	m := api.Mutation{
		SetNquads: []byte(stmt),
	}

	apiReq := api.Request{
		Mutations: []*api.Mutation{&m},
		CommitNow: true,
	}

	res, err := client.NewTxn().Do(context.Background(), &apiReq)
	if err != nil {
		return errx.Wrapf(errx.ERROR_DB, "failed to follow in dgraph, err:%v", err)
	}

	fmt.Println("addfriend:", res.Json)
	return nil
}

// ok
var qfriend = `query all($uid: string) {
	G as var(func: uid($uid)) {
		F as follows
	}

	all(func: uid(G)) @normalize {
		friends : ~follows @filter(uid(F)) {
			uid: uid
			nickname: nickname
			sex: sex
			age: age
			phone: phone
			avatar: avatar
			address: address
			public_key: public_key
		}
	}
}`

func QueryFriends(client *dgo.Dgraph, uid uint64) ([]User, error) {
	res, err := client.NewTxn().QueryWithVars(context.Background(), qfriend, map[string]string{"$uid": "0x" + strconv.FormatUint(uid, 16)})
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_DB, "failed to query uid:%v friends from dgraph", uid)
	}
	fmt.Println("f:", string(res.Json))

	var decode struct {
		All []User `json:"all,omitempty"`
	}

	fmt.Println(string(res.Json))

	err = json.Unmarshal(res.Json, &decode)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to unmarshall json: %v, err:%v", string(res.Json), err)
	}

	return decode.All, nil
}

// ok
var qmember = `query all($uid: string) {
	all(func: uid($uid))@normalize  {
		~ingroups{
			uid: uid
			nickname: nickname
			sex: sex
			age: age
			phone: phone
			avatar: avatar
			address: address
			public_key: public_key
		}
	}
}`

func QueryMembers(client *dgo.Dgraph, uid uint64) ([]User, error) {
	res, err := client.NewTxn().QueryWithVars(context.Background(), qmember, map[string]string{"$uid": "0x" + strconv.FormatUint(uid, 16)})
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_DB, "failed to query uid:%v friends from dgraph", uid)
	}

	fmt.Println(string(res.Json))

	var decode struct {
		All []User
	}

	err = json.Unmarshal(res.Json, &decode)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to unmarshall json: %v, err:%v", string(res.Json), err)
	}

	return decode.All, nil
}

func testNewExpr(client *dgo.Dgraph, uid uint64) ([]User, error) {
	qfmt := `
	{  
		G as var(func: uid("%s")) {
			F as follows
			F2 as ~follows 
	 	}

	 	result(func: uid(G)) {
			uid
			requestsLeft : follows @filter(NOT uid(F2)) {
				uid
			}
			accepted : ~follows @filter(uid(F)) {
				uid
			}
			allFriends : follows {
				uid
			}
		}
	}
	`
	q := fmt.Sprintf(qfmt, "0x"+strconv.FormatUint(uid, 16))

	// res, err := client.NewTxn().QueryWithVars(context.Background(), q, map[string]string{"$uid": "0x" + strconv.FormatUint(uid, 16)})
	res, err := client.NewTxn().Query(context.Background(), q)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_DB, "failed to query uid:%v friends from dgraph", uid)
	}

	fmt.Println(string(res.Json))

	var decode struct {
		All []User
	}

	err = json.Unmarshal(res.Json, &decode)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to unmarshall json: %v, err:%v", string(res.Json), err)
	}

	return decode.All, nil
}

func JoinGroup(client *dgo.Dgraph, groupId uint64, userId uint64) error {
	g, u := utils.StringUid(groupId), utils.StringUid(userId)
	stmt := fmt.Sprintf("<%s> <members> <%s> .", g, u)

	qFormat := `
  query {
		group as var(func: uid("%s")) 
		user as var(func: uid("%s")) 
	}`
	query := fmt.Sprintf(qFormat, g, u)

	m := api.Mutation{
		Cond:      `@if(eq(len(user), 1) AND eq(len(group), 1))`,
		SetNquads: []byte(stmt),
	}

	apiReq := api.Request{
		Query:     query,
		Mutations: []*api.Mutation{&m},
		CommitNow: true,
	}

	res, err := client.NewTxn().Do(context.Background(), &apiReq)
	if err != nil {
		return errx.Wrapf(errx.ERROR_DB, "failed to join group in dgraph, err:%v", err)
	}

	if res.Metrics.NumUids["mutation_cost"] == 0 {
		return errx.Wrapf(errx.ERROR_BAD_REQUEST, "用户或群组不存在")
	}

	return nil
}

func main() {
	endpoint := "localhost:9080"
	d, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)

	// testNewExpr(client, 820)

	// u, err := QueryFriends(client, 820)
	// for _, uu := range u {
	// 	fmt.Println("friend:", uu.Uid)
	// }

	// m, err := QueryMembers(client, 291)
	// for _, uu := range m {
	// 	fmt.Println("member:", uu.Uid)
	// }
	// return

	user := User{
		Phone:     1032203,
		Nickname:  "apple",
		Sex:       1,
		Age:       20,
		Avatar:    "asdf",
		Address:   "asdf",
		PublicKey: "asdf",
		Password:  "asasdfasdfdf",
	}
	userId, err := user.Store(client)
	if err != nil {
		panic(err)
	}
	fmt.Println("userid:", userId)

	group := Group{
		GroupName:   "group_1",
		GroupAvatar: "asdf",
		Members: []User{
			{Uid: "0x" + strconv.FormatUint(userId, 16)},
		},
	}
	groupId, err := group.Store(client)
	if err != nil {
		panic(err)
	}
	fmt.Println("groupid:", groupId)

	userx, err := QueryUser(client, userId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", userx)

	groupx, err := QueryGroup(client, groupId)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v\n", groupx)

	var decode struct {
		All []Group
	}
	err = Query(client, qgroup, &decode, groupId)
	if err != nil {
		panic(err)
	}
	fmt.Println(decode)

	// fmt.Println()
	// err = JoinGroup(client, groupId, userId)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println()
	// err = JoinGroup(client, 444, userId)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println()
	// err = JoinGroup(client, groupId, 1233)
	// if err != nil {
	// 	panic(err)
	// }

	// err = AddFriend(client, userId, 12)
	// if err != nil {
	// 	panic(err)
	// }

	// userx, err = QueryUser(client, userId)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%#v\n", userx)
}

func Query(client *dgo.Dgraph, qFormat string, decode interface{}, uid uint64) error {
	res, err := client.NewTxn().QueryWithVars(context.Background(), qFormat, map[string]string{"$uid": utils.StringUid(uid)})
	if err != nil {
		return errx.Wrapf(errx.ERROR_DB, "failed to query uid:%v from dgraph", uid)
	}

	err = json.Unmarshal(res.Json, decode)
	if err != nil {
		return errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to unmarshall json: %v, err:%v", string(res.Json), err)
	}

	return nil
}
