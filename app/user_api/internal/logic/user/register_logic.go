package user

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/dgraph-io/dgo/v210/protos/api"
	"github.com/pkg/errors"
	"github.com/wymli/bcsns/app/auth_rpc/pb"
	"github.com/wymli/bcsns/app/user_api/internal/svc"
	"github.com/wymli/bcsns/app/user_api/internal/types"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/utils"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) RegisterLogic {
	return RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req types.RegisterReq) (resp *types.RegisterResp, err error) {
	c := l.svcCtx.DgraphClient

	user := types.User{
		Phone:     req.Phone,
		Nickname:  req.Nickname,
		Sex:       req.Sex,
		Age:       req.Age,
		Avatar:    req.Avatar,
		Address:   req.Address,
		PublicKey: req.PublicKey,
		Password:  req.Password,
	}

	userJson, err := json.Marshal(user)
	if err != nil {
		return nil, errors.Wrapf(errx.ERROR_SERVER_COMMON, "failed to marshall json, err:%v", err)
	}

	qFormat := `
  query {
      user as var(func: eq(phone, "%s"))
  }`
	query := fmt.Sprintf(qFormat, req.Phone)

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

	res, err := c.NewTxn().Do(context.Background(), apiReq)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_DB, "failed to insert user in dgraph, err:%v", err)
	}

	if len(res.Uids) == 0 {
		return nil, errx.Wrapf(errx.ERROR_USER_DUPLICATE, "duplicate register of user phone: %v", req.Phone)
	}

	_, v, ok := utils.ExtractOneFromStringMap(res.Uids)
	if !ok {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to extract uid from dgraph response, want 1 uid, get %d uid", len(res.Uids))
	}

	uid, err := strconv.ParseUint(v, 16, 64)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to convert uid from string to uint64: %v, err: %v", uid, err)
	}

	authRsp, err := l.svcCtx.AuthClient.GenerateToken(context.Background(), &pb.GenerateTokenReq{
		UserId: uid,
	})
	if err != nil {
		return nil, errx.Wrapf(err, "failed to generate token with userid:%d", uid)
	}

	return &types.RegisterResp{
		UserId:       uid,
		AccessToken:  authRsp.AccessToken,
		AccessExpire: authRsp.AccessExpire,
		RefreshAfter: authRsp.RefreshAfter,
	}, nil
}
