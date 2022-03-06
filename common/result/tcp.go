package result

import (
	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/pkg/codec"
	pb "github.com/wymli/bcsns/dependency/pb/tcp"
	"github.com/wymli/bcsns/pkg/server_framework/tcp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TcpResult(conn *tcp.ConnCtx, req protoreflect.ProtoMessage, rsp protoreflect.ProtoMessage, err error) {
	defer func() {
		if err != nil {
			conn.Logger.Error().Interface("req", req).Interface("rsp", rsp).Msg(err.Error())
		} else {
			conn.Logger.Info().Interface("req", req).Interface("rsp", rsp).Msg(err.Error())
		}
	}()

	if err != nil {
		stdErr := errx.ToApiError(err)
		rsp = &pb.CommonResp{
			Code: int64(stdErr.Code),
			Msg:  stdErr.Msg,
		}
	}

	respBody, err := proto.Marshal(rsp)

	frame := codec.TcpFrame{
		Length:       uint32(len(respBody)),
		Path:         0,
		CompressType: codec.COMPRESS_NONE,
		Payload:      respBody,
	}

	n, err := conn.Conn.Write(frame.Encode())
	if err != nil {
		logx.Errorf("failed to write msg to tcp.conn, err:%v", err)
	} else if n != frame.Size() {
		logx.Errorf("failed to write enough msg to tcp.conn, expected:%v actual:%v", frame.Size(), n)
	}
}
