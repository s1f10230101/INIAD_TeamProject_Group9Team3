# =====================================================================
# Base Stage: 依存関係のインストールとキャッシュ
# =====================================================================
FROM node:24-alpine AS base

# pnpmのインストール
RUN npm install -g pnpm

# pnpmが利用するキャッシュディレクトリを作成
ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN mkdir -p $PNPM_HOME

# アプリケーションディレクトリの作成
WORKDIR /app


# =====================================================================
# Dependencies Stage: 依存関係のみをインストールするステージ
# =====================================================================
FROM base AS deps

# package.jsonとpnpm-lock.yamlのみをコピー
COPY package.json pnpm-lock.yaml ./

# --frozen-lockfile: lockファイルとpackage.jsonの乖離を防ぐ
# pnpm fetch: 依存関係のtarballのみをダウンロード（キャッシュに保存）
# ここでキャッシュが生成される
RUN pnpm fetch

# pnpm install --offline: キャッシュから依存関係をインストール
# これにより、ビルド時にネットワークアクセスが発生しない
RUN pnpm install --frozen-lockfile --offline


# =====================================================================
# Builder Stage: 本番用のアプリケーションをビルドするステージ
# =====================================================================
FROM base AS builder

# 依存関係のキャッシュを前のステージからコピー
COPY --from=deps /app/node_modules ./node_modules
COPY --from=deps /pnpm /pnpm

# アプリケーションの全ソースコードをコピー
COPY . .

# SvelteKitアプリケーションをビルド
# 環境変数はここで設定するか、実行時に--build-argで渡す
# ARG PUBLIC_BASE_PATH
# ENV PUBLIC_BASE_PATH=${PUBLIC_BASE_PATH}
RUN pnpm build


# =====================================================================
# Production Stage: 本番用の軽量イメージ
# =====================================================================
FROM base AS production

# ビルド成果物のみをコピー
COPY --from=builder /app/build ./build
COPY --from=builder /app/package.json ./package.json

# 本番用のnode_modulesをコピー (ビルド成果物の実行に必要なもの)
# COPY --from=builder /app/node_modules ./node_modules
# もしくは、本番に必要な依存関係のみを再インストール
# --prodフラグでdevDependenciesを除外
RUN pnpm install --prod --offline

# 実行ポートを開放
EXPOSE 3000

# アプリケーションの実行
# SvelteKitのadapter-nodeが出力するサーバーを実行
CMD [ "node", "build" ]


# =====================================================================
# Development Stage: 開発用のイメージ
# =====================================================================
FROM base AS development

# 依存関係のキャッシュをdepsステージからコピー
COPY --from=deps /app/node_modules ./node_modules
COPY --from=deps /pnpm /pnpm

# アプリケーションの全ソースコードをコピー
COPY . .

# 開発サーバーの実行ポートを開放
EXPOSE 5173

# 開発サーバーを起動
# --hostフラグでコンテナ外からのアクセスを許可
CMD [ "pnpm", "dev", "--host" ]
