import "@typespec/http";
import "@typespec/openapi3";
import "@typespec/openapi";
import "@typespec/versioning";
import "./main.tsp";

using TypeSpec.Http;
using OpenAPI;

@tagMetadata("Plans", #{ description: "旅行プラン関連のAPI" })
namespace Trip_Canvas;

model Prompt {
    @doc("旅行の希望や条件")
    prompt: string;
}

@mediaTypeHint("text/event-stream")
@encodedName("text/event-stream", "txt")
model PlanResponse {
    @doc("プランの名称")
    name: string;
    @doc("プランの詳細説明")
    description: string;
    @doc("プランに含まれる観光スポットのリスト")
    spots: SpotResponse[];
}

@tag("Plans")
@route("/plans")
interface PlansAPI {
    @summary("旅行プランの生成")
    @doc("ユーザーの希望に基づいて新しい旅行プランを生成します。")
    @operationId("generatePlan")
    @post
    generatePlan(@body prompt: Prompt): PlanResponse | Error<400 | 500>;
}
