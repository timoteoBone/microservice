package mapper

import "github.com/timoteoBone/final-project-microservice/grpc-service/entities"

func CreateUserRequestToUser(userReq entities.CreateUserRequest) entities.User {

	user := entities.User{
		userReq.Name,
		userReq.Pass,
		userReq.Age,
	}
	return user
}
