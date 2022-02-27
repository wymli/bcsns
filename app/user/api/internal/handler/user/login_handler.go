package user

import (
	"net/http"

	"github.com/wymli/bcsns/app/user/api/internal/logic/user"
	"github.com/wymli/bcsns/app/user/api/internal/svc"
	"github.com/wymli/bcsns/app/user/api/internal/types"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoginHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errx.Wrapf(errx.ERROR_BAD_REQUEST, "failed to parse bind request params, err:%v", err)
			result.HttpApiResult(r, w, nil, nil, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), ctx)
		resp, err := l.Login(req)
		result.HttpApiResult(r, w, req, resp, err)
	}
}
