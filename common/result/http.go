package result

import (
	"net/http"

	"github.com/wymli/bcsns/common/errx"

	"github.com/wymli/bcsns/common/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

// HttpResult 错误信息映射
func HttpApiResult(r *http.Request, w http.ResponseWriter, req interface{}, rsp interface{}, err error) {
	httpResult(r, w, req, rsp, err)
}

func httpResult(r *http.Request, w http.ResponseWriter, req interface{}, rsp interface{}, err error) {
	var ret interface{}
	if err == nil {
		ret = Success(rsp)
		logx.WithTraceCtx(r.Context()).Info().Msgf("%v %v req:%v rsp:%v", r.Method, r.URL.String(), req, ret)
		httpx.WriteJson(w, http.StatusOK, ret)
		return
	}

	apiErr := errx.ToStdError(err)

	ret = Error(apiErr.Code, apiErr.ApiMsg)
	logx.WithTraceCtx(r.Context()).Error().Msgf("%v %v req:%v rsp:%v err:%v", r.Method, r.URL.String(), req, ret, err)
	httpx.WriteJson(w, apiErr.ToHttpStatusCode(), ret)
}
