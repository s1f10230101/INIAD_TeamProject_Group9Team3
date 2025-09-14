目的: システムの全体像と設計思想を伝える。「どのように作られているか」の地図を示す。
# アーキテクチャ

## 1. アーキテクチャ図

旅行プランの提案フロー
```mermaid
sequenceDiagram
    participant ユーザー
    participant Frontend
    participant APIサーバー
    participant AIサービス
    participant Database

    ユーザー ->> Frontend: 希望の旅行体験を入力
    Frontend ->> APIサーバー: APIにリクエスト
    APIサーバー ->> AIサービス: AIに問い合わせ
    AIサービス ->> Database: レビュー情報を取得
    AIサービス ->> Database: データを取得
    AIサービス -->> APIサーバー: 旅行プランを生成
    APIサーバー -->> Frontend: プランを受信
    Frontend -->> ユーザー: 提案を表示
```
観光施設の登録フロー
```mermaid
sequenceDiagram
    participant ユーザー
    participant Frontend
    participant APIサーバー
    participant AIサービス
    participant Database

    ユーザー ->> Frontend: 登録ページを開く
    Frontend ->> APIサーバー: APIにリクエスト
    APIサーバー ->> AIサービス: インタビューAIで質問を生成
    AIサービス -->> APIサーバー: 質問を生成
    APIサーバー -->> Frontend: 質問を受信
    Frontend -->> ユーザー: 質問を表示
    ユーザー ->> Frontend: 質問に回答
    Frontend ->> APIサーバー: 回答を送信
    APIサーバー ->> Database: 施設情報をDBに登録
    Database -->> APIサーバー: データを保存
    APIサーバー -->> Frontend: 登録完了
    Frontend -->> ユーザー: 完了を通知
```
レビューの投稿フロー
```mermaid
sequenceDiagram
    participant ユーザー
    participant Frontend
    participant APIサーバー
    participant AIサービス
    participant Database

    ユーザー ->> Frontend: 施設ページにアクセス
    Frontend ->> APIサーバー: 施設情報を取得
    APIサーバー ->> Database: 情報取得を要求
    Database -->> APIサーバー: データを取得
    APIサーバー -->> Frontend: 情報を受信
    Frontend -->> ユーザー: 情報を表示
    ユーザー ->> Frontend: レビューを入力
    Frontend ->> APIサーバー: レビューを投稿
    APIサーバー ->> Database: レビューをDBに保存
    Database -->> APIサーバー: データを保存
    APIサーバー -->> Frontend: 登録完了
    Frontend -->> ユーザー: 完了を通知
```
図の説明
  - 
## 2. 設計原則 (Design Principles)
  - OpenAPI仕様書に基づくマイクロサービス連携
    - 各サービスはOpenAPI仕様書に従い、openapi-codegen(Golang)やopenapi-typescript(TypeScript)で自動生成されたクライアントを使用して通信。

## 3. 主要コンポーネントの責務

| コンポーネント名 | 責務・役割 |
| --- | --- |
|go-backend | Goで実装されたバックエンドAPI。ユーザー認証、データ管理、ビジネスロジックを担当。 |

## 4. ディレクトリ構造
```
.
├── docs/               # このドキュメント群
├── go-backend/        # Goバックエンドのソースコード
├── .github/            # CI/CD (GitHub Actions) の設定
├── openapi.1.0.yaml  # OpenAPI仕様書
```
