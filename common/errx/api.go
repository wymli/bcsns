package errx

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StdError struct {
	Code uint32
	Msg  string
}

func (e StdError) Error() string {
	return fmt.Sprintf("Code:%d, Msg:%s", e.Code, e.Msg)
}

// api 对外返回
// 4位错误码
var (
	SUCCESS              StdError = StdError{0, "成功"}
	ERROR_SERVER_COMMON  StdError = StdError{1001, "服务器内部错误"}
	ERROR_BAD_REQUEST    StdError = StdError{1002, "参数错误"}
	ERROR_TOKEN_INVALID  StdError = StdError{1003, "无效的Token"}
	ERROR_USER_NOT_FOUND StdError = StdError{1005, "该用户不存在"}
	ERROR_USER_DUPLICATE StdError = StdError{1006, "存在重复用户"}
	ERROR_USER_WRONG_PWD StdError = StdError{1007, "密码错误"}
	ERROR_USER_UNAUTHEN  StdError = StdError{1007, "未登录"}

	ERROR_SERVER_UNIMPLEMENTED StdError = StdError{2001, "功能未实现"}

	ERROR_MQ    StdError = StdError{3001, "服务器内部错误"}
	ERROR_DB    StdError = StdError{3002, "数据库繁忙"}
	ERROR_REDIS StdError = StdError{3003, "服务器内部错误"}
	ERROR_SEQID StdError = StdError{3004, "服务器内部错误"}

	ERROR_UNKNOWN StdError = StdError{9999, "未知错误"}
)

var code2ErrMap = map[uint32]StdError{
	0:    SUCCESS,
	1001: ERROR_SERVER_COMMON,
	1002: ERROR_BAD_REQUEST,
	1003: ERROR_TOKEN_INVALID,
	1004: ERROR_DB,
	1005: ERROR_USER_NOT_FOUND,
	1006: ERROR_UNKNOWN,

	2001: ERROR_SERVER_UNIMPLEMENTED,
}

func ToApiErrFromCode(code uint32) StdError {
	stdErr, ok := code2ErrMap[code]
	if !ok {
		return ERROR_SERVER_COMMON
	}
	return stdErr
}

func (cm *StdError) ToHttpStatusCode() int {
	switch cm.Code {
	case 0:
		return http.StatusOK
	case 1001, 1004, 1006:
		return http.StatusInternalServerError
	case 1002:
		return http.StatusBadRequest
	case 1003:
		return http.StatusUnauthorized
	case 1005:
		return http.StatusUnprocessableEntity
	default:
		return http.StatusInternalServerError
	}
}

// ToGrpcError 将各种错误转成grpc error格式(主要为了保留错误码)
// 两种情况
// 1. 底层是grpc error
// 2. 底层是CodeMsg标准错误
func ToGrpcError(err error) error {
	if err == nil {
		return nil
	}

	iErr := errors.Cause(err)

	gErr, ok := status.FromError(iErr)
	if ok {
		// 透传底层错误码,但返回所有层的信息
		return status.Error(gErr.Code(), err.Error())
	}

	stdErr := ERROR_UNKNOWN
	switch z := iErr.(type) {
	case StdError:
		stdErr = z
	case *StdError:
		stdErr = *z
	default:
	}

	// 透传底层错误码,但返回所有层的信息
	return status.Error(codes.Code(stdErr.Code), err.Error())
}

func ToGrpcErrorf(err error, format string, v ...interface{}) error {
	err = errors.Wrapf(err, format, v...)
	return ToGrpcError(err)
}

// func IsStdError(err error, stdErr StdError) bool {
// 	gErr, ok := status.FromError(err)
// 	if !ok {

// 	}

// 	if gErr.Code() == codes.Code(stdErr.Code) {
// 		return true
// 	}
// 	return false
// }

// func FromError(err error) StdError {
// 	switch z := err.(type) {
// 	case StdError:
// 		return z
// 	case *StdError:
// 		return *z
// 	default:
// 		return ERROR_UNKNOWN
// 	}
// }

// ToApiError 返回标准的给api的错误
// 两种情况
// 1. 底层是grpc error,那么根据错误码转成标准错误
// 2. 底层直接是自定义的CodeMsg,那么我们相信一定是一个标准错误,直接返回
func ToApiError(err error) StdError {
	iErr := errors.Cause(err)
	gErr, ok := status.FromError(iErr)
	if ok {
		return ToApiErrFromCode(uint32(gErr.Code()))
	}

	switch z := iErr.(type) {
	case StdError:
		return z
	case *StdError:
		return *z
	default:
		return ERROR_UNKNOWN
	}
}

func IsGrpcBizErr(code codes.Code) bool {
	if _, ok := code2ErrMap[uint32(code)]; ok {
		return true
	}
	return false
}

func Wrapf(err error, format string, v ...interface{}) error {
	return errors.Wrapf(err, format, v...)
}

func Wrap(err error, msg string) error {
	return errors.Wrapf(err, msg)
}
