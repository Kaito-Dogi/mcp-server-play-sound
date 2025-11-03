package main

import (
	"context"
	"fmt"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct {
	Name string `json:"name" jsonschema:"the name of the person to greet"`
}

type Output struct {
	Greeting string `json:"greeting" jsonschema:"the greeting to tell to the user"`
}

func SayHello(
	ctx context.Context,
	req *mcp.CallToolRequest,
	input Input,
) (
	*mcp.CallToolResult,
	Output,
	error,
) {
	return nil, Output{Greeting: "Hello, " + input.Name + "!"}, nil
}

func main() {
	fmt.Println(`あああ: start main`)

	// MCP サーバーを作成
	server := mcp.NewServer(&mcp.Implementation{Name: "mcp-server-play-sound", Version: "v0.0.1"}, nil)

	fmt.Println(`あああ: created a mcp server`)

	// ツールを登録
	mcp.AddTool(server, &mcp.Tool{Name: "greet", Description: "Say Hello"}, SayHello)

	fmt.Println(`あああ: added a tool`)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}

	fmt.Println(`あああ: running mcp server...`)
}
