目的: アプリケーションのデプロイ方法や、本番環境での運用について記述する。
# 運用ドキュメント (Operations)

# 1. デプロイフロー
まず全体像としては、
1. GCP プロジェクトと API の準備
2. コンテナイメージを Artifact Registry に push
3. Cloud SQL for PostgreSQL インスタンス作成
4. Cloud Run サービス作成（Cloud SQL と接続・環境変数設定）
5. ローカルから Cloud SQL に接続してマイグレーション実行
という流れ

## 0. 事前に決めておく値（例）

```bash
# GCP プロジェクト
PROJECT_ID="iniad-teamproject"

# リージョン（例: 東京）
REGION="asia-northeast1"

# Artifact Registry
REPOSITORY_NAME="iniad-repo"
IMAGE_NAME="go-api"
IMAGE_TAG="v1"
IMAGE_URI="${REGION}-docker.pkg.dev/${PROJECT_ID}/${REPOSITORY_NAME}/${IMAGE_NAME}:${IMAGE_TAG}"

# Cloud Run
SERVICE_NAME="iniad-prg-service"

# Cloud SQL
INSTANCE_NAME="my-postgres"
DB_NAME="my_app_db"
DB_USER="app_user"
DB_PASSWORD="strong-password-here"  # 実際はもっと複雑に

# Cloud SQL インスタンス接続名（後で取得）
# INSTANCE_CONNECTION_NAME="PROJECT_ID:REGION:INSTANCE_NAME"
```

## 1. GCP プロジェクト・ツール準備

### 1-1. gcloud の準備

1. ローカルに以下をインストール
  * Google Cloud SDK (gcloud)
  * Cloud sql proxy
```bash
mise use gcloud cloud-sql-proxy
```
2. ログイン・プロジェクト設定
```bash
gcloud auth login
gcloud config set project ${PROJECT_ID}
gcloud config set run/region ${REGION}
```
3. 課金が有効か確認（コンソール → 課金 → プロジェクトが紐づいているか）

### 1-2. 必要な API を有効化

```bash
gcloud services enable \
  run.googleapis.com \
  sqladmin.googleapis.com \
  artifactregistry.googleapis.com \
  cloudbuild.googleapis.com \
  compute.googleapis.com
```

## 2. Artifact Registry にコンテナイメージを push

### 2-1. Artifact Registry リポジトリ作成

```bash
gcloud artifacts repositories create ${REPOSITORY_NAME} \
  --repository-format=docker \
  --location=${REGION} \
  --description="App container repo"
```

すでに同名リポジトリがある場合はこのステップはスキップします。

### 2-2. Docker から Artifact Registry へのログイン設定

```bash
gcloud auth configure-docker ${REGION}-docker.pkg.dev
```

### 2-3. Docker イメージをビルド
```bash
$ ~/go-backend
docker build -t ${IMAGE_URI} .
```

### 2-4. Artifact Registry へ push
```bash
docker push ${IMAGE_URI}
```

## 3. Cloud SQL for PostgreSQL インスタンス作成

### 3-1. コンソールで操作（クリック手順）
1. GCP コンソール左メニュー → **SQL**
2. 「**インスタンスを作成**」 → **PostgreSQL**
3. 以下を設定
   * インスタンス ID: `my-postgres`（= `${INSTANCE_NAME}`）
   * パスワード: `DB_PASSWORD` で決めた値
   * リージョン: `${REGION}` と同じ (例: asia-northeast1)
   * マシンタイプ・ストレージ: 開発環境なら小さめで OK
4. 「**作成**」をクリック

### 3-2. gcloud コマンドで作成する場合

```bash
gcloud sql instances create ${INSTANCE_NAME} \
  --database-version=POSTGRES_15 \
  --tier=db-f1-micro \
  --region=${REGION}
```

### 3-3. DB ユーザーと DB 作成

#### ユーザー作成

インスタンス作成時にデフォルト `postgres` ユーザができますが、アプリ用ユーザを作っておくとよいです。

```bash
gcloud sql users create ${DB_USER} \
  --instance=${INSTANCE_NAME} \
  --password=${DB_PASSWORD}
```

#### データベース作成

```bash
gcloud sql databases create ${DB_NAME} \
  --instance=${INSTANCE_NAME}
```

### 3-4. インスタンス接続名の取得

```bash
gcloud sql instances describe ${INSTANCE_NAME} \
  --format="value(connectionName)"
```

出力例:

```text
my-gcp-project:asia-northeast1:my-postgres
```

これを `INSTANCE_CONNECTION_NAME` として控えます

## 4. Cloud Run サービスを Cloud SQL に接続してデプロイ
ポイントは：
* Cloud Run サービスに **Cloud SQL Client** 権限を持つサービスアカウントを付与
* `--add-cloudsql-instances` で Cloud SQL を紐づける
* 環境変数 `DB_HOST, POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB` を渡す
  * `DB_HOST` は `/cloudsql/${INSTANCE_CONNECTION_NAME}` を指定（Unix ソケット）([Google Developer forums][3])

