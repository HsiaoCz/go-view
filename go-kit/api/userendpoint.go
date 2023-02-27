package api

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type UserRequest struct {
	UserID int `json:"user_id"`
}

type UserResponse struct {
	Result string `json:"result"`
}

func GenUserEndpoint(userService IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		r := request.(UserRequest)
		result := userService.GetUsername(r.UserID)
		return UserResponse{Result: result}, nil
	}
}
