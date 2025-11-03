// FIXME: ChatGPT による自動生成コードのため、後ほどリファクタリングする

package main

import (
	"context"
	"errors"
	"log"
	"os/exec"
	"runtime"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct{}

type Output struct {
	Status string `json:"status" jsonschema:"playback status"`
}

// PlayGlass plays the built‑in macOS system sound "Glass.aiff" using afplay.
func PlayGlass(
	ctx context.Context,
	req *mcp.CallToolRequest,
	_ Input,
) (
	*mcp.CallToolResult,
	Output,
	error,
) {
	if runtime.GOOS != "darwin" {
		return nil, Output{Status: "unsupported os"}, errors.New("PlayGlass only supports macOS (darwin)")
	}

	// Absolute paths to avoid PATH lookup confusion.
	afplay := "/usr/bin/afplay"
	sound := "/System/Library/Sounds/Glass.aiff"

	cmd := exec.CommandContext(ctx, afplay, sound)
	if err := cmd.Run(); err != nil {
		return nil, Output{Status: "error"}, err
	}
	return nil, Output{Status: "played"}, nil
}

func main() {
	log.Println("mcp-server-play-sound: start")

	// Create MCP server
	server := mcp.NewServer(&mcp.Implementation{Name: "mcp-server-play-sound", Version: "v0.0.1"}, nil)
	log.Println("mcp-server-play-sound: created server")

	// Register tools
	mcp.AddTool(server, &mcp.Tool{Name: "play_glass", Description: "Play the macOS system sound Glass.aiff using afplay"}, PlayGlass)
	log.Println("mcp-server-play-sound: registered tools: play_glass")

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}

	log.Println("mcp-server-play-sound: running…")
}
