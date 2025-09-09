import "@typespec/http";
import "@typespec/openapi3";
import "@typespec/openapi";
import "@typespec/versioning";

import "./spots.tsp";
import "./reviews.tsp";
import "./plan.tsp";
import "./auth.tsp";

using TypeSpec.Http;
using OpenAPI;
using Versioning;

@service(#{ title: "AI Travel Planner API" })
@doc("AI旅行プラン提案サービスのAPI仕様")
@info(#{ summary: "AI旅行プラン提案サービスのAPI仕様", })
@server("https://localhost:8080/v1", "ローカル開発用サーバー")

@versioned(Versions)
namespace Trip_Canvas;
enum Versions {
  v1: "1.0",
}

@encode("uuid")
scalar Uuid extends string;

model Success<StatusCode extends integer, T> {
  @statusCode _: StatusCode;
  @bodyRoot body: T;
}

// 400: 不正なリクエスト
// 401: 認証エラー
// 403: 権限エラー
// 404: リソースが見つからない
// 500: サーバーエラー
model Error<StatusCode extends integer> {
  @statusCode _: StatusCode;
  message: string;
}
