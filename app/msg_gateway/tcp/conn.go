package tcp

import (
	"fmt"
	"net"
	"sync"

	"github.com/wymli/bcsns/common/codec"
	"github.com/wymli/bcsns/common/errx"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type UserConn struct {
	UserId uint64
	Conn   net.Conn
	exit   bool
}

func NewUserConn(userId uint64, conn net.Conn) *UserConn {
	return &UserConn{
		UserId: userId,
		Conn:   conn,
	}
}

// Send ...
func (c *UserConn) Send(bs []byte) error {
	n, err := c.Conn.Write(bs)
	if err != nil {
		return errx.Wrapf(errx.ERROR_IO, "failed to write frame to socket, err:%v", err)
	}

	if n != len(bs) {
		return errx.Wrapf(errx.ERROR_IO, "failed to write enough bytes of frame to socket, expected:%v, actual:%v", len(bs), n)
	}

	return nil
}

// SendFrame ...
func (c *UserConn) SendFrame(frame codec.TcpFrame) error {
	return c.Send(frame.Encode())
}

// SendPBMsg ...
func (c *UserConn) SendPBMsg(msg protoreflect.ProtoMessage) error {
	bs, err := proto.Marshal(msg)
	if err != nil {
		return errx.Wrapf(errx.ERROR_MARSHALL, "failed to marshall proto msg:%v, err:%v", msg, err)
	}

	frame := codec.TcpFrame{
		Length:       uint32(len(bs)),
		Path:         0,
		CompressType: codec.COMPRESS_NONE,
		Payload:      bs,
	}

	return c.SendFrame(frame)
}

// UserConnPool ...
type UserConnPool struct {
	rwm      sync.RWMutex
	ConnPool map[uint64]*UserConn
}

// NewUserConnPool ...
func NewUserConnPool() *UserConnPool {
	return &UserConnPool{
		rwm:      sync.RWMutex{},
		ConnPool: map[uint64]*UserConn{},
	}
}

// AddConn ...
func (c *UserConnPool) AddConn(userId uint64, conn net.Conn) {
	c.rwm.Lock()
	defer c.rwm.Unlock()

	c.ConnPool[userId] = NewUserConn(userId, conn)
}

// RemoveConn ...
func (c *UserConnPool) CloseConn(userId uint64) error {
	c.rwm.Lock()
	defer c.rwm.Unlock()

	conn, exist := c.GetConn(userId)
	if !exist {
		return fmt.Errorf("failed to close a not-found connection of userId:%v", userId)
	}

	defer delete(c.ConnPool, conn.UserId)

	err := conn.Conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection of userId:%v, err:%v", userId, err)
	}

	return nil
}

// GetConn ...
func (c *UserConnPool) GetConn(userId uint64) (*UserConn, bool) {
	c.rwm.RLock()
	defer c.rwm.RUnlock()

	conn, exists := c.ConnPool[userId]
	return conn, exists
}

// HasConn ...
func (c *UserConnPool) HasConn(userId uint64) bool {
	c.rwm.RLock()
	defer c.rwm.RUnlock()

	_, exists := c.ConnPool[userId]
	return exists
}
