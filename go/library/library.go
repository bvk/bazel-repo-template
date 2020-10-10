package library

import (
	"context"

	librarypb "github.com/bvk/bazel-repo-template/proto/library"
)

type Service struct {
}

func New() (*Service, error) {
	s := new(Service)
	return s, nil
}

func (s *Service) Close() error {
	return nil
}

func (s *Service) CreateBook(ctx context.Context, req *librarypb.CreateBookRequest) (*librarypb.CreateBookResponse, error) {
	resp := &librarypb.CreateBookResponse{
		Book: &librarypb.Book{
			Name: req.Name,
		},
	}
	return resp, nil
}
