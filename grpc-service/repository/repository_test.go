package repository_test

import (
	"context"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
	"github.com/timoteoBone/final-project-microservice/grpc-service/entities"
	"github.com/timoteoBone/final-project-microservice/grpc-service/mocks"
	"github.com/timoteoBone/final-project-microservice/grpc-service/repository"
)

func TestNewRepo(t *testing.T) {
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	db, _ := mocks.NewMock(logger)

	repo := repository.NewSQL(db, logger)

	assert.NotNil(t, repo)
}

func TestCreateUser(t *testing.T) {

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "grpcUserService",
			"time:", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}

	db, mock := mocks.NewMock(logger)
	defer db.Close()

	repo := repository.NewSQL(db, logger)

	testCases := []struct {
		Name              string
		Identifier        string
		User              entities.User
		Query             string
		ExpectedRespError error
	}{

		{
			Name:       "Create User Valid Case",
			Identifier: "CreateUser",
			User: entities.User{
				Name: "Timo",
				Age:  19,
				Pass: "1234",
			},
			Query:             "INSERT INTO USER (first_name, age, pass) VALUES(?,?,?)",
			ExpectedRespError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			ctx := context.Background()

			prep := mock.ExpectPrepare(tc.Query)
			prep.ExpectExec().WithArgs(tc.User.Name, tc.User.Age, tc.User.Pass).WillReturnResult(sqlmock.NewResult(int64(1), int64(1)))

			_, err := repo.CreateUser(ctx, tc.User)
			assert.NoError(t, err)
		})
	}
}
