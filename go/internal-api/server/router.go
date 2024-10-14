package server

import (
	"net/http"

	"connectrpc.com/grpcreflect"

	"github.com/Shu-Tanabe/next-connect/go/internal-api/pb/user/v1/userv1connect"
	"github.com/Shu-Tanabe/next-connect/go/internal-api/service"
)

func newRouter() *http.ServeMux {
	userv1Server := service.NewUserv1Server()

	router := http.NewServeMux()
	reflector := grpcreflect.NewStaticReflector(
		userv1connect.UserServiceName,
	)
	router.Handle(grpcreflect.NewHandlerV1((reflector)))
	router.Handle(userv1connect.NewUserServiceHandler(userv1Server))

	return router
}