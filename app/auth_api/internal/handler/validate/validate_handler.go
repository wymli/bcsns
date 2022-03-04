package validate

import (
	"net/http"

	"github.com/wymli/bcsns/app/auth_api/internal/logic/validate"
	"github.com/wymli/bcsns/app/auth_api/internal/svc"
	"github.com/wymli/bcsns/app/auth_api/internal/types"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ValidateHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ValidateTokenReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errx.Wrapf(errx.ERROR_BAD_REQUEST, "failed to parse bind request params, err:%v", err)
			result.HttpApiResult(r, w, nil, nil, err)
			return
		}

		l := validate.NewValidateLogic(r.Context(), ctx)
		resp, err := l.Validate(req, r)
		result.HttpApiResult(r, w, req, resp, err)
	}
}
