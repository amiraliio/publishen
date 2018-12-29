package handler

import (
	"context"
	pb "github.com/amiraliio/publishen/product/model/product"
	"github.com/amiraliio/publishen/product/repository"
)

//Services client for product handler
type Service struct {
	//other services client
}

func (s *Service) getRepository() repository.Adapter {
	return &repository.Repository{}
}

//CreateProduct service
func (s *Service) CreateProduct(ctx context.Context, req *pb.Product, rsp *pb.Response) error {
	repo := s.getRepository()
	create, err := repo.Create(req)
	if err != nil {
		rsp.Success = false
		rsp.ErrorMessage = err.Error()
		return nil
	}
	rsp.DocumentID = create
	rsp.Success = true
	return nil
}
