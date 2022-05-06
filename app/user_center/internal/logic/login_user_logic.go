package logic

import (
	"context"
	"encoding/json"
	"strconv"

	pbauth "github.com/wymli/bcsns/app/auth_rpc/pb"
	"github.com/wymli/bcsns/app/user_center/internal/model"
	"github.com/wymli/bcsns/app/user_center/internal/svc"
	"github.com/wymli/bcsns/app/user_center/pb"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/common/utils"
)

type LoginUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginUserLogic {
	return &LoginUserLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginUserLogic) LoginUser(in *pb.LoginUserReq) (*pb.LoginUserResp, error) {
	switch in.LoginType {
	case "sms_code":
		return nil, errx.Wrapf(errx.ERROR_SERVER_UNIMPLEMENTED, "unimplemented login_type: sms_code")
	case "password":
	default:
		return nil, errx.Wrapf(errx.ERROR_BAD_REQUEST, "invalid login_type: %v", in.LoginType)
	}

	txn := l.svcCtx.DgraphClient.NewTxn()
	res, err := txn.QueryWithVars(l.ctx, model.QLogin,
		map[string]string{"$phone": strconv.FormatInt(int64(in.Phone), 64), "$password": in.Password})
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_DB, "failed to check phone:%v password from dgraph", in.Phone)
	}

	login := model.LoginDecode{}
	err = json.Unmarshal(res.Json, &login)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to unmarshall json: %v, err:%v", string(res.Json), err)
	}

	switch {
	case len(login.Check) == 0:
		return nil, errx.Wrapf(errx.ERROR_USER_NOT_FOUND, "no user found of phone:%v", in.Phone)
	case len(login.Check) > 1:
		return nil, errx.Wrapf(errx.ERROR_USER_DUPLICATE, "duplicate user found of phone:%v", in.Phone)
	}

	if !login.Check[0].Pass {
		return nil, errx.ERROR_USER_WRONG_PWD
	}

	uid := utils.UintUid(login.Check[0].Uid)

	authRsp, err := l.svcCtx.AuthClient.GenerateToken(l.ctx, &pbauth.GenerateTokenReq{
		UserId: uid,
	})
	if err != nil {
		return nil, errx.Wrapf(err, "failed to generate token with userid:%d", uid)
	}

	return &pb.LoginUserResp{
		UserId:       uid,
		AccessToken:  authRsp.AccessToken,
		AccessExpire: authRsp.AccessExpire,
		RefreshAfter: authRsp.RefreshAfter,
	}, nil
}
