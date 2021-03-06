package tcp

import (
	"context"
	"log"
	"net"

	"github.com/wymli/bcsns/common/logx"
)

type (
	Handler func(connCtx *ConnCtx, body []byte)
	Framer  func(conn net.Conn) ([]byte, error)        // return frame
	Decoder func(frame []byte) (uint32, []byte, error) // return path,body
)

type Route struct {
	Path    uint32
	Handler Handler
}

type Config struct {
	ListenOn string
}

type ConnCtx struct {
	Conn   net.Conn
	Ctx    context.Context
	Logger logx.Logger // todo: shoud be declared as an interface
}

type Server struct {
	Routes  []Route
	Framer  Framer
	Decoder Decoder
	Config  Config
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) RegisterFramer(framer Framer) {
	s.Framer = framer
}

func (s *Server) RegisterDecoder(decoder Decoder) {
	s.Decoder = decoder
}

func (s *Server) AddRoutes(routes []Route) {
	s.Routes = append(s.Routes, routes...)
}

func (s *Server) Start() {
	if s.Framer == nil {
		log.Fatal("tcp_server: set framer first")
	}
	if s.Decoder == nil {
		log.Fatal("tcp_server: set decoder first")
	}

	lis, err := net.Listen("tcp", s.Config.ListenOn)
	if err != nil {
		log.Fatalf("failed to listen on addr:%s, err:%v\n", s.Config.ListenOn, err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatalf("failed to accept connection from socket on addr:%s, error:%v", s.Config.ListenOn, err)
		}

		exit := make(chan struct{})
		go func() {
			defer func() {
				if err := recover(); err != nil {
					log.Printf("[error] goroutine catch a panic, err:%s", err)
				}
			}()

			for {
				frame, err := s.Framer(conn)
				if err != nil {
					log.Printf("failed to frame stream from connection, err:%v\n", err)
					close(exit)
					return
				}

				path, body, err := s.Decoder(frame)
				if err != nil {
					log.Printf("failed to decode package from frame, err:%v\n", err)
					close(exit)
					return
				}

				connCtx := &ConnCtx{
					Conn: conn,
					Ctx:  context.Background(),
				}

				// 当路由数<7左右时,array比map快
				for _, route := range s.Routes {
					if route.Path == path {
						route.Handler(connCtx, body)
						break
					}
				}
			}
		}()

		for range exit {
		}
		_ = conn.Close()

	}
}

func (s *Server) Stop() {
}

