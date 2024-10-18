"use client";
import { ListUsersResponse, User } from "@/types/user";
import { fetchData } from "@/lib/api";
import { useEffect, useState } from "react";

export function UsersClientComponent() {
  const [users, setUsers] = useState<User[] | null>(null);

  useEffect(() => {
    async function fetchUsers() {
      const { users } = await fetchData<ListUsersResponse>("users");
      setUsers(users);
    }
    fetchUsers();
  }, []);

  if (users === null) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h3>Users（Client）</h3>
      <ul>
        {users.map((user) => (
          <li key={user.id}>{user.userName}</li>
        ))}
      </ul>
    </div>
  );
}
