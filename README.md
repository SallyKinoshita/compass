# DB 設計・API 実装課題

## 概要

本リポジトリは、ある課題に基づき作成された Web サーバーおよび MySQL データベースの実装です。本プロジェクトでは、Go 言語を使用し、オニオンアーキテクチャを採用しています。

## 構成

- **課題 1**: 表 1 のデータを基にした MySQL のテーブル設計およびデータの取り込み。
- **課題 2**: 提供された API 仕様に基づいた Web サーバーの実装。

## 使用技術

- **言語**: Go
- **データベース**: MySQL
- **インフラ**: Docker / Docker Compose
- **アーキテクチャ**: オニオンアーキテクチャ

## ディレクトリ構成

```plaintext
.
├── cmd
│   └── api                  # アプリケーションのエントリーポイント
│       └── main.go
├── internal
│   ├── application
│   │   └── usecases         # ビジネスロジック
│   ├── domain
│   │   ├── models           # ドメインモデル
│   │   └── repositories     # リポジトリインターフェース
│   ├── infrastructure
│   │   ├── db               # データベース接続設定
│   │   └── persistence      # リポジトリ実装
│   ├── interfaces
│   │   └── controllers      # コントローラー
│   └── gen
│       └── openapi          # OpenAPI生成コード
├── migrations               # マイグレーションファイル
├── docs                     # OpenAPI仕様書
├── docker　　　　　　　　　　  # Docker構成
├── docker-compose.yml       # Docker構成ファイル
├── .env                     # 環境変数ファイル
├── atlas.hcl                # DBスキーマ管理ファイル
├── go.mod                   # Goモジュール設定
└── Makefile                 # 開発タスクのスクリプト
```

## セットアップ

### 必要条件

以下がインストールされている必要があります：

- Docker
- Docker Compose

### 起動手順

以下の手順でプロジェクトをセットアップし、MySQL と Web サーバーを起動します。

```bash
git clone https://github.com/SallyKinoshita/compass.git
cd compass
docker compose up
```

---

## API 仕様

### エンドポイント

- **GET** `/students`

### リクエストパラメータ

| パラメータ名     | 型       | 必須/任意 | デフォルト | 説明                                  |
| ---------------- | -------- | --------- | ---------- | ------------------------------------- |
| `facilitator_id` | `int`    | 必須      | -          | 教師 ID                               |
| `page`           | `int`    | 任意      | `1`        | ページ数                              |
| `limit`          | `int`    | 任意      | `10`       | 1 ページあたりの表示数                |
| `sort`           | `string` | 任意      | `id`       | ソートキー（`id`, `name`, `loginId`） |
| `order`          | `string` | 任意      | `asc`      | ソート順序（`asc`, `desc`）           |

### リクエスト例

```bash
curl 'http://127.0.0.1:48080/students?facilitator_id=1'
```

### レスポンス例

```json
{
  "students": [
    {
      "id": 1,
      "name": "佐藤",
      "loginId": "foo123",
      "classroom": {
        "id": 1,
        "name": "クラスA"
      }
    }
  ],
  "totalCount": 1
}
```

### エラーレスポンス

| 状況                         | ステータスコード  |
| ---------------------------- | ----------------- |
| 該当する生徒が存在しない場合 | `404 Not Found`   |
| リクエストに問題がある場合   | `400 Bad Request` |

---

## テスト

### テストの実行

以下のコマンドでテストを実行します。

```bash
make test
```

## 補足説明

### アーキテクチャ

- **オニオンアーキテクチャ**:
  - **Domain**: 業務ロジックやエンティティを管理。
  - **Application**: ユースケースや DTO を定義。
  - **Infrastructure**: データベースや永続化に関する実装。
  - **Interfaces**: 外部とのインターフェース（HTTP ハンドラー）。

### マイグレーション

- **Atlas** を使用してデータベースのスキーマを管理。
- `migrations` ディレクトリにマイグレーションファイルを格納。
- 今回は起動手順を考慮して seeder ファイルも `migrations` 配下で管理しているが、本来は別で管理する方が保守の観点で良い。

### .env について

- 起動手順を考慮して `.env` の内容をそのままコミットしているが、本来は環境変数の値はコミットすべきでなく secret として別で管理すべきである。
