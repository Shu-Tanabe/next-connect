import { UsersClientComponent } from "@/components/users/users-client-component";
import { UsersServerComponent } from "@/components/users/users-server-component";
import { Suspense } from "react";

export default async function Home() {
  return (
    <div>
      <h2>Server Component</h2>
      <Suspense fallback={<div>Loading...</div>}>
        <UsersServerComponent />
      </Suspense>
      <h2>Client Component</h2>
      <UsersClientComponent />
    </div>
  );
}
