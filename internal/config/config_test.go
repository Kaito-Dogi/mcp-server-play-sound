package config

import (
	"os"
	"testing"
)

func TestLoadFromEnv_Defaults(t *testing.T) {
	// Clear environment variables
	os.Unsetenv("MCP_SOUND_FILE")
	os.Unsetenv("MCP_SOUND_VOLUME")
	os.Unsetenv("MCP_RESTORE_VOLUME")

	cfg := LoadFromEnv()

	if cfg.SoundFile != "/System/Library/Sounds/Glass.aiff" {
		t.Errorf("Expected default sound file, got: %s", cfg.SoundFile)
	}
	if cfg.Volume != -1 {
		t.Errorf("Expected default volume -1, got: %d", cfg.Volume)
	}
	if !cfg.RestoreVolume {
		t.Error("Expected default RestoreVolume to be true")
	}
}

func TestLoadFromEnv_CustomValues(t *testing.T) {
	// Set environment variables
	os.Setenv("MCP_SOUND_FILE", "/test/custom.aiff")
	os.Setenv("MCP_SOUND_VOLUME", "75")
	os.Setenv("MCP_RESTORE_VOLUME", "false")
	defer func() {
		os.Unsetenv("MCP_SOUND_FILE")
		os.Unsetenv("MCP_SOUND_VOLUME")
		os.Unsetenv("MCP_RESTORE_VOLUME")
	}()

	cfg := LoadFromEnv()

	if cfg.SoundFile != "/test/custom.aiff" {
		t.Errorf("Expected custom sound file '/test/custom.aiff', got: %s", cfg.SoundFile)
	}
	if cfg.Volume != 75 {
		t.Errorf("Expected volume 75, got: %d", cfg.Volume)
	}
	if cfg.RestoreVolume {
		t.Error("Expected RestoreVolume to be false")
	}
}

func TestLoadFromEnv_InvalidValues(t *testing.T) {
	// Set invalid environment variables (should fall back to defaults)
	os.Setenv("MCP_SOUND_VOLUME", "invalid")
	os.Setenv("MCP_RESTORE_VOLUME", "invalid")
	defer func() {
		os.Unsetenv("MCP_SOUND_VOLUME")
		os.Unsetenv("MCP_RESTORE_VOLUME")
	}()

	cfg := LoadFromEnv()

	// Should use defaults when parsing fails
	if cfg.Volume != -1 {
		t.Errorf("Expected default volume -1 for invalid input, got: %d", cfg.Volume)
	}
	if !cfg.RestoreVolume {
		t.Error("Expected default RestoreVolume true for invalid input")
	}
}
