package handler

import (
	"context"
	pb "github.com/amiraliio/publishen/publish/model/publish"
	"github.com/amiraliio/publishen/publish/repository"
	"net/http"
)

//Service client for publish handler
type Service struct {
	//other services client
}

func (s *Service) getRepository() repository.Adapter {
	return &repository.Repository{}
}

//CreatePublish service
func (s *Service) CreatePublish(ctx context.Context, req *pb.Publish, res *pb.Response) error {
	repo := s.getRepository()
	create, err := repo.Create(req)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		return nil
	}
	res.Code = http.StatusOK
	res.Data.Publish = create
	return nil
}

//GetPublish service
func (s *Service) GetPublish(ctx context.Context, req *pb.Publish, res *pb.Response) error {
	repo := s.getRepository()
	publish, err := repo.Get(req)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		return nil
	}
	res.Code = http.StatusOK
	res.Data.Publish = publish
	return nil
}

//ListOfUserPublishes publish service
func (s *Service) ListOfUserPublishes(ctx context.Context, req *pb.Request, res *pb.Response) error {
	repo := s.getRepository()
	publishes, err := repo.ListOfUserPublishes(req)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		return nil
	}
	res.Code = http.StatusOK
	res.Data.Publishes = publishes
	return nil
}

//UpdatePublish service
func (s *Service) UpdatePublish(ctx context.Context, req *pb.Publish, res *pb.Response) error {
	repo := s.getRepository()
	publish, err := repo.Update(req)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		return nil
	}
	res.Code = http.StatusOK
	res.Data.Publish = publish
	return nil
}

//DeletePublish service
func (s *Service) DeletePublish(ctx context.Context, req *pb.Publish, res *pb.Response) error {
	repo := s.getRepository()
	publish, err := repo.delete(req)
	if err != nil {
		res.Code = http.StatusInternalServerError
		res.Message = err.Error()
		return nil
	}
	res.Code = http.StatusOK
	res.Data.Publish = publish
	return nil
}
