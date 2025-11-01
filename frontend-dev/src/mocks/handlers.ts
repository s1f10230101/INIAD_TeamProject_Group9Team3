import { http, HttpResponse } from "msw";

// ここでバックエンドAPIのモックを定義します
export const handlers = [
    // 旅行プラン生成API (/v1/plans) のモック
    http.post("http://localhost:8080/v1/plans", async () => {
        const stream = new ReadableStream({
            async start(controller) {
                const encoder = new TextEncoder();

                // 実際のストリーミングを模倣して、データを少しずつ送る
                const chunk1 = { text: "最高の" };
                controller.enqueue(
                    encoder.encode(`data: ${JSON.stringify(chunk1)}\n\n`),
                );
                await new Promise((r) => setTimeout(r, 10));

                const chunk2 = { text: "東京旅行" };
                controller.enqueue(
                    encoder.encode(`data: ${JSON.stringify(chunk2)}\n\n`),
                );
                await new Promise((r) => setTimeout(r, 10));

                const chunk3 = {
                    text: "のプランです。東京の魅力を満喫する3日間... ",
                };
                controller.enqueue(
                    encoder.encode(`data: ${JSON.stringify(chunk3)}\n\n`),
                );
                await new Promise((r) => setTimeout(r, 10));

                // コンポーネントが期待している終了イベントを送る
                controller.enqueue(encoder.encode("event: done\n\n"));
                controller.close();
            },
        });

        return new HttpResponse(stream, {
            headers: {
                "Content-Type": "text/event-stream",
            },
        });
    }),

    // 他のAPIエンドポイントのモックもここに追加できます
    // http.get(...)
];
