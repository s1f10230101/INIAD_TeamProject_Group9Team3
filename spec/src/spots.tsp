import "@typespec/http";
import "@typespec/openapi3";
import "@typespec/openapi";
import "@typespec/versioning";
import "./main.tsp";

using TypeSpec.Http;
using OpenAPI;

@tagMetadata("Spots", #{ description: "観光スポットに関するAPI" })

namespace Trip_Canvas;

model Spot {
  @doc("観光スポットの一意の識別子")
  @visibility(Lifecycle.Read)
  id: string;

  @doc("観光スポットの名称")
  name: string;

  @doc("観光スポットの説明")
  discription: string;

  @doc("観光スポットの所在地")
  address: string;

  @doc("作成日時")
  @visibility(Lifecycle.Read)
  createdAt: string;
}

@tag("Spots")
@route("/spots")
interface SpotsAPI {
  @summary("観光スポット一覧の取得")
  @doc("登録されている観光スポットの一覧を取得します。")
  @operationId("getAllSpots")
  @get
  spots(): Spot[] | Error<500>;

  @summary("観光スポットの詳細取得")
  @doc("指定したIDの観光スポットの詳細情報を取得します。")
  @operationId("getSpotById")
  @get
  spotById(@path id: string): Spot | Error<404 | 500>;

  @summary("観光スポットの登録")
  @doc("新しい観光スポットを登録します。")
  @operationId("createSpot")
  @post
  createSpot(@body newSpot: Spot): Http.CreatedResponse | Error<400 | 500>;

  @summary("観光スポットの更新")
  @doc("指定したIDの観光スポットの情報を更新します。")
  @operationId("updateSpot")
  @put
  updateSpot(@path id: string, @body updatedSpot: Spot): Spot | Error<400 | 404 | 500>;
}
