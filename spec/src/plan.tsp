import "@typespec/http";
import "@typespec/openapi3";
import "@typespec/openapi";
import "@typespec/sse";
import "@typespec/events";
import "./main.tsp";

using TypeSpec.Http;
using OpenAPI;

@tagMetadata("Plans", #{ description: "旅行プラン関連のAPI" })
namespace Trip_Canvas;

model Prompt {
    @doc("旅行の希望や条件")
    prompt: string;
}


model PlanResponse extends TypeSpec.SSE.SSEStream<PlanResponseEvent> {
@header `Cache-Control`: "no-cache",
@header `Connection`: "keep-alive",
}
@TypeSpec.Events.events
union PlanResponseEvent {
    aiGeneratedText: TexteventStream;
    done: "event: done\\n\\n";
}
model TexteventStream {
    @Events.contentType("text/plain")
    @Events.data
    texteventStream: string;
}

@tag("Plans")
@route("/plans")
interface PlansAPI {
    @summary("旅行プランの生成")
    @doc("""
        ユーザーの希望に基づいて新しい旅行プランを生成します。"
        AIが生成する旅行プラン
        形式: data: {text: "生成されたテキスト"}\\n\\n
        SSEの形式の末尾が\\n\\nである都合上、もし例えば"data: 生成されたテキスト\\n\\n"のようにしてしまうと、生成されたテキストに\\n\\nが含まれる場合に正しくパースできなくなってしまう。
        そのため、SSE形式の文字列にはJSONエスケープを施すこととし、上記の例では"data: {text: "生成されたテキスト"}\\n\\n"のようにする。
        最後に、生成が完了したことを示すために"event: done\\n\\n"というイベントが送信されます。
        """)
    @operationId("generatePlan")
    @post
    generatePlan(@body prompt: Prompt): PlanResponse | Error<400 | 500>;
}
