package service

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/timoteoBone/final-project-microservice/grpc-service/entities"
)

type Repository interface {
	CreateUser(ctx context.Context, rq entities.CreateUserRequest) (entities.CreateUserResponse, error)
	GetUser(ctx context.Context, rq entities.GetUserRequest) (entities.GetUserResponse, error)
}

type service struct {
	Repo   Repository
	Logger log.Logger
}

func NewService(repo Repository, logger log.Logger) *service {
	return &service{Repo: repo, Logger: logger}
}

func (s *service) CreateUser(ctx context.Context, rq entities.CreateUserRequest) (entities.CreateUserResponse, error) {

	logger := log.With(s.Logger, "create user request", "recevied")

	res, err := s.Repo.CreateUser(ctx, rq)

	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return entities.CreateUserResponse{}, err
	}

	return res, nil

}

func (s *service) GetUser(ctx context.Context, rq entities.GetUserRequest) (entities.GetUserResponse, error) {
	logger := log.With(s.Logger, "get user request", "recevied")

	res, err := s.Repo.GetUser(ctx, rq)
	if err != nil {
		level.Error(logger).Log("error", err.Error())
		return entities.GetUserResponse{}, err
	}

	return res, nil
}
