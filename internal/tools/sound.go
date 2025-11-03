// Package tools implements MCP tools for sound playback operations.
package tools

import (
	"context"
	"log"

	"mcp-server-play-sound/internal/platform"
	"mcp-server-play-sound/internal/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// SoundTool handles sound playback operations using a platform-specific player.
type SoundTool struct {
	player platform.SoundPlayer
	config *types.Config
}

// NewSoundTool creates a new SoundTool with the given player and configuration.
func NewSoundTool(player platform.SoundPlayer, config *types.Config) *SoundTool {
	return &SoundTool{
		player: player,
		config: config,
	}
}

// PlayGlass plays the configured sound file at maximum volume.
// It saves the current volume, sets it to the configured volume, plays the sound,
// then restores the original volume if configured to do so.
//
// Returns:
//   - CallToolResult: nil (no additional result needed)
//   - Output: Status indicating "played", "error", or "unsupported os"
//   - error: Any error that occurred during playback
func (st *SoundTool) PlayGlass(
	ctx context.Context,
	req *mcp.CallToolRequest,
	_ types.Input,
) (
	*mcp.CallToolResult,
	types.Output,
	error,
) {
	// Validate platform
	if st.player.PlatformName() != "darwin" {
		return nil, types.Output{Status: "unsupported os"}, types.ErrUnsupportedPlatform
	}

	// Get current volume if we need to restore it
	var originalVolume int
	var volumeRetrieved bool
	if st.config.RestoreVolume {
		vol, err := st.player.GetVolume(ctx)
		if err != nil {
			log.Printf("Warning: failed to get current volume: %v", err)
			// Continue anyway - we'll just skip volume restoration
		} else {
			originalVolume = vol
			volumeRetrieved = true
		}
	}

	// Set volume to configured level
	targetVolume := st.config.Volume
	if targetVolume == -1 {
		targetVolume = 100 // Maximum volume
	}
	if err := st.player.SetVolume(ctx, targetVolume); err != nil {
		log.Printf("Warning: failed to set volume to %d: %v", targetVolume, err)
		// Continue with current volume
	}

	// Play the sound
	playErr := st.player.Play(ctx, st.config.SoundFile, targetVolume)

	// Restore original volume if configured and retrieved
	if st.config.RestoreVolume && volumeRetrieved {
		if err := st.player.SetVolume(ctx, originalVolume); err != nil {
			log.Printf("Warning: failed to restore volume to %d: %v", originalVolume, err)
		}
	}

	// Check for playback errors
	if playErr != nil {
		return nil, types.Output{Status: "error"}, playErr
	}

	return nil, types.Output{Status: "played"}, nil
}
