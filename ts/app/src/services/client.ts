import { ServiceType } from "@bufbuild/protobuf";
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-node";

const transport = createConnectTransport({
  baseUrl: "http://localhost:8088",
  httpVersion: "1.1",
});

export function createConnectRpcClient<T extends ServiceType>(service: T) {
  return createClient(service, transport);
}
