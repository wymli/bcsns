package logic

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	"github.com/wymli/bcsns/app/auth_rpc/internal/svc"
	"github.com/wymli/bcsns/app/auth_rpc/pb"
	"github.com/wymli/bcsns/common/errx"

	"github.com/wymli/bcsns/common/logx"
)

type ValidateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateTokenLogic {
	return &ValidateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithTraceCtx(ctx),
	}
}

//  validateToken
func (l *ValidateTokenLogic) ValidateToken(in *pb.ValidateTokenReq) (*pb.ValidateTokenResp, error) {
	// userTokenKey := fmt.Sprintf(CacheUserTokenKey, in.UserId)
	// dbToken, err := l.svcCtx.RedisClient.Get(userTokenKey)
	// if err != nil {
	// 	return nil, errors.Wrapf(common.ErrValidateToken, "failed to get redis, userId:%d, err:%v", in.UserId, err)
	// }

	token, err := jwt.Parse(in.Token, func(t *jwt.Token) (interface{}, error) {
		return []byte(l.svcCtx.Config.JwtAuth.AccessSecret), nil
	})
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to parse jwt, err:%v", err)
	}

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return &pb.ValidateTokenResp{
			UserId: 0,
			Ok:     false,
		}, nil
	}

	if !token.Valid {
		return &pb.ValidateTokenResp{
			UserId: 0,
			Ok:     false,
		}, nil
	}

	return &pb.ValidateTokenResp{
		UserId: token.Claims.(jwt.MapClaims)["userid"].(uint64),
		Ok:     true,
	}, nil
}
