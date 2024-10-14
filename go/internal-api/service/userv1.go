package service

import (
	"context"
	"time"

	"connectrpc.com/connect"
	userv1 "github.com/Shu-Tanabe/next-connect/go/internal-api/pb/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/google/uuid"
)

type UserService struct {}

func NewUserv1Server() *UserService {
	return &UserService{}
}

func (s *UserService) ListUsers(ctx context.Context, req *connect.Request[userv1.ListUsersRequest]) (*connect.Response[userv1.ListUsersResponse], error) {
	var nextPageToken string = ""
	res := connect.NewResponse(
		&userv1.ListUsersResponse{
			Users: []*userv1.User{
				{
					Id:   uuid.New().String(),
					UserName: "Alice",
					Email: "",
					Description: "",
					CreatedAt: time.Now().Local().String(),
				},
			},
			NextPageToken: &nextPageToken,
		},
	)

	return res, nil
}

func (s *UserService) GetUser(ctx context.Context, req *connect.Request[userv1.GetUserRequest]) (*connect.Response[userv1.User], error) {
	res := connect.NewResponse(
		&userv1.User{
			Id:   uuid.New().String(),
			UserName: "Alice",
			Email: "",
			Description: "",
			CreatedAt: time.Now().Local().String(),
		},
	)

	return res, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *connect.Request[userv1.CreateUserRequest]) (*connect.Response[userv1.User], error) {
	res := connect.NewResponse(
		&userv1.User{
			Id:   uuid.New().String(),
			UserName: req.Msg.UserName,
			Email: req.Msg.Email,
			Description: req.Msg.Description,
			CreatedAt: time.Now().Local().String(),
		},
	)

	return res, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *connect.Request[userv1.UpdateUserRequest]) (*connect.Response[userv1.User], error) {
	res := connect.NewResponse(
		&userv1.User{
			Id:   req.Msg.Id,
			UserName: *req.Msg.UserName,
			Email: *req.Msg.Email,
			Description: *req.Msg.Description,
			CreatedAt: time.Now().Local().String(),
		},
	)

	return res, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *connect.Request[userv1.DeleteUserRequest]) (*connect.Response[emptypb.Empty], error) {
	res := connect.NewResponse[emptypb.Empty](
		&emptypb.Empty{},
	)

	return res, nil
}