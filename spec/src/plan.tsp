import "@typespec/http";
import "@typespec/openapi3";
import "@typespec/openapi";
import "@typespec/versioning";
import "./main.tsp";

using TypeSpec.Http;
using OpenAPI;

@tagMetadata("Plans", #{ description: "旅行プラン関連のAPI" })
namespace Trip_Canvas;

@mediaTypeHint("text/event-stream")
model Plan {
    @doc("プランの一意の識別子")
    @visibility(Lifecycle.Read)
    id: string;
    @doc("プランの名称")
    name: string;
    @doc("プランの詳細説明")
    description: string;
    @doc("プランに含まれる観光スポットのリスト")
    spots: Spot[];
    @doc("作成日時")
    @visibility(Lifecycle.Read)
    createdAt: string;
}

@tag("Plans")
@route("/plans")
interface PlansAPI {
    @summary("旅行プランの生成")
    @doc("ユーザーの希望に基づいて新しい旅行プランを生成します。")
    @operationId("generatePlan")
    @post
    generatePlan(@body prompt: {prompt: string}): Plan | Error<400 | 500>;
}
