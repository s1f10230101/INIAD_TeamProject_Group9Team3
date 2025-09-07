# バックエンドのアーキテクチャ

## 設計原則
1. **スキーマファースト**
  - OpenAPI仕様書をoapi-codegenでスキーマをGoコードに変換
  - sqlcでSQLクエリを型安全に扱う
2. **層化アーキテクチャ**
  - Handler, Usecase, Repositoryの3層に分割
  - 各層はInterfaceを介して疎結合に設計
3. **依存性の逆転**
  - 上位層が下位層に依存しない(下位層を入れ替え可能)ように設計
  - 具体的な実装はmain.goで組み立てる

## 開発手順

### 1. ツールのインストール

このプロジェクトでは `sqlc` と `golang-migrate` を使用します。以下のコマンドでインストールしてください。

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### 2. データベースの起動

開発にはDockerが必要です。以下のコマンドでデータベースを起動します。

```bash
docker compose up -d
```

作業が完了したら、以下のコマンドで停止してください。

```bash
docker compose down
```

### 3. データベースマイグレーション

テーブルスキーマの変更はマイグレーションファイルで行います。

**新しいマイグレーションファイルの作成:**
```bash
# <name> の部分を `create_users_table` のように命名します
migrate create -ext sql -dir db/migration -seq <name>
```

**マイグレーションの適用:**
```bash
# DBを最新の状態にします
migrate -database "postgres://app_user:password@localhost:5432/app_db?sslmode=disable" -path db/migration up
```

**マイグレーションを1つ戻す:**
```bash
migrate -database "postgres://app_user:password@localhost:5432/app_db?sslmode=disable" -path db/migration down 1
```

### 4. SQLCによるコード生成

`db/query/` ディレクトリのSQLファイルを変更した後は、以下のコマンドでGoのコードを再生成する必要があります。

```bash
sqlc generate
```
(このコマンドは `go-backend` ディレクトリで実行してください)

## 使用ライブラリ
- [oapi-codegen]
  - OpenAPI仕様書からGoコードを自動生成するツール
  - 自動生成コードの更新`oapi-codegen --config=api/oapi_config.yaml ../openapi.yaml`
- [air]
  - Go製のホットリロードツール
  - `air -c .air.toml`で起動

## ディレクトリ構成
```plaintext
.
├── api/
│   └── server.gen.go   # 👈 OpenAPI Codegenで自動生成されたコード
├── main.go         # 👈 各パーツを組み立てて起動する役目に特化
├── internal/
│   ├── handler/        # 📦 HTTPリクエスト・レスポンスを扱う層
│   │   └── post.go
│   │   └── post_test.go
│   ├── usecase/        # 🧠 アプリケーションのビジネスロジック層
│   └── repository/     # 🗄️ データの永続化（保存・取得）を担う層
├── go.mod
└── go.sum
```
* **Handler層**
  - HTTPの世界の言葉を話します。リクエストを受け取り、Usecaseに処理を依頼し、結果をレスポンスとして返します。
* **Usecase層** 
  - アプリケーションの本体です。HTTPのことは何も知りません。「投稿を作成する」といった純粋なビジネスロジックを担当します。
* **Repository層**
  - データの保存場所です。Usecaseから依頼を受けて、メモリやデータベースにデータを保存・取得します。  
