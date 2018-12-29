package main

import (
	"github.com/amiraliio/publishen/product/handler"
	pb "github.com/amiraliio/publishen/product/model/product"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

func main() {

	srv := micro.NewService(
		micro.Name("publishen_product_service"),
		micro.Version("0.1"),
	)

	srv.Init()

	if err := pb.RegisterProductServiceHandler(srv.Server(), new(handler.Service)); err != nil {
		log.Fatal(err)
	}

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
