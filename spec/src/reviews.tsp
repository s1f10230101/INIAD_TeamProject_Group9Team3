import "@typespec/http";
import "@typespec/openapi3";
import "@typespec/openapi";
import "@typespec/versioning";
import "./main.tsp";

using TypeSpec.Http;
using OpenAPI;

@tagMetadata("Reviews", #{ description: "観光スポットのレビュー関連のAPI" })

namespace Trip_Canvas;

model Review {
    @doc("レビューの一意の識別子")
    @visibility(Lifecycle.Read)
    id: string;

    @doc("レビューを投稿したユーザーの識別子")
    userId: string;

    @doc("レビュー対象の観光スポットの識別子")
    spotId: string;

    @doc("レビューの評価（1から5までの整数）")
    @maxValue(5)
    @minValue(1)
    rating: integer;

    @doc("レビューのコメント")
    comment: string;

    @doc("作成日時")
    @visibility(Lifecycle.Read)
    createdAt: string;
}

@tag("Reviews")
@route("/reviews")
interface ReviewsAPI {
    @summary("観光スポットのレビュー一覧の取得")
    @doc("指定した観光スポットのレビュー一覧を取得します。")
    @operationId("getReviewsBySpotId")
    @get
    getReviewsBySpotId(@query spotId: string): Review[] | Error<500>;

    @summary("レビューの投稿")
    @doc("新しいレビューを投稿します。")
    @operationId("createReview")
    @post
    createReview(@body newReview: Review): Http.CreatedResponse | Error<400 | 500>;
}
