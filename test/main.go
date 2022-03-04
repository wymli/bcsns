package main

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"sync"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/wymli/bcsns/app/seq_service/seq"
	"github.com/wymli/bcsns/common/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type User struct {
	Phone     int    `json:"phone"`
	Nickname  string `json:"nickname"`
	Sex       int    `json:"sex"`
	Age       int    `json:"age"`
	Avatar    string `json:"avatar"`
	Address   string `json:"address"`
	PublicKey string `json:"public_key"`
	Password  string `json:"password"`
}

var (
	testPhone = 11122133
	testPwd   = "123456789"
)

func main() {
	// uid := condMutate()
	// query(uid)
	// checkpwd(uid)
	c := seq.NewSeq(zrpc.MustNewClient(zrpc.RpcClientConf{
		Target: "localhost:9009",
	}))
	done := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		done.Add(1)
		ii := i
		go func() {
			defer done.Done()
			l := []uint64{}
			for j := 0; j < 1; j++ {
				rsp, err := c.GetSeqId(context.TODO(), &seq.GetSeqIdReq{
					UserId: uint64(ii),
				})
				if err != nil {
					panic(err)
				}
				l = append(l, rsp.SeqId)
				fmt.Println(ii, rsp, err)
			}
			for x := range l[:len(l)-1] {
				if l[x] > l[x+1] {
					panic(fmt.Sprintf("bad:%v", l))
				}
			}
		}()
	}
	done.Wait()
}

func checkpwd(uid string) {
	qLogin := `
query check($phone: int, $password: string){
  check(func: eq(phone, $phone)) {
		uid
    pass: checkpwd(password, $password)
  }
}
`
	c := newDgraphClient("localhost:9080")

	fmt.Printf("check uid %v\n", uid)

	txn := c.NewTxn()
	res, err := txn.QueryWithVars(context.Background(), qLogin,
		map[string]string{"$phone": strconv.Itoa(testPhone), "$password": testPwd})
	if err != nil {
		fmt.Println(err)
	}
	var decode struct {
		Check []struct {
			Uid  string
			Pass bool
		}
	}
	fmt.Printf("%s\n", string(res.Json))

	err = json.Unmarshal(res.Json, &decode)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", decode)
}

func query(uid string) {
	q := `query all($a: string) {
    all(func: uid($a)) {
			uid
      phone
			nickname
			sex
			age
			avatar
			address
			public_key
			password
    }
  }`

	c := newDgraphClient("localhost:9080")

	txn := c.NewTxn()
	fmt.Println("querying uid ", uid)
	res, err := txn.QueryWithVars(context.Background(), q, map[string]string{"$a": uid})
	if err != nil {
		fmt.Println(err)
	}
	var decode struct {
		All []User
	}
	fmt.Printf("%s\n", string(res.Json))

	err = json.Unmarshal(res.Json, &decode)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v\n", decode)
	// {"all":[{"uid":"0x1a","phone":1,"nickname":"apple","sex":1,"age":12,"avatar":"s","address":"asdfsadf","public_key":"asdfasdfs"}]}
	// struct { All []main.User }{All:[]main.User{main.User{Phone:1, Nickname:"apple", Sex:1, Age:12, Avatar:"s", Address:"asdfsadf", PublicKey:"asdfasdfs", Password:""}}}
}

func condMutate() string {
	user := User{
		Phone:     testPhone,
		Nickname:  "apx",
		Sex:       1,
		Age:       12,
		Avatar:    "s",
		Address:   "asdfsadf",
		PublicKey: "asdfasdfs",
		Password:  testPwd,
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("failed to marshall json, err:%v", err)
		return ""
	}

	qFormat := `
  query {
      user as var(func: eq(phone, "%d"))
  }`
	query := fmt.Sprintf(qFormat, user.Phone)
	mu := &api.Mutation{
		Cond: `@if(eq(len(user), 0))`, // Only mutate if "wrong_email@dgraph.io" belongs to single user.
		// SetNquads: []byte(`uid(user) <email> "correct_email@dgraph.io" .`),
		SetJson: userJson,
	}
	apiReq := &api.Request{
		Query:     query,
		Mutations: []*api.Mutation{mu},
		CommitNow: true,
	}

	c := newDgraphClient("localhost:9080")

	// Update email only if exactly one matching uid is found.
	res, err := c.NewTxn().Do(context.Background(), apiReq)
	if err != nil {
		// failed to do txn, err:rpc error: code = Unknown desc = Password too short, i.e. should have at least 6 chars<nil>
		fmt.Printf("failed to do txn, err:%v", err)
	}
	fmt.Println(res)
	// get uid
	// 如果插入成功.会返回uid,否则uid是空
	var uid string
	for k, v := range res.Uids {
		fmt.Println("uid_k:", k, "  uid_v:", v)
		uid = v
	}
	// json:"{}" txn:<start_ts:866 commit_ts:867 preds:"1-0-address" preds:"1-0-age" preds:"1-0-avatar" preds:"1-0-nickname" preds:"1-0-password" preds:"1-0-phone" preds:"1-0-public_key" preds:"1-0-sex" > latency:<parsing_ns:1313800 processing_ns:78688900 encoding_ns:100600 assign_timestamp_ns:834600 total_ns:81114700 > metrics:<num_uids:<key:"" value:1 > num_uids:<key:"_total" value:11 > num_uids:<key:"mutation_cost" value:9 > num_uids:<key:"phone" value:0 > num_uids:<key:"user" value:1 > > uids:<key:"dg.3488257281.20" value:"0x1a" > hdrs:<key:"content-type" value:<value:"application/grpc" > > hdrs:<key:"dgraph-toucheduids" value:<value:"11" > >
	return uid
}

func newDgraphClient(endpoint string) *dgo.Dgraph {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial(endpoint, grpc.WithInsecure())
	logx.FatalIfErrf(err, "failed to dial dgraph endpoint:%s", endpoint)

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	)
}
