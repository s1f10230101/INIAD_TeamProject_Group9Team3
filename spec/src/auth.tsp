import "@typespec/http";
import "@typespec/openapi3";
import "@typespec/openapi";
import "@typespec/versioning";
import "./main.tsp";

using TypeSpec.Http;
using OpenAPI;

@tagMetadata("Auth", #{ description: "認証関連のAPI" })

namespace Trip_Canvas;

model JWTAuth {
    @doc("JWTトークン")
    token: string;
}

model UserModel {
    @doc("ユーザーの一意の識別子")
    id: string;

    @doc("ユーザー名")
    username: string;

    @doc("メールアドレス")
    email: string;

    @doc("パスワード")
    @secret
    password: string;
}

model UserResister is OmitProperties<UserModel, "id">;

model UserResponse is OmitProperties<UserModel, "password">;

@tag("Auth")
@route("/auth")
interface AuthAPI {
    @summary("ユーザー登録")
    @doc("新しいユーザーを登録し、JWTトークンを発行します。")
    @operationId("registerUser")
    @route("/register")
    @post
    registerUser(@body newUser: UserResister): JWTAuth | Error<400 | 500>;

    @summary("ユーザーログイン")
    @doc("既存のユーザーがログインし、JWTトークンを取得します。")
    @operationId("loginUser")
    @route("/login")
    @post
    loginUser(@body credentials: { email: string; password: string }): JWTAuth | Error<401 | 500>;

    @summary("ユーザー情報取得")
    @doc("JWTトークンを使用して認証されたユーザーの情報を取得します。")
    @useAuth(Http.BearerAuth)
    @operationId("getUserInfo")
    @route("/me")
    @get
    getUserInfo(@header("Authorization") token: string): UserResponse | Error<401 | 500>;
}
