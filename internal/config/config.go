// Package config provides configuration management for the mcp-server-play-sound.
// It supports loading configuration from environment variables and provides defaults.
package config

import (
	"os"
	"strconv"

	"mcp-server-play-sound/internal/types"
)

// LoadFromEnv loads configuration from environment variables.
// Supported environment variables:
//   - MCP_SOUND_FILE: Path to the sound file (default: /System/Library/Sounds/Glass.aiff)
//   - MCP_SOUND_VOLUME: Playback volume 0-100, -1 for max (default: -1)
//   - MCP_RESTORE_VOLUME: Whether to restore volume after playback (default: true)
func LoadFromEnv() *types.Config {
	cfg := types.DefaultConfig()

	if soundFile := os.Getenv("MCP_SOUND_FILE"); soundFile != "" {
		cfg.SoundFile = soundFile
	}

	if volumeStr := os.Getenv("MCP_SOUND_VOLUME"); volumeStr != "" {
		if volume, err := strconv.Atoi(volumeStr); err == nil {
			cfg.Volume = volume
		}
	}

	if restoreStr := os.Getenv("MCP_RESTORE_VOLUME"); restoreStr != "" {
		if restore, err := strconv.ParseBool(restoreStr); err == nil {
			cfg.RestoreVolume = restore
		}
	}

	return cfg
}
