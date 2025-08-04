目的: コードの品質と一貫性を保つためのルールを明記する。
# コーディング規約 (Coding Guidelines)

## 1. 基本方針
* **リーダブルコード:** 他の人が読んで理解しやすいコードを書くことを最優先します。
* **YAGNI (You Ain't Gonna Need It):** 必要になるまで機能や抽象化を実装しません。
* **リンター/フォーマッターに従う:** Prettier と ESLint の設定は、このプロジェクトの「正義」です。

## 2. リンターとフォーマッター

* **Formatter:** Prettier
* **Linter:** ESLint
* **設定ファイル:** `prettier.config.js`, `.eslintrc.js`
* **実行方法:**
```bash
# フォーマットの適用
pnpm format

# lintの実行と自動修正
pnpm lint:fix
```
* VSCode拡張機能を導入し、保存時に自動でフォーマットがかかるように設定することを強く推奨します。

## 3. 命名規則

| 対象         | 規則          | 例                               |
| :----------- | :------------ | :------------------------------- |
| ファイル     | `kebab-case`  | `user-profile.tsx`               |
| 変数・関数   | `camelCase`   | `const userName = ...`           |
| コンポーネント | `PascalCase`  | `function UserProfile() {...}`   |
| 型・インターフェース | `PascalCase`  | `interface UserProfile {...}`    |

## 4. コーディングスタイル

### 非同期処理
* 原則として `async/await` を使用します。 `.then()` や `.catch()` のチェインは、特別な理由がない限り避けてください。

### エラーハンドリング
* `try...catch` ブロックを使用して、エラーが発生しうる箇所を明確に囲んでください。
* API層では、エラー内容をログに出力し、クライアントには汎用的なエラーメッセージを返すように設計します。

## 5. プルリクエスト (PR) のルール

* **ブランチ名:** `feature/issue-123-add-login-page` のように、`[type]/[issue番号]-[概要]` の形式で作成します。
* **PRのタイトル:** `feat: ログイン機能を追加 (#123)` のように、Conventional Commits の形式に沿って記述します。
* **レビュー:** 最低1人以上の `Approve` を得てからマージします。セルフマージは原則禁止です。
