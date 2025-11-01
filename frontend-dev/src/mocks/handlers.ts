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
      http.get('http://localhost:8080/v1/spots', () => {
        return HttpResponse.json([
          {
            id: 'c1b5c1c8-0b8f-4b1a-8b1a-0b8f4b1a8b1a',
            name: '東京タワー',
            address: '東京都港区芝公園４丁目２−８',
            description: '東京のシンボル的なタワーです。',
            createdAt: new Date().toISOString(),
          },
          {
            id: 'd2c6d2d9-1c9g-5c2b-9c2b-1c9g5c2b9c2b',
            name: '浅草寺',
            address: '東京都台東区浅草２丁目３−１',
            description: '都内最古の寺院です。',
            createdAt: new Date().toISOString(),
          },
        ]);
      }),
    
      // レビューのモックAPI (動的パス)
      http.get('http://localhost:8080/v1/spots/:spotId/reviews', ({ params }) => {
        const { spotId } = params;
        return HttpResponse.json([
          {
            spotId: spotId,
            userId: 'user-123',
            comment: 'とても良かったです！',
            rating: 5,
            createdAt: new Date().toISOString(),
          },
        ]);
      }),
    ];
