// Package server provides MCP server initialization and tool registration.
package server

import (
	"context"
	"log"
	"runtime"

	"mcp-server-play-sound/internal/config"
	"mcp-server-play-sound/internal/platform"
	"mcp-server-play-sound/internal/tools"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Version is the current version of the server.
// This will be overridden at build time using ldflags.
var Version = "v0.0.1-dev"

// New creates and configures a new MCP server with the play_glass tool registered.
// It automatically detects the platform and uses the appropriate sound player.
// Configuration is loaded from environment variables.
func New() *mcp.Server {
	log.Println("mcp-server-play-sound: initializing server")

	// Create MCP server
	server := mcp.NewServer(&mcp.Implementation{
		Name:    "mcp-server-play-sound",
		Version: Version,
	}, nil)
	log.Printf("mcp-server-play-sound: created server (version: %s)", Version)

	// Load configuration
	cfg := config.LoadFromEnv()
	log.Printf("mcp-server-play-sound: loaded config (sound: %s, volume: %d, restore: %v)",
		cfg.SoundFile, cfg.Volume, cfg.RestoreVolume)

	// Create platform-specific player
	var player platform.SoundPlayer
	switch runtime.GOOS {
	case "darwin":
		player = platform.NewMacOSPlayer()
		log.Println("mcp-server-play-sound: using macOS player")
	default:
		log.Printf("mcp-server-play-sound: WARNING - unsupported platform: %s", runtime.GOOS)
		// Still create a player that will return errors
		player = platform.NewMacOSPlayer() // Will fail on non-darwin
	}

	// Create tool handler
	soundTool := tools.NewSoundTool(player, cfg)

	// Register tools
	mcp.AddTool(
		server,
		&mcp.Tool{
			Name:        "play_glass",
			Description: "Play the macOS system sound Glass.aiff at maximum volume",
		},
		soundTool.PlayGlass,
	)
	log.Println("mcp-server-play-sound: registered tool: play_glass")

	return server
}

// Run starts the MCP server with stdio transport.
func Run(ctx context.Context, server *mcp.Server) error {
	log.Println("mcp-server-play-sound: starting server...")
	if err := server.Run(ctx, &mcp.StdioTransport{}); err != nil {
		return err
	}
	log.Println("mcp-server-play-sound: server stopped")
	return nil
}
