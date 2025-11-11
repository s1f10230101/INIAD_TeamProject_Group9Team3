目的: アプリケーションのデプロイ方法や、本番環境での運用について記述する。
# 運用ドキュメント (Operations)

## 1. デプロイフロー
### 1. Google Artifact Registryにdocker imageを保存
  1. まずgloudをインストールする`mise use gcloud`
  2. googleに認証する`gcloud auth login`
  3. ブラウザで設定したリージョンasia-northeast1に設定`gcloud auth configure-docker asia-northeast1-docker.pkg.dev`
  4. Dokerイメージをgoogleのコンソールに入力した名前やリージョンの設定通りに名前をつけてビルド
  ```bash
  docker build -t "asia-northeast1-docker.pkg.dev/iniadteamprojectgroup9team3/iniad-team-project/iniad-go-api" .
  ```
  5. GCPのArtifactにpush
  ```bash
  docker push "asia-northeast1-docker.pkg.dev/iniadteamprojectgroup9team3/iniad-team-project/iniad-go-api"
  ```
### 2. Cloud Runでimageを起動
  1. ブラウザ上でボタンを見つけて押す

| ブランチ名 | マージ先 | デプロイ先環境 | URL                                | 備考                               |
| :--------- | :------- | :------------- | :--------------------------------- | :--------------------------------- |

## 2. 環境変数

本番環境 (Production) およびステージング環境 (Staging) で設定が必要な環境変数の一覧です。
**注意: ここには実際の値を記述しないでください。**

| 変数名              | 説明                             | 設定場所                  |
| :------------------ | :------------------------------- | :------------------------ |

## 3. 監視とアラート

| 監視項目             | 使用ツール        | アラート条件                             | 通知先                  |
| :------------------- | :---------------- | :--------------------------------------- | :---------------------- |

## 4. トラブルシューティング (FAQ)

### Q. 
**A.**
1.  
