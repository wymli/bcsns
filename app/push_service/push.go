package main

import (
	"flag"
	"fmt"

	"github.com/wymli/bcsns/app/push_service/internal/config"
	"github.com/wymli/bcsns/app/push_service/internal/handler"
	"github.com/wymli/bcsns/app/push_service/internal/svc"
	"github.com/wymli/bcsns/pkg/mq"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/message.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := mq.NewServer(c.Kafka.Broker.Endpoints, c.Kafka.GroupId)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting push consumer")
	server.Start()
}
