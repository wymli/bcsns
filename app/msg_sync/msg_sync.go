package main

import (
	"flag"
	"fmt"

	"github.com/wymli/bcsns/app/msg_sync/internal/config"
	"github.com/wymli/bcsns/app/msg_sync/internal/handler"
	"github.com/wymli/bcsns/app/msg_sync/internal/svc"
	"github.com/wymli/bcsns/common/server_framework/mq"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/msg_sync.yaml", "the config file")

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
