package model

import (
	"context"
	"fmt"

	"github.com/dgraph-io/dgo/v210"
	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/utils"
)

func Follow(client *dgo.Dgraph, from uint64, to uint64) error {
	stmt := fmt.Sprintf("<%s> <follows> <%s> .", utils.StringUid(from), utils.StringUid(to))

	m := api.Mutation{
		SetNquads: []byte(stmt),
	}

	apiReq := api.Request{
		Mutations: []*api.Mutation{&m},
		CommitNow: true,
	}

	res, err := client.NewTxn().Do(context.Background(), &apiReq)
	if err != nil {
		return errx.Wrapf(errx.ERROR_DB, "failed to follow in dgraph, err:%v, json:%v", err, string(res.Json))
	}

	return nil
}

func JoinGroup(client *dgo.Dgraph, groupId uint64, userId uint64) error {
	g, u := utils.StringUid(groupId), utils.StringUid(userId)
	stmt := fmt.Sprintf("<%s> <members> <%s> .", g, u)

	qFormat := `
  query {
		group as var(func: uid("%s")) @filter(has("group_name"))
		user as var(func: uid("%s")) @filter(has("nickname"))
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
		return errx.Wrapf(errx.ERROR_GROUP_NOTFOUND, "user or group not found")
	}

	return nil
}
