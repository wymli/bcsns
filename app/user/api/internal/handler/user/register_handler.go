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

func RegisterHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RegisterReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errx.Wrapf(errx.ERROR_BAD_REQUEST, "failed to parse bind request params, err:%v", err)
			result.HttpApiResult(r, w, nil, nil, err)
			return
		}

		l := user.NewRegisterLogic(r.Context(), ctx)
		resp, err := l.Register(req)
		result.HttpApiResult(r, w, req, resp, err)
	}
}
