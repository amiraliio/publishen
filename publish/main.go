package main

import (
	"github.com/amiraliio/publishen/publish/handler"
	pb "github.com/amiraliio/publishen/publish/model/publish"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/config"
	"github.com/micro/go-micro/util/log"
	registry "github.com/micro/go-plugins/registry/nats"
	server "github.com/micro/go-plugins/server/grpc"
	transport "github.com/micro/go-plugins/transport/grpc"
)

func main() {

	err := config.LoadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	srv := micro.NewService(
		micro.Name(config.Get("service", "name").String("publishen_publish_service")),
		micro.Version(config.Get("service", "version").String("1.0.0")),
		micro.Server(server.NewServer()),
		micro.Transport(transport.NewTransport()),
		micro.Registry(registry.NewRegistry()),
	)

	srv.Init()

	if err := pb.RegisterPublishServiceHandler(srv.Server(), new(handler.Service)); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
