package handler

import (
	"context"
	pb "github.com/amiraliio/publishen/publish/model/publish"
	"github.com/amiraliio/publishen/publish/repository"
)

//Service client for publish handler
type Service struct {
	//other services client
}

func (s *Service) getRepository() repository.Adapter {
	return &repository.Repository{}
}

//CreatePublish service
func (s *Service) CreatePublish(ctx context.Context, req *pb.Publish, rsp *pb.Response) error {
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

//GetPublish service
func (s *Service) GetPublish(ctx context.Context, req *pb.Publish, res *pb.Response) error {
	repo := s.getRepository()
	publish, err := repo.Get(req)
	if err != nil {
		res.Success = false
		res.ErrorMessage = err.Error()
		return nil
	}
	res.Success = true
	res.Publish = publish
	return nil
}


//ListOfPublishes publish service
func (s *Service) ListOfPublishes(ctx context.Context, req *pb.Request, res *pb.Response) error {
	repo := s.getRepository()
	publishes, err := repo.List(req)
	if err !=nil{
		res.Success = false
		res.ErrorMessage = err.Error()
		return nil
	}
	res.Success = true
	res.Publishes = publishes
	return nil
}