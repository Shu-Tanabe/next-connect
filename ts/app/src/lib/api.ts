export async function fetchData<T>(path: string): Promise<T> {
  const host = process.env.NEXT_PUBLIC_HOST;
  if (!host) {
    throw new Error("Host is not found");
  }

  const apiBase = `${host}`;

  const res = await fetch(new URL(`${apiBase}/api/${path}`));

  const json = await res.json();
  return json;
}
