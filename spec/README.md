# TypeSpec 仕様書作成プロジェクト
## 概要
このリポジトリは、TypeSpecを使用してOpenAPI仕様書を生成するためのプロジェクトです。

## インストール
1. ツールをインストールします。
   ```bash
   mise install
   ```
2. 依存関係をインストールしてTypeSpecをコンパイルします。
   ```bash
   tsp install
   tsp compile ./src/main.tsp
   ```

## ディレクトリ構成
```pl
..
├── openapi.1.0.yaml  # 生成されたOpenAPI仕様書
├── spec
      ├── package.json
      ├── README.md
      ├── src # TypeSpecソースコードはこのディレクトリに配置
      └── tspconfig.yaml
```
