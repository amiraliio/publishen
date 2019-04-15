package main

import (
	"log"
	"time"

	"github.com/amiraliio/publishen/publish/handler"
	pb "github.com/amiraliio/publishen/publish/model/publish"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-config"
	registry "github.com/micro/go-plugins/registry/kubernetes"
	server "github.com/micro/go-plugins/server/grpc"
	transport "github.com/micro/go-plugins/transport/grpc"
)

func main() {

	err := config.LoadFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	srv := micro.NewService(
		micro.Name(config.Get("service","name")),
		micro.Version(config.Get("service","version")),
		micro.Server(server.NewServer()),
		micro.Transport(transport.NewTransport()),
		micro.Registry(registry.NewRegistry()),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	srv.Init()

	if err := pb.RegisterPublishServiceHandler(srv.Server(), new(handler.Service)); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
