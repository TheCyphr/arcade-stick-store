export async function pingServer() {
  const url = "http://localhost:8080/ping";

  return fetch(url).then(
    (response) => response.json() as Promise<{ value: string }>,
  );
}
