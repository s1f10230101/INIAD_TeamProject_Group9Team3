![アプリアイコン](./img/appli-icon.png)<br />
![Dockerアイコン](https://img.shields.io/badge/docker-28.0.0-blue.svg?logo=docker&style=flat)
![Golangアイコン](https://img.shields.io/badge/golang-1.25-blue.svg?logo=go&style=flat)
[![Go Ci](https://github.com/s1f10230101/INIAD_TeamProject_Group9Team3/actions/workflows/go_ci.yaml/badge.svg)](https://github.com/s1f10230101/INIAD_TeamProject_Group9Team3/actions/workflows/go_ci.yaml)

# プロジェクト概要
大学のチーム実習のための演習用Webアプリケーションです。  
詳細は: [OVERVIEW.md](/docs/00_OVERVIEW.md)を参照してください。  
コーディング規約は
[CONTRIBUTING.md](/docs/CONTRIBUTING.md)に記載されています。  
他ドキュメントはすべて`docs`ディレクトリにあります。
- [README.md](/docs/README.md)
- [CONTRIBUTING.md](/docs/CONTRIBUTING.md)
- [00_OVERVIEW.md](/docs/00_OVERVIEW.md)
- [01_ARCHITECTURE.md](/docs/01_ARCHITECTURE.md)
- [02_DOMAIN_KNOWLEDGE.md](/docs/02_DOMAIN_KNOWLEDGE.md)
- [03_DECISION_LOG.md](/docs/03_DECISION_LOG.md)
- [04_OPERATIONS.md](/docs/04_OPERATIONS.md)
# 開発環境のセットアップ手順 (Getting Started)

## 1. 前提条件 (Prerequisites)

| ソフトウェア | 必須バージョン | インストール方法/リンク                                 |
| :--- | :--- | :--- |
| Docker | 28.0.0 以上 | [Docker公式サイト](https://docs.docker.com/get-docker/) |
| Go     | 1.25 以上 | [Go公式サイト](https://go.dev/dl/)                   |

## 2. インストールと起動手順

1.  **リポジトリをクローンする**
    ```bash
    git clone git@github.com:s1f10230101/INIAD_TeamProject_Group9Team3.git
    cd INIAD_TeamProject_Group9Team3
    ```

2. 環境変数ファイル `.env` を作成
    ```bash
    cp go-backend/.env.example go-backend/.env
    ```

3.  **Dockerコンテナのビルドと起動**
    ```bash
    docker compose up --build -d
    ```

## 3. テストの実行方法

```bash
docker compose run --rm backend-dev go test -tags="integration" ./... -v
```
