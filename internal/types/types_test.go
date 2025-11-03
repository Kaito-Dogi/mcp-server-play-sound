package types

import "testing"

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	if cfg.SoundFile != "/System/Library/Sounds/Glass.aiff" {
		t.Errorf("Expected default sound file '/System/Library/Sounds/Glass.aiff', got: %s", cfg.SoundFile)
	}
	if cfg.Volume != -1 {
		t.Errorf("Expected default volume -1, got: %d", cfg.Volume)
	}
	if !cfg.RestoreVolume {
		t.Error("Expected default RestoreVolume to be true")
	}
}
