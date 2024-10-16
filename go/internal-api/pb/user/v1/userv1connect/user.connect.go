// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: user/v1/user.proto

package userv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/Shu-Tanabe/next-connect/go/internal-api/pb/user/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// UserServiceName is the fully-qualified name of the UserService service.
	UserServiceName = "user.v1.UserService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// UserServiceListUsersProcedure is the fully-qualified name of the UserService's ListUsers RPC.
	UserServiceListUsersProcedure = "/user.v1.UserService/ListUsers"
	// UserServiceGetUserProcedure is the fully-qualified name of the UserService's GetUser RPC.
	UserServiceGetUserProcedure = "/user.v1.UserService/GetUser"
	// UserServiceCreateUserProcedure is the fully-qualified name of the UserService's CreateUser RPC.
	UserServiceCreateUserProcedure = "/user.v1.UserService/CreateUser"
	// UserServiceUpdateUserProcedure is the fully-qualified name of the UserService's UpdateUser RPC.
	UserServiceUpdateUserProcedure = "/user.v1.UserService/UpdateUser"
	// UserServiceDeleteUserProcedure is the fully-qualified name of the UserService's DeleteUser RPC.
	UserServiceDeleteUserProcedure = "/user.v1.UserService/DeleteUser"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	userServiceServiceDescriptor          = v1.File_user_v1_user_proto.Services().ByName("UserService")
	userServiceListUsersMethodDescriptor  = userServiceServiceDescriptor.Methods().ByName("ListUsers")
	userServiceGetUserMethodDescriptor    = userServiceServiceDescriptor.Methods().ByName("GetUser")
	userServiceCreateUserMethodDescriptor = userServiceServiceDescriptor.Methods().ByName("CreateUser")
	userServiceUpdateUserMethodDescriptor = userServiceServiceDescriptor.Methods().ByName("UpdateUser")
	userServiceDeleteUserMethodDescriptor = userServiceServiceDescriptor.Methods().ByName("DeleteUser")
)

