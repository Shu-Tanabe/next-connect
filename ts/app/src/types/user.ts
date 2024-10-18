import {
  User as CUser,
  ListUsersResponse as CListUsersResponse,
} from "@/services/user/v1/user_pb";

export type User = CUser;

export type ListUsersResponse = CListUsersResponse;
