# API ドキュメント

## MCP ツール仕様

### `play_glass`

macOS システムサウンド "Glass.aiff" を最大音量で再生するツールです。

#### 基本情報

| 項目 | 値 |
|-----|---|
| ツール名 | `play_glass` |
| 説明 | Play the macOS system sound Glass.aiff at maximum volume |
| サポートOS | macOS (darwin) のみ |
| 入力パラメータ | なし |

#### リクエスト

**入力スキーマ:**
```json
{}
```

パラメータは不要です。空のオブジェクトを渡します。

**MCP リクエスト例:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "method": "tools/call",
  "params": {
    "name": "play_glass",
    "arguments": {}
  }
}
```

#### レスポンス

**出力スキーマ:**
```json
{
  "type": "object",
  "properties": {
    "status": {
      "type": "string",
      "description": "playback status"
    }
  }
}
```

**成功時:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "{\"status\":\"played\"}"
      }
    ]
  }
}
```

**失敗時（サポート外OS）:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "{\"status\":\"unsupported os\"}"
      }
    ],
    "isError": true
  }
}
```

**失敗時（再生エラー）:**
```json
{
  "jsonrpc": "2.0",
  "id": 1,
  "result": {
    "content": [
      {
        "type": "text",
        "text": "{\"status\":\"error\"}"
      }
    ],
    "isError": true
  }
}
```

#### 動作仕様

1. **プラットフォームチェック**
   - `runtime.GOOS` が "darwin" でない場合、エラーを返却
   - ステータス: `"unsupported os"`

2. **音量の取得**
   - `osascript -e "output volume of (get volume settings)"` を実行
   - 失敗してもエラーとせず、警告ログを出力して続行

3. **音量の最大化**
   - デフォルト: `osascript -e "set volume output volume 100"` を実行
   - 環境変数 `MCP_SOUND_VOLUME` で音量を変更可能（0-100、-1=最大）
   - 失敗してもエラーとせず、警告ログを出力して続行

4. **音声の再生**
   - デフォルト: `/usr/bin/afplay /System/Library/Sounds/Glass.aiff` を実行
   - 環境変数 `MCP_SOUND_FILE` でファイルパスを変更可能
   - **失敗した場合はエラーを返却** (ステータス: `"error"`)

5. **音量の復元**
   - ステップ2で取得した音量に戻す
   - デフォルトで有効（環境変数 `MCP_RESTORE_VOLUME=false` で無効化可能）
   - 失敗してもエラーとせず、警告ログを出力（処理は成功扱い）

#### 環境変数

| 環境変数 | 型 | デフォルト値 | 説明 |
|---------|---|------------|------|
| `MCP_SOUND_FILE` | string | `/System/Library/Sounds/Glass.aiff` | 再生する音声ファイルのパス |
| `MCP_SOUND_VOLUME` | int | `-1` | 再生時の音量 (0-100, -1=最大) |
| `MCP_RESTORE_VOLUME` | bool | `true` | 音量を元に戻すかどうか |

#### エラーコード

| ステータス | 説明 | 対処方法 |
|----------|------|---------|
| `played` | 成功 | - |
| `unsupported os` | macOS 以外の OS で実行 | macOS で実行してください |
| `error` | 音声再生失敗 | afplay が利用可能か確認してください |

#### 使用例

**Claude Code から:**
```
音声を鳴らして
```

**カスタム設定で使用:**

`.mcp.json`:
```json
{
  "mcpServers": {
    "play-sound": {
      "type": "stdio",
      "command": "/path/to/mcp-server-play-sound",
      "env": {
        "MCP_SOUND_FILE": "/System/Library/Sounds/Ping.aiff",
        "MCP_SOUND_VOLUME": "75"
      }
    }
  }
}
```

#### 制限事項

1. **同期実行**
   - 音声の再生は同期的に実行され、再生完了まで待機します
   - 長い音声ファイルの場合、ツール呼び出しがブロックされます

2. **単一音声**
   - 一度に1つの音声しか再生できません
   - 複数の音声を同時に再生する機能はありません

3. **macOS 専用**
   - 現在は macOS (darwin) のみサポート
   - Linux/Windows では動作しません

#### ログ出力

実行時に以下のログが出力されます:

```
2025/11/03 19:44:51 mcp-server-play-sound: starting...
2025/11/03 19:44:51 mcp-server-play-sound: initializing server
2025/11/03 19:44:51 mcp-server-play-sound: created server (version: v0.0.1)
2025/11/03 19:44:51 mcp-server-play-sound: loaded config (sound: /System/Library/Sounds/Glass.aiff, volume: -1, restore: true)
2025/11/03 19:44:51 mcp-server-play-sound: using macOS player
2025/11/03 19:44:51 mcp-server-play-sound: registered tool: play_glass
2025/11/03 19:44:51 mcp-server-play-sound: starting server...
```

音量制御で問題が発生した場合:
```
2025/11/03 19:44:51 Warning: failed to get current volume: <error>
2025/11/03 19:44:51 Warning: failed to set volume to 100: <error>
2025/11/03 19:44:51 Warning: failed to restore volume to 50: <error>
```

## 将来のツール（計画）

### `play_sound` (将来実装予定)

任意の音声ファイルを再生するツール。

**パラメータ:**
```json
{
  "sound_file": "/path/to/sound.aiff",
  "volume": 75
}
```

### `list_sounds` (将来実装予定)

利用可能なシステムサウンドの一覧を取得するツール。

**パラメータ:** なし

**戻り値:**
```json
{
  "sounds": [
    "/System/Library/Sounds/Glass.aiff",
    "/System/Library/Sounds/Ping.aiff",
    ...
  ]
}
```
