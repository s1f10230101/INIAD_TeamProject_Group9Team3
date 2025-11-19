import createClient from "openapi-fetch";
import type { paths } from "$lib/types/api";
const baseUrl = import.meta.env.PUBLIC_API_BASE_URL || "http://localhost:8080/v1";
//const client = createClient<paths>({ baseUrl: "http://localhost:8080/v1" }); // バックエンドのURL
const client = createClient<paths>({ baseUrl });
/**
 * ストリーム送信を受け取るためのヘルパー関数
 * @param response SSE fetchのResponse
 *
 * @param callback callbackで変数代入をする
 *   ```
 *   let accString = ""
 *   (additionnalText) => accString += additionnalText
 *   ```
 */
export async function streamingRecvHelper(
  response: Response,
  callback: (recvText: string) => void,
) {
  if (response.body) {
    const reader = response.body.getReader();
    const decoder = new TextDecoder();

    while (true) {
      const { done, value } = await reader.read();
      if (done) break;

      const chunk = decoder.decode(value, { stream: true });
      const lines = chunk.split("\n\n");

      for (const line of lines) {
        if (line.startsWith("data:")) {
          const data = line.substring(5).trim();
          const parsed = JSON.parse(data);
          if (parsed.text) {
            callback(parsed.text);
          }
        } else if (line.includes("event: done")) {
          reader.cancel();
          return;
        }
      }
    }
  }
}

export default client;
