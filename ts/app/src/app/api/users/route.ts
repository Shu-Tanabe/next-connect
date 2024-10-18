import { createConnectRpcClient } from "@/services/client";
import { UserService } from "@/services/user/v1/user_connect";
import { ListUsersResponse } from "@/types/user";
import { NextResponse } from "next/server";

// client component で実行する場合に呼び出す handler
export async function GET() {
  const { listUsers } = createConnectRpcClient(UserService);
  const res = await listUsers({});

  return NextResponse.json<ListUsersResponse>(res);
}