### 4-1. Cloud Run 用サービスアカウントに権限付与

デフォルトのコンピュートサービスアカウントを使う想定:
```bash
SA_EMAIL="$(gcloud iam service-accounts list \
  --filter='Compute Engine default service account' \
  --format='value(email)' | head -n1)"

echo ${SA_EMAIL}
# => PROJECT_NUMBER-compute@developer.gserviceaccount.com
```

この SA に Cloud SQL Client ロールを付与:

```bash
gcloud projects add-iam-policy-binding ${PROJECT_ID} \
  --member="serviceAccount:${SA_EMAIL}" \
  --role="roles/cloudsql.client"
```

※必要なら Cloud Run 用に専用のサービスアカウントを作成しても OK です。

### 4-2. gcloud コマンドで Cloud Run にデプロイ

```bash
INSTANCE_CONNECTION_NAME="$(gcloud sql instances describe ${INSTANCE_NAME} \
  --format='value(connectionName)')"

gcloud run deploy ${SERVICE_NAME} \
  --image=${IMAGE_URI} \
  --region=${REGION} \
  --platform=managed \
  --allow-unauthenticated \
  --add-cloudsql-instances=${INSTANCE_CONNECTION_NAME} \
  --set-env-vars=DB_HOST="/cloudsql/${INSTANCE_CONNECTION_NAME}" \
  --set-env-vars=POSTGRES_USER="${DB_USER}" \
  --set-env-vars=POSTGRES_PASSWORD="${DB_PASSWORD}" \
  --set-env-vars=POSTGRES_DB="${DB_NAME}" \
  --set-env-vars=DB_PORT="5432" \
  --service-account=${SA_EMAIL}
```
## 5. ローカルから Cloud SQL に接続してマイグレーション実行
アプリ側からDBのマイグレーションは行わないので、自分のシェルから Cloud SQL に接続してマイグレーションコマンドを実行したい

### 方針

1. ローカルマシンに **Cloud SQL Auth Proxy** を起動
   → ローカルの `localhost:5432` が Cloud SQL PostgreSQL にトンネルされる
2. マイグレーションコマンド
   環境変数 `DB_HOST=127.0.0.1` を読むようにして実行

### 5-2. Cloud SQL Auth Proxy を起動
1. 端末を 1 つ開いて、次のコマンドを実行
```bash
INSTANCE_CONNECTION_NAME="$(gcloud sql instances describe ${INSTANCE_NAME} \
  --format='value(connectionName)')"

cloud-sql-proxy ${INSTANCE_CONNECTION_NAME}
```

正常に起動すると、こんなメッセージが出ます:

```text
Listening on 127.0.0.1:5432 for INSTANCE_CONNECTION_NAME
Ready for new connections
```

> ※ ここではデフォルトで `localhost:5432` にバインドされます。すでにローカルに PostgreSQL が動いている場合はポート番号を変える必要があります（例: `--port 5433`）。

2. このターミナルは **開いたまま** にしておきます（Proxy が動いている状態）。

### 5-3. 別ターミナルで DB にアクセスできるか確認（任意）

PostgreSQL クライアント `psql` が入っている場合:

```bash
psql "host=127.0.0.1 port=5432 sslmode=disable dbname=${DB_NAME} user=${DB_USER}"
```

パスワード入力を求められるので `${DB_PASSWORD}` を入力。
`\q` で抜けられます。

### 5-4. 環境変数を設定してマイグレーションコマンドを実行

#### 例1: ローカル環境にマイグレーションツールが入っている場合

```bash
export DB_HOST="127.0.0.1"
export DB_PORT="5432"
export POSTGRES_USER="${DB_USER}"
export POSTGRES_PASSWORD="${DB_PASSWORD}"
export POSTGRES_DB="${DB_NAME}"

# ここでマイグレーションコマンドを実行
migrate -path ./db/migration -database "postgres://$POSTGRES_USER:$POSTGRES_PASSWORD@$DB_HOST:5432/$POSTGRES_DB" up
```

アプリがこれらの環境変数を読んで DB 接続するようになっていれば、そのまま Cloud SQL に対してマイグレーションが走ります。

# 2. 環境変数

本番環境 (Production) およびステージング環境 (Staging) で設定が必要な環境変数の一覧です。
**注意: ここには実際の値を記述しないでください。**

| 変数名              | 説明                             | 設定場所                  |
| :------------------ | :------------------------------- | :------------------------ |

# 3. 監視とアラート

| 監視項目             | 使用ツール        | アラート条件                             | 通知先                  |
| :------------------- | :---------------- | :--------------------------------------- | :---------------------- |

# 4. トラブルシューティング (FAQ)

### Q. 
**A.**
1.  
