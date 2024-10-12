// @generated by protoc-gen-connect-es v1.6.0 with parameter "target=ts"
// @generated from file user/v1/user.proto (package user.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateUserRequest, DeleteUserRequest, GetUserRequest, ListUsersRequest, ListUsersResponse, UpdateUserRequest, User } from "./user_pb.js";
import { Empty, MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service user.v1.UserService
 */
export const UserService = {
  typeName: "user.v1.UserService",
  methods: {
    /**
     * @generated from rpc user.v1.UserService.ListUsers
     */
    listUsers: {
      name: "ListUsers",
      I: ListUsersRequest,
      O: ListUsersResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc user.v1.UserService.GetUser
     */
    getUser: {
      name: "GetUser",
      I: GetUserRequest,
      O: User,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc user.v1.UserService.CreateUser
     */
    createUser: {
      name: "CreateUser",
      I: CreateUserRequest,
      O: User,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc user.v1.UserService.UpdateUser
     */
    updateUser: {
      name: "UpdateUser",
      I: UpdateUserRequest,
      O: User,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc user.v1.UserService.DeleteUser
     */
    deleteUser: {
      name: "DeleteUser",
      I: DeleteUserRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
  }
} as const;

