import "@typespec/http";
import "@typespec/openapi3";
import "@typespec/openapi";
import "@typespec/versioning";
import "./main.tsp";

using TypeSpec.Http;
using OpenAPI;

@tagMetadata("Reviews", #{ description: "観光スポットのレビュー関連のAPI" })

namespace Trip_Canvas;

model ReviewModel {
    @doc("レビューを投稿したユーザーの識別子")
    userId: Uuid;

    @doc("レビュー対象の観光スポットの識別子")
    spotId: Uuid;

    @doc("レビューの評価（1から5までの整数）")
    @maxValue(5)
    @minValue(1)
    rating: integer;

    @doc("レビューのコメント")
    comment: string;

    @doc("レビューの作成日時")
    createdAt: utcDateTime;
}

model ReviewResister is OmitProperties<ReviewModel, "userId" | "createdAt">;

model ReviewResponse is ReviewModel;

@tag("Spots")
@tag("Reviews")
@route("/spots")
interface ReviewsAPI {
    @summary("観光スポットのレビュー一覧の取得")
    @doc("指定した観光スポットのレビュー一覧を取得します。")
    @operationId("getReviewsBySpotId")
    @route("{spotId}/reviews")
    @get
    getReviewsBySpotId(@path spotId: Uuid): ReviewResponse[] | Error<500>;

    @summary("レビューの投稿")
    @doc("新しいレビューを投稿します。")
    @operationId("createReview")
    @route("{spotId}/reviews")
    @post
    createReview(@path spotId: Uuid, @body newReview: ReviewResister): Success<201, ReviewResponse> | Error<400 | 500>;
}
