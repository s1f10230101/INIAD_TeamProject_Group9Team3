import createClient from "openapi-fetch";
import type { paths } from "$lib/types/api";

const client = createClient<paths>({ baseUrl: "http://localhost:8080/v1" }); // バックエンドのURL

export default client;
