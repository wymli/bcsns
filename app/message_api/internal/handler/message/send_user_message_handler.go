package message

import (
	"net/http"

	"github.com/wymli/bcsns/app/message_api/internal/logic/message"
	"github.com/wymli/bcsns/app/message_api/internal/svc"
	"github.com/wymli/bcsns/app/message_api/internal/types"
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SendUserMessageHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendUserMsgReq
		if err := httpx.Parse(r, &req); err != nil {
			err = errx.Wrapf(errx.ERROR_BAD_REQUEST, "failed to parse bind request params, err:%v", err)
			result.HttpApiResult(r, w, nil, nil, err)
			return
		}

		l := message.NewSendUserMessageLogic(r.Context(), ctx)
		resp, err := l.SendUserMessage(req)
		result.HttpApiResult(r, w, req, resp, err)
	}
}