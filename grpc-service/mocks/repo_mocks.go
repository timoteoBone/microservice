package mocks

import (
	"context"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/mock"
	"github.com/timoteoBone/final-project-microservice/grpc-service/entities"
)

type RepoSitoryMock struct {
	mock.Mock
}

func (repo *RepoSitoryMock) CreateUser(ctx context.Context, user entities.User) (int64, error) {
	args := repo.Called(ctx, user)
	fmt.Println(args.Int(0))

	return int64(args.Int(0)), args.Error(1)
}

func (repo *RepoSitoryMock) GetUser(ctx context.Context, userId int64) (entities.User, error) {
	args := repo.Called(ctx, userId)
	id := args[0]

	if id == nil {
		return entities.User{}, args.Error(1)
	}

	return id.(entities.User), nil

}
