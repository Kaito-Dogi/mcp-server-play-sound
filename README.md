# mcp-server-play-sound

音声を再生する MCP サーバーです。

|       項目       |    内容     |
| :--------------: | :---------: |
|     対応 OS      |    macOS    |
|     実行環境     |  ローカル   |
| 推奨クライアント | Claude Code |

## 使用方法

### 1. 実行ファイルのビルド

```bash
go build
```

### 2. MCP サーバーの登録

#### 2-1. 実行ファイルの実行権限の付与

```bash
chmod +x </path/to/mcp-server-play-sound>
```

#### 2-2. MCP サーバーの登録

```bash
claude mcp add --transport stdio --scope project play-sound -- </path/to/mcp-server-play-sound>
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

登録した直後は表示されない可能性があります。

### 3. ツールの呼び出し方法

TBD

## 使用技術

| 項目 |             内容              |
| :--: | :---------------------------: |
| 言語 |              Go               |
| SDK  | [modelcontextprotocol/go-sdk] |

## 環境構築

TBD

## 参考記事

- [VSCode で Go 言語の開発環境を構築する - Qiita]
- [MCP の Go 公式 SDK が公開されたので早速使ってみた - Zenn]
- [Claude Code を MCP 経由でツールに接続する - Claude Docs]

<!-- Links -->

[modelcontextprotocol/go-sdk]: https://github.com/modelcontextprotocol/go-sdk
[VSCode で Go 言語の開発環境を構築する - Qiita]: https://qiita.com/melty_go/items/c977ba594efcffc8b567
[MCP の Go 公式 SDK が公開されたので早速使ってみた - Zenn]: https://zenn.dev/canary_techblog/articles/191579daa04f2a
[Claude Code を MCP 経由でツールに接続する - Claude Docs]: https://docs.claude.com/ja/docs/claude-code/mcp
