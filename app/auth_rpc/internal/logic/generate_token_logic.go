package logic

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/wymli/bcsns/app/auth_rpc/internal/svc"
	"github.com/wymli/bcsns/app/auth_rpc/pb"
	"github.com/wymli/bcsns/common/errx"

	"github.com/wymli/bcsns/common/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithTraceCtx(ctx),
	}
}

//  生成token，只针对用户服务开放访问
func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.generateJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.UserId)
	if err != nil {
		return nil, errx.Wrapf(errx.ERROR_SERVER_COMMON, "failed to generate jwt token, err:%v", err)
	}

	// userTokenKey := fmt.Sprintf(CacheUserTokenKey, in.UserId)
	// err = l.svcCtx.RedisClient.Setex(userTokenKey, accessToken, int(accessExpire))
	// if err != nil {
	// 	return nil, errors.Wrapf(common.ErrGenerateToken, "failed to store token to redis, userId:%d, err:%v", in.UserId, err)
	// }

	return &pb.GenerateTokenResp{
		AccessToken:  accessToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *GenerateTokenLogic) generateJwtToken(secretKey string, iat, exp int64, userId uint64) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    iat + exp,
		"iat":    iat,
		"userid": userId,
	}).SignedString([]byte(secretKey))
}
