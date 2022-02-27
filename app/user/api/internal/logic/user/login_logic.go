package user

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/wymli/bcsns/app/auth/rpc/pb"
	"github.com/wymli/bcsns/app/user/api/internal/svc"
	"github.com/wymli/bcsns/app/user/api/internal/types"
	"github.com/wymli/bcsns/common/errx"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) LoginLogic {
	return LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var qLogin = `
query check($phone: int, $password: string){
  check(func: eq(phone, $phone)) {
    pass: checkpwd(password, $password)
  }
}
`

func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginResp, err error) {
	switch req.LoginType {
	case "sms_code":
		return nil, errx.Wrapf(errx.ERROR_SERVER_UNIMPLEMENTED, "unimplemented login_type: sms_code")
	case "password":
	default:
		return nil, errx.Wrapf(errx.ERROR_BAD_REQUEST, "invalid login_type: %v", req.LoginType)
	}

	txn := l.svcCtx.DgraphClient.NewTxn()
	res, err := txn.QueryWithVars(context.Background(), qLogin,
		map[string]string{"$phone": strconv.FormatInt(int64(req.Phone), 64), "$password": req.Password})
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_DB, "failed to check phone:%v password from dgraph", req.Phone)
	}

	var decode struct {
		Check []struct {
			Uid  string
			Pass bool
		}
	}

	err = json.Unmarshal(res.Json, &decode)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to unmarshall json: %v, err:%v", string(res.Json), err)
	}

	switch {
	case len(decode.Check) == 0:
		return nil, errx.Wrapf(errx.ERROR_USER_NOT_FOUND, "no user found of phone:%v", req.Phone)
	case len(decode.Check) > 1:
		return nil, errx.Wrapf(errx.ERROR_USER_DUPLICATE, "duplicate user found of phone:%v", req.Phone)
	}

	if !decode.Check[0].Pass {
		return nil, errx.ERROR_USER_WRONG_PWD
	}

	uid, err := strconv.ParseUint(decode.Check[0].Uid, 16, 64)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to parseUint: %v, err:%v", decode.Check[0].Uid, err)
	}

	authRsp, err := l.svcCtx.AuthClient.GenerateToken(context.Background(), &pb.GenerateTokenReq{
		UserId: uid,
	})
	if err != nil {
		return nil, errx.Wrapf(err, "failed to generate token with userid:%d", uid)
	}

	return &types.LoginResp{
		UserId:       uid,
		AccessToken:  authRsp.AccessToken,
		AccessExpire: authRsp.AccessExpire,
		RefreshAfter: authRsp.RefreshAfter,
	}, nil
}
