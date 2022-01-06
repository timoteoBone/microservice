package util

import (
	err "github.com/timoteoBone/final-project-microservice/http-service/errors"

	"github.com/timoteoBone/final-project-microservice/grpc-service/entities"
)

func ValidateCreateUserRequest(user entities.CreateUserRequest) error {
	if user.Age < 1 || len(user.Name) < 5 || len(user.Pass) < 5 {
		return err.ErrInvalidDataForm
	}
	return nil
}
