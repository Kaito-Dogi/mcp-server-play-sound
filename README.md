# mcp-server-play-sound

音声を再生する MCP サーバーです。

|       項目       |    内容     |
| :--------------: | :---------: |
|     対応 OS      |    macOS    |
|     実行環境     |  ローカル   |
| 推奨クライアント | Claude Code |

## 機能

### 提供ツール

#### `play_glass`

macOS システムサウンド "Glass.aiff" を最大音量で再生します。

**特徴:**
- 現在の音量を自動で保存
- 再生時に音量を最大に設定
- 再生完了後に元の音量を復元
- macOS 専用（darwin）

**パラメータ:** なし

**戻り値:** `{"status": "played"}` または `{"status": "error"}`

## 使用方法

### 1. 実行ファイルのビルド

#### Option A: Makefile を使用（推奨）

```bash
make build
```

#### Option B: go build を直接使用

```bash
go build -o mcp-server-play-sound ./cmd/server
```

### 2. MCP サーバーの登録

#### 2-1. 実行ファイルの実行権限の付与

```bash
chmod +x </path/to/mcp-server-play-sound>
```

#### 2-2. MCP サーバーの登録

```bash
claude mcp add --transport stdio --scope project play-sound -- </path/to/mcp-server-play-sound>

## Added stdio MCP server play-sound with command: </path/to/mcp-server-play-sound>  to project config
## File modified: </path/to/.mcp.json>
```

または `.mcp.json` に直接追加します。

```json
{
  "mcpServers": {
    "play-sound": {
      "type": "stdio",
      "command": "/path/to/mcp-server-play-sound",
      "args": [],
      "env": {}
    }
  }
}
```

詳細は [Claude Code を MCP 経由でツールに接続する - Claude Docs] をご参照ください。

#### 2-3. 登録した MCP サーバーの確認

```bash
claude mcp list

# Checking MCP server health...
#
# play-sound: /Users/k_dogi/Developer/mcp-server/mcp-server-play-sound/mcp-server-play-sound  - ✓ Connected
```

以下のように、登録した直後は MCP サーバーが認識されていない場合があります。<br>
この場合も [3. ツールの呼び出し方法](#3-ツールの呼び出し方法) に進んで OK です。

```bash
## No MCP servers configured. Use `claude mcp add` to add a server.
```

### 3. ツールの呼び出し方法

#### Claude Code から使用

Claude Code を起動して、以下のように依頼します:

```
音声を鳴らして
```

または

```
play_glass ツールを実行して
```

Claude Code が自動的に `play_glass` ツールを呼び出し、Glass サウンドを再生します。

#### プログラムから使用

MCP クライアントライブラリを使用して、`play_glass` ツールを呼び出すことができます。

### 4. 環境変数による設定（Optional）

以下の環境変数で動作をカスタマイズできます:

| 環境変数 | 説明 | デフォルト値 |
|---------|------|------------|
| `MCP_SOUND_FILE` | 再生する音声ファイルのパス | `/System/Library/Sounds/Glass.aiff` |
| `MCP_SOUND_VOLUME` | 再生時の音量 (0-100, -1=最大) | `-1` |
| `MCP_RESTORE_VOLUME` | 音量を復元するか (true/false) | `true` |

**使用例:**

```json
{
  "mcpServers": {
    "play-sound": {
      "type": "stdio",
      "command": "/path/to/mcp-server-play-sound",
      "args": [],
      "env": {
        "MCP_SOUND_FILE": "/System/Library/Sounds/Ping.aiff",
        "MCP_SOUND_VOLUME": "75",
        "MCP_RESTORE_VOLUME": "true"
      }
    }
  }
}
```

## 使用技術

| 項目 |             内容              |
| :--: | :---------------------------: |
| 言語 |              Go               |
| SDK  | [modelcontextprotocol/go-sdk] |

## 開発

### 必要な環境

- Go 1.23.2 以上
- macOS (darwin)

### プロジェクト構造

```
mcp-server-play-sound/
├── cmd/server/           # エントリーポイント
├── internal/
│   ├── config/          # 設定管理
│   ├── platform/        # プラットフォーム固有実装
│   ├── server/          # MCP サーバー初期化
│   ├── tools/           # MCP ツール実装
│   └── types/           # 共通型定義
├── examples/            # サンプルコード
├── tests/               # 統合テスト
├── Makefile            # ビルド自動化
└── README.md
```

### 開発コマンド

```bash
# ビルド
make build

# テスト実行
make test

# カバレッジレポート生成
make coverage

# リント実行（golangci-lint が必要）
make lint

# クリーンアップ
make clean

# ヘルプ表示
make help
```

### テスト

ユニットテストは各パッケージに `*_test.go` ファイルとして配置されています。

```bash
# 全テスト実行
go test ./...

# カバレッジ付きテスト
go test -cover ./...
```

### バージョン管理

バージョンは Git タグベースで管理されます。ビルド時に自動的にバージョン情報が埋め込まれます。

```bash
# バージョン確認
make version
```

## 参考記事

- [VSCode で Go 言語の開発環境を構築する - Qiita]
- [MCP の Go 公式 SDK が公開されたので早速使ってみた - Zenn]
- [Claude Code を MCP 経由でツールに接続する - Claude Docs]

<!-- Links -->

[modelcontextprotocol/go-sdk]: https://github.com/modelcontextprotocol/go-sdk
[VSCode で Go 言語の開発環境を構築する - Qiita]: https://qiita.com/melty_go/items/c977ba594efcffc8b567
[MCP の Go 公式 SDK が公開されたので早速使ってみた - Zenn]: https://zenn.dev/canary_techblog/articles/191579daa04f2a
[Claude Code を MCP 経由でツールに接続する - Claude Docs]: https://docs.claude.com/ja/docs/claude-code/mcp
