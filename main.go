// FIXME: ChatGPT による自動生成コードのため、後ほどリファクタリングする

package main

import (
	"context"
	"errors"
	"log"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type Input struct{}

type Output struct {
	Status string `json:"status" jsonschema:"playback status"`
}

// PlayGlass plays the built‑in macOS system sound "Glass.aiff" using afplay at maximum volume.
// It saves the current volume, sets it to 100%, plays the sound, then restores the original volume.
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

	// Get current volume
	getCurrentVolumeCmd := exec.CommandContext(ctx, "osascript", "-e", "output volume of (get volume settings)")
	currentVolumeBytes, err := getCurrentVolumeCmd.Output()
	if err != nil {
		log.Printf("Warning: failed to get current volume: %v", err)
		// Continue anyway - we'll just skip volume restoration
	}
	currentVolume := strings.TrimSpace(string(currentVolumeBytes))

	// Set volume to maximum (100%)
	setMaxVolumeCmd := exec.CommandContext(ctx, "osascript", "-e", "set volume output volume 100")
	if err := setMaxVolumeCmd.Run(); err != nil {
		log.Printf("Warning: failed to set max volume: %v", err)
		// Continue with current volume
	}

	// Play the sound
	afplay := "/usr/bin/afplay"
	sound := "/System/Library/Sounds/Glass.aiff"
	cmd := exec.CommandContext(ctx, afplay, sound)
	playErr := cmd.Run()

	// Restore original volume
	if currentVolume != "" {
		// Validate that currentVolume is a valid number
		if _, parseErr := strconv.Atoi(currentVolume); parseErr == nil {
			restoreVolumeCmd := exec.CommandContext(ctx, "osascript", "-e", "set volume output volume "+currentVolume)
			if err := restoreVolumeCmd.Run(); err != nil {
				log.Printf("Warning: failed to restore volume to %s: %v", currentVolume, err)
			}
		}
	}

	if playErr != nil {
		return nil, Output{Status: "error"}, playErr
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