// UserServiceClient is a client for the user.v1.UserService service.
type UserServiceClient interface {
	ListUsers(context.Context, *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error)
	GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.User], error)
	CreateUser(context.Context, *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.User], error)
	UpdateUser(context.Context, *connect.Request[v1.UpdateUserRequest]) (*connect.Response[v1.User], error)
	DeleteUser(context.Context, *connect.Request[v1.DeleteUserRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewUserServiceClient constructs a client for the user.v1.UserService service. By default, it uses
// the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewUserServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) UserServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &userServiceClient{
		listUsers: connect.NewClient[v1.ListUsersRequest, v1.ListUsersResponse](
			httpClient,
			baseURL+UserServiceListUsersProcedure,
			connect.WithSchema(userServiceListUsersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getUser: connect.NewClient[v1.GetUserRequest, v1.User](
			httpClient,
			baseURL+UserServiceGetUserProcedure,
			connect.WithSchema(userServiceGetUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		createUser: connect.NewClient[v1.CreateUserRequest, v1.User](
			httpClient,
			baseURL+UserServiceCreateUserProcedure,
			connect.WithSchema(userServiceCreateUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		updateUser: connect.NewClient[v1.UpdateUserRequest, v1.User](
			httpClient,
			baseURL+UserServiceUpdateUserProcedure,
			connect.WithSchema(userServiceUpdateUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		deleteUser: connect.NewClient[v1.DeleteUserRequest, emptypb.Empty](
			httpClient,
			baseURL+UserServiceDeleteUserProcedure,
			connect.WithSchema(userServiceDeleteUserMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// userServiceClient implements UserServiceClient.
type userServiceClient struct {
	listUsers  *connect.Client[v1.ListUsersRequest, v1.ListUsersResponse]
	getUser    *connect.Client[v1.GetUserRequest, v1.User]
	createUser *connect.Client[v1.CreateUserRequest, v1.User]
	updateUser *connect.Client[v1.UpdateUserRequest, v1.User]
	deleteUser *connect.Client[v1.DeleteUserRequest, emptypb.Empty]
}

// ListUsers calls user.v1.UserService.ListUsers.
func (c *userServiceClient) ListUsers(ctx context.Context, req *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error) {
	return c.listUsers.CallUnary(ctx, req)
}

// GetUser calls user.v1.UserService.GetUser.
func (c *userServiceClient) GetUser(ctx context.Context, req *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.User], error) {
	return c.getUser.CallUnary(ctx, req)
}

// CreateUser calls user.v1.UserService.CreateUser.
func (c *userServiceClient) CreateUser(ctx context.Context, req *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.User], error) {
	return c.createUser.CallUnary(ctx, req)
}

// UpdateUser calls user.v1.UserService.UpdateUser.
func (c *userServiceClient) UpdateUser(ctx context.Context, req *connect.Request[v1.UpdateUserRequest]) (*connect.Response[v1.User], error) {
	return c.updateUser.CallUnary(ctx, req)
}

// DeleteUser calls user.v1.UserService.DeleteUser.
func (c *userServiceClient) DeleteUser(ctx context.Context, req *connect.Request[v1.DeleteUserRequest]) (*connect.Response[emptypb.Empty], error) {
	return c.deleteUser.CallUnary(ctx, req)
}

// UserServiceHandler is an implementation of the user.v1.UserService service.
type UserServiceHandler interface {
	ListUsers(context.Context, *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error)
	GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.User], error)
	CreateUser(context.Context, *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.User], error)
	UpdateUser(context.Context, *connect.Request[v1.UpdateUserRequest]) (*connect.Response[v1.User], error)
	DeleteUser(context.Context, *connect.Request[v1.DeleteUserRequest]) (*connect.Response[emptypb.Empty], error)
}

// NewUserServiceHandler builds an HTTP handler from the service implementation. It returns the path
// on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewUserServiceHandler(svc UserServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	userServiceListUsersHandler := connect.NewUnaryHandler(
		UserServiceListUsersProcedure,
		svc.ListUsers,
		connect.WithSchema(userServiceListUsersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceGetUserHandler := connect.NewUnaryHandler(
		UserServiceGetUserProcedure,
		svc.GetUser,
		connect.WithSchema(userServiceGetUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceCreateUserHandler := connect.NewUnaryHandler(
		UserServiceCreateUserProcedure,
		svc.CreateUser,
		connect.WithSchema(userServiceCreateUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceUpdateUserHandler := connect.NewUnaryHandler(
		UserServiceUpdateUserProcedure,
		svc.UpdateUser,
		connect.WithSchema(userServiceUpdateUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	userServiceDeleteUserHandler := connect.NewUnaryHandler(
		UserServiceDeleteUserProcedure,
		svc.DeleteUser,
		connect.WithSchema(userServiceDeleteUserMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/user.v1.UserService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case UserServiceListUsersProcedure:
			userServiceListUsersHandler.ServeHTTP(w, r)
		case UserServiceGetUserProcedure:
			userServiceGetUserHandler.ServeHTTP(w, r)
		case UserServiceCreateUserProcedure:
			userServiceCreateUserHandler.ServeHTTP(w, r)
		case UserServiceUpdateUserProcedure:
			userServiceUpdateUserHandler.ServeHTTP(w, r)
		case UserServiceDeleteUserProcedure:
			userServiceDeleteUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedUserServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedUserServiceHandler struct{}

func (UnimplementedUserServiceHandler) ListUsers(context.Context, *connect.Request[v1.ListUsersRequest]) (*connect.Response[v1.ListUsersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.UserService.ListUsers is not implemented"))
}

func (UnimplementedUserServiceHandler) GetUser(context.Context, *connect.Request[v1.GetUserRequest]) (*connect.Response[v1.User], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.UserService.GetUser is not implemented"))
}

func (UnimplementedUserServiceHandler) CreateUser(context.Context, *connect.Request[v1.CreateUserRequest]) (*connect.Response[v1.User], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.UserService.CreateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) UpdateUser(context.Context, *connect.Request[v1.UpdateUserRequest]) (*connect.Response[v1.User], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.UserService.UpdateUser is not implemented"))
}

func (UnimplementedUserServiceHandler) DeleteUser(context.Context, *connect.Request[v1.DeleteUserRequest]) (*connect.Response[emptypb.Empty], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("user.v1.UserService.DeleteUser is not implemented"))
}
