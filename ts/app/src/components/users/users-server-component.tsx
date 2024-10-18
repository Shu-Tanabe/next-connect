import { fetchUsers } from "./fetchUsers";

export async function UsersServerComponent() {
  const { users } = await fetchUsers();

  return (
    <div>
      <h3>Users（Server）</h3>
      <ul>
        {users.map((user) => (
          <li key={user.id}>{user.userName}</li>
        ))}
      </ul>
    </div>
  );
}
