package codec

import (
	"encoding/binary"
	"fmt"
	"net"

	"github.com/wymli/bcsns/common/server_framework/tcp"
)

type CompressType uint8

const (
	COMPRESS_NONE CompressType = iota
	COMPRESS_GZIP
	COMPRESS_ZIP
)

type TcpFrame struct {
	Length       uint32       // payload length
	Path         uint32       // 路由
	CompressType CompressType // 压缩类型
	Payload      []byte       // body
}

func (tf *TcpFrame) Size() int {
	return 9 + len(tf.Payload)
}

func (tf *TcpFrame) Encode() []byte {
	buf := make([]byte, 9+len(tf.Payload))
	binary.BigEndian.PutUint32(buf, tf.Length)
	binary.BigEndian.PutUint32(buf[4:], tf.Path)
	buf[8] = byte(tf.CompressType)
	copy(buf[9:], tf.Payload)
	return buf
}

func Framer() tcp.Framer {
	return func(conn net.Conn) ([]byte, error) {
		// todo: use bytebufferpool instead
		buf := make([]byte, 9)

		n, err := conn.Read(buf)
		if err != nil || n != 9 {
			return nil, fmt.Errorf("unable to read enough size frame header from socket,want 9, get %d, err:%s", n, err.Error())
		}

		len := binary.BigEndian.Uint32(buf[:4])

		buff := make([]byte, len)

		n, err = conn.Read(buff)
		if err != nil || n != int(len) {
			return nil, fmt.Errorf("unable to read enough size payload after read frame header,want %d, get %d, error: %s", len, n, err.Error())
		}

		res := make([]byte, 9+len)
		copy(res, buf)
		copy(res[9:], buff)

		return res, nil
	}
}

func Decoder() tcp.Decoder {
	return func(frame []byte) (uint32, []byte, error) {
		path := binary.BigEndian.Uint32(frame[4:8])
		compressType := CompressType(frame[8])
		rawPayload := frame[9:]

		switch compressType {
		case COMPRESS_NONE:
		case COMPRESS_ZIP:
			return 0, nil, fmt.Errorf("unsupported compress type")
		case COMPRESS_GZIP:
			return 0, nil, fmt.Errorf("unsupported compress type")
		}

		return path, rawPayload, nil
	}
}
