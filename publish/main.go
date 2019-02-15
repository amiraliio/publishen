package main

import (
	"github.com/amiraliio/publishen/publish/handler"
	pb "github.com/amiraliio/publishen/publish/model/publish"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"time"
)

func main() {

	srv := micro.NewService(
		micro.Name("publishen_publish_service"),
		micro.Version("1.0.0"),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)

	srv.Init()

	if err := pb.RegisterPublishServiceHandler(srv.Server(), new(handler.Service)); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
