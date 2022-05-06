package model

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
		return 0, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to insert group in dgraph, json:%v", string(res.Json))
	}

	_, v, ok := utils.ExtractOneFromStringMap(res.Uids)
	if !ok {
		return 0, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to extract uid from dgraph response, want 1 uid, get %d uid, json:%v", len(res.Uids), string(res.Json))
	}

	return utils.UintUid(v), nil
}
