import {http, HttpResponse} from 'msw';

export const handlers = [
    http.post('http://localhost:8080/v1/plans', async () => {
        const stream = new ReadableStream({
            start(controller) {
                const text = '{"id":"plan-1","title":"最高の東京旅行","description":"東京の魅力を満喫する3日間のプランです。"}';
                const encoder = new TextEncoder();
                controller.enqueue(encoder.encode(`data: ${text}\n\n`));
                controller.close();
            },
        });

        return new HttpResponse(stream, {
            headers: {
                'Content-Type': 'text/event-stream',
            },
        });
    }),

    
]