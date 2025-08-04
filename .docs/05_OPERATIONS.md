目的: アプリケーションのデプロイ方法や、本番環境での運用について記述する。
# 運用ドキュメント (Operations)

## 1. デプロイフロー

| ブランチ名 | マージ先 | デプロイ先環境 | URL                                | 備考                               |
| :--------- | :------- | :------------- | :--------------------------------- | :--------------------------------- |
| `feature/*`| `develop`| -              | -                                  | レビュー後 `develop` へマージ      |
| `develop`  | `main`   | Staging        | `https://stg.example.com`          | `develop` へのプッシュで自動デプロイ |
| `main`     | -        | Production     | `https://app.example.com`          | `main` へのプッシュで自動デプロイ    |

## 2. 環境変数

本番環境 (Production) およびステージング環境 (Staging) で設定が必要な環境変数の一覧です。
**注意: ここには実際の値を記述しないでください。**

| 変数名              | 説明                             | 設定場所                  |
| :------------------ | :------------------------------- | :------------------------ |
| `DATABASE_URL`      | PostgreSQLの接続文字列           | AWS Secrets Manager       |
| `OPENAI_API_KEY`    | OpenAI APIのシークレットキー     | AWS Secrets Manager       |
| `JWT_SECRET`        | JWTの署名に使用する秘密鍵        | AWS Secrets Manager       |
| `NEXT_PUBLIC_API_URL` | フロントが参照するAPIのエンドポイント | Vercel Environment Variables |

## 3. 監視とアラート

| 監視項目             | 使用ツール        | アラート条件                             | 通知先                  |
| :------------------- | :---------------- | :--------------------------------------- | :---------------------- |
| APIサーバーCPU使用率 | AWS CloudWatch    | 5分間連続で80%以上                     | Slack `#alerts` チャンネル |
| API 5xxエラーレート  | AWS CloudWatch    | 1分間に5回以上発生                     | Slack `#alerts` チャンネル, PagerDuty |
| フロントエンド稼働状況 | Vercel Health Checks | ヘルスチェックが2回連続で失敗          | Email, Slack `#alerts` チャンネル |

## 4. トラブルシューティング (FAQ)

### Q. 5xxエラーが多発している場合、どこを見ればよいか？
**A.**
1.  **Slack `#alerts` チャンネル**で、アラート内容を確認する。
2.  **AWS CloudWatch Logs**にアクセスし、該当時刻の `api-server` のロググループを確認する。
3.  エラーメッセージを元に、原因を特定する。特に `Error:` や `FATAL:` といったキーワードで検索する。

### Q. ローカルでDBに接続できない。
**A.**
1.  `docker ps` コマンドで、PostgreSQLのコンテナが正常に起動しているか確認する。
2.  `packages/api/.env` ファイルの `DATABASE_URL` の設定が正しいか確認する。
3.  `docker-compose logs db` で、コンテナのログにエラーが出ていないか確認する。
