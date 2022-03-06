package validate

import (
	"context"
	"net/http"
	"strings"

	"github.com/wymli/bcsns/app/auth_api/internal/svc"
	"github.com/wymli/bcsns/app/auth_api/internal/types"
	"github.com/wymli/bcsns/app/auth_rpc/auth"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
)

const CtxKeyJwtUserId = "jwtUserId"

type ValidateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateLogic(ctx context.Context, svcCtx *svc.ServiceContext) ValidateLogic {
	return ValidateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithTraceCtx(ctx),
	}
}

func (l *ValidateLogic) Validate(req types.ValidateTokenReq, r *http.Request) (resp *types.ValidateTokenResp, err error) {
	realRequestPath := r.Header.Get("X-Original-Uri")
	if realRequestPath == "" {
		realRequestPath = r.RequestURI
	}

	if strings.Contains(realRequestPath, "?") {
		realRequestPath = strings.Split(realRequestPath, "?")[0]
	}

	if l.urlNoAuth(realRequestPath) {
		return &types.ValidateTokenResp{
			UserId: 0,
			Ok:     true,
		}, nil
	}

	// 获取token
	token := ""
	authorization := r.Header.Get("Authorization")
	cookie := r.Header.Get("Cookie")
	switch {
	case authorization != "":
		token = authorization
	case cookie != "":
		token = cookie
	default:
	}

	// if token == "" {
	// 	return &types.ValidateTokenResp{
	// 		UserId: "",
	// 		Ok:     false,
	// 	}, nil
	// }

	// 解析用户
	rsp, err := l.svcCtx.AuthRpc.ValidateToken(l.ctx, &auth.ValidateTokenReq{
		UserId: 0,
		Token:  token,
	})
	if err != nil {
		return nil, errx.Wrap(err, "failed to validate token")
	}

	return &types.ValidateTokenResp{
		UserId: rsp.UserId,
		Ok:     rsp.Ok,
	}, nil
}

// 验证 url 是否 `不` 需要认证
func (l *ValidateLogic) urlNoAuth(path string) bool {
	for _, val := range l.svcCtx.Config.NoAuthUrls {
		if val == path {
			return true
		}
	}
	return false
}
