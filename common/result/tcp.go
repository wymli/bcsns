package result

import (
	"net"

	"github.com/wymli/bcsns/common/errx"
	"github.com/wymli/bcsns/common/logx"
	"github.com/wymli/bcsns/dependency/codec"
	pb "github.com/wymli/bcsns/dependency/pb/tcp"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TcpResult(conn net.Conn, req protoreflect.ProtoMessage, rsp protoreflect.ProtoMessage, err error) {
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

	logx.Infof("req:%#v rsp:%#v err:%v", req, rsp, err)

	n, err := conn.Write(frame.Encode())
	if err != nil {
		logx.Errorf("failed to write msg to tcp.conn, err:%v", err)
	} else if n != frame.Size() {
		logx.Errorf("failed to write enough msg to tcp.conn, expected:%v actual:%v", frame.Size(), n)
	}
}
