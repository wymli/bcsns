package mq

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"github.com/wymli/bcsns/common/logx"
)

type Handler func(value []byte) error

type Route struct {
	Topic   string
	Handler Handler
	metric  metric
}

type metric struct {
	total int64
	fail  int64
}

type Server struct {
	Routes  []Route
	Brokers []string
	GroupId string
}

func (s *Server) AddRoutes(routes []Route) {
	s.Routes = append(s.Routes, routes...)
}

func NewServer(brokers []string, groupId string) *Server {
	return &Server{
		Brokers: brokers,
		GroupId: groupId,
	}
}

func (s *Server) Start() {
	closeCh := make(chan struct{})
	for i := range s.Routes {
		route := &s.Routes[i]
		reader := kafka.NewReader(kafka.ReaderConfig{
			Brokers:  s.Brokers,
			Topic:    route.Topic,
			GroupID:  s.GroupId,
			MinBytes: 10e3, // 10KB
			MaxBytes: 10e6, // 10MB
		})

		go func() {
			for {
				msg, err := reader.ReadMessage(context.Background())
				if err != nil {
					logx.Errorf("failed to fetch msg from kafka with topic:%v, err:%v", route.Topic, err)
					close(closeCh)
					return
				}
				err = route.Handler(msg.Value)
				if err != nil {
					// logx.Errorf("failed to consume and handle msg from kafka, err:%v", err)
					// 不在这里打log,拿不到业务信息,让handler自己打log
					route.metric.fail++
				}
				route.metric.total++
			}
		}()
	}

	// 收集metric
	// go func() {
	// 	for {
	// 		//
	// 	}
	// }()

	for range closeCh {
		return
	}

	return
}

func (s *Server) Stop() {
	for _, route := range s.Routes {
		fmt.Printf("[exit] topic:%s fail:%d total:%d\n", route.Topic, route.metric.fail, route.metric.total)
	}
}
