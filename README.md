# プロジェクト概要
大学のチーム実習のための演習用Webアプリケーションです。
詳細は: [OVERVIEW.md](./.docs/00_OVERVIEW.md)を参照してください。

# 開発環境のセットアップ手順 (Getting Started)

## 1. 前提条件 (Prerequisites)

| ソフトウェア | 必須バージョン | インストール方法/リンク                                 |
| :----------- | :------------- | :------------------------------------------------------ |
| Node.js      | `v20.x`        | `nvm` や `volta` での管理を推奨: [https://nodejs.org/](https://nodejs.org/) |
| Docker       | `v24.x`        | [https://www.docker.com/products/docker-desktop/](https://www.docker.com/products/docker-desktop/) |
| pnpm         | `v9.x`         | `npm install -g pnpm`                                   |

## 2. インストールと起動手順

1.  **リポジトリをクローンする**
    ```bash
    git clone [リポジトリのURL]
    cd [リポジトリ名]
    ```

2.  **依存パッケージをインストールする**
    ```bash
    pnpm install
    ```

3.  **環境変数を設定する**
    * `packages/api/` ディレクトリにある `.env.example` をコピーして `.env` ファイルを作成します。
    ```bash
    cp packages/api/.env.example packages/api/.env
    ```
    * 作成した `.env` ファイルを開き、必要な値を設定します。（例: `DATABASE_URL`, `OPENAI_API_KEY`）

4.  **Dockerコンテナ（DBなど）を起動する**
    ```bash
    docker-compose up -d
    ```

5.  **データベースをマイグレーションする**
    ```bash
    pnpm --filter api run db:migrate
    ```

6.  **開発サーバーを起動する**
    ```bash
    pnpm dev
    ```
    * フロントエンド: `http://localhost:3000`
    * バックエンドAPI: `http://localhost:8000`

## 3. テストの実行方法

```bash
# すべてのテストを実行
pnpm test

# APIのテストのみ実行
pnpm --filter api test
```
