import { createConnectRpcClient } from "@/services/client";
import { UserService } from "@/services/user/v1/user_connect";
import { ListUsersResponse } from "@/types/user";

export async function fetchUsers(): Promise<ListUsersResponse> {
  const { listUsers } = createConnectRpcClient(UserService);
  const res = await listUsers({});

  return res;
}
