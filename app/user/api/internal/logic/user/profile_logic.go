package user

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/wymli/bcsns/app/user/api/internal/svc"
	"github.com/wymli/bcsns/app/user/api/internal/types"
	"github.com/wymli/bcsns/common/errx"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) ProfileLogic {
	return ProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var qProfile = `query all($uid: string) {
	all(func: uid($uid)) {
		uid
		phone
		nickname
		sex
		age
		avatar
		address
		public_key
	}
}`

func (l *ProfileLogic) Profile(req types.UserProfileReq) (resp *types.UserProfileResp, err error) {
	txn := l.svcCtx.DgraphClient.NewTxn()
	res, err := txn.QueryWithVars(context.Background(), qProfile, map[string]string{"$uid": "0x" + strconv.FormatUint(req.UserId, 16)})
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_DB, "failed to query uid:%v profile from dgraph", req.UserId)
	}

	var decode struct {
		All []types.User
	}

	err = json.Unmarshal(res.Json, &decode)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to unmarshall json: %v, err:%v", string(res.Json), err)
	}

	switch {
	case len(decode.All) == 0:
		return nil, errx.Wrapf(errx.ERROR_USER_NOT_FOUND, "no user found of uid:%v", req.UserId)
	case len(decode.All) > 1:
		return nil, errx.Wrapf(errx.ERROR_USER_DUPLICATE, "duplicate user found of uid:%v", req.UserId)
	}

	user := decode.All[0]

	return &types.UserProfileResp{
		UserId:    req.UserId,
		Nickname:  user.Nickname,
		Sex:       user.Sex,
		Age:       user.Sex,
		Avatar:    user.Avatar,
		Address:   user.Address,
		PublicKey: user.PublicKey,
	}, nil
}
