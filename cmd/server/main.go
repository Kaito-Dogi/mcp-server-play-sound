// mcp-server-play-sound is an MCP server that provides sound playback functionality.
// It currently supports playing macOS system sounds at maximum volume.
package main

import (
	"context"
	"log"

	"mcp-server-play-sound/internal/server"
)

func main() {
	log.Println("mcp-server-play-sound: starting...")

	// Create and run server
	srv := server.New()
	if err := server.Run(context.Background(), srv); err != nil {
		log.Fatal(err)
	}
}
