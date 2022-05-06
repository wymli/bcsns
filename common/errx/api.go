package errx

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type StdError struct {
	Code     uint32
	InnerMsg string
	ApiMsg   string
}

func (e StdError) Error() string {
	return fmt.Sprintf("Code:%d, InnerMsg:%s, ApiMsg:%s", e.Code, e.InnerMsg, e.ApiMsg)
}

// api 对外返回
// 4位错误码
var (
	SUCCESS             StdError = StdError{0, "", "成功"}
	ERROR_BAD_REQUEST   StdError = StdError{1002, "", "参数错误"}
	ERROR_TOKEN_INVALID StdError = StdError{1003, "", "无效的Token"}

	ERROR_USER_NOT_FOUND StdError = StdError{2001, "", "该用户不存在"}
	ERROR_USER_DUPLICATE StdError = StdError{2002, "", "存在重复用户"}
	ERROR_USER_WRONG_PWD StdError = StdError{2003, "", "密码错误"}
	ERROR_USER_UNAUTHEN  StdError = StdError{2004, "", "用户未登录"}
	ERROR_USER_OFFLINE   StdError = StdError{2005, "", "用户未在线"}

	ERROR_GROUP_DUPLICATE StdError = StdError{2051, "", "存在重复群组"}
	ERROR_GROUP_NOTFOUND  StdError = StdError{2052, "", "该群组不存在"}

	ERROR_SERVER_COMMON          StdError = StdError{3000, "一般性错误", "服务器内部错误"}
	ERROR_MQ                     StdError = StdError{3001, "消息队列错误", "服务器内部错误"}
	ERROR_DB                     StdError = StdError{3002, "数据库错误", "数据库繁忙"}
	ERROR_REDIS                  StdError = StdError{3003, "Redis错误", "服务器内部错误"}
	ERROR_SEQID                  StdError = StdError{3004, "序列号错误", "服务器内部错误"}
	ERROR_GATEWAY_USER_NOT_FOUND StdError = StdError{3005, "用户不在此gateway上", "服务器内部状态错误"}
	ERROR_IO                     StdError = StdError{3006, "socket io 错误", "服务器内部错误"}
	ERROR_MARSHALL               StdError = StdError{3007, "序列化错误", "服务器内部错误"}
	ERROR_PART                   StdError = StdError{3008, "部分错误", "服务器内部错误"}

	ERROR_BC_GAS         StdError = StdError{4001, "gas 相关错误", "区块链错误"}
	ERROR_BC_NONCE       StdError = StdError{4002, "nonce 相关错误", "区块链错误"}
	ERROR_BC_TRANSACTION StdError = StdError{4003, "交易错误", "区块链错误"}

	ERROR_UNKNOWN              StdError = StdError{9001, "", "未知错误"}
	ERROR_SERVER_UNIMPLEMENTED StdError = StdError{9002, "", "功能未实现"}
)

var code2ErrMap = map[uint32]StdError{
	0:    SUCCESS,
	1002: ERROR_BAD_REQUEST,
	1003: ERROR_TOKEN_INVALID,
	2001: ERROR_USER_NOT_FOUND,
	2002: ERROR_USER_DUPLICATE,
	2003: ERROR_USER_WRONG_PWD,
	2004: ERROR_USER_UNAUTHEN,
	2005: ERROR_USER_OFFLINE,
	3000: ERROR_SERVER_COMMON,
	3001: ERROR_MQ,
	3002: ERROR_DB,
	3003: ERROR_REDIS,
	3004: ERROR_SEQID,
	3005: ERROR_GATEWAY_USER_NOT_FOUND,
	3006: ERROR_IO,
	3007: ERROR_MARSHALL,
	3008: ERROR_PART,
	9001: ERROR_UNKNOWN,
	9002: ERROR_SERVER_UNIMPLEMENTED,
}

func ToStdErrFromCode(code uint32) StdError {
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
	case 2003, 2004:
		return http.StatusUnauthorized
	case 2001:
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

// ToStdError 返回标准的给api的错误
// 两种情况
// 1. 底层是grpc error,那么根据错误码转成标准错误
// 2. 底层直接是自定义的CodeMsg,那么我们相信一定是一个标准错误,直接返回
func ToStdError(err error) StdError {
	iErr := errors.Cause(err)
	gErr, ok := status.FromError(iErr)
	if ok {
		return ToStdErrFromCode(uint32(gErr.Code()))
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

func Is(err1 error, err2 StdError) bool {
	return ToStdError(err1).Code == err2.Code
}
