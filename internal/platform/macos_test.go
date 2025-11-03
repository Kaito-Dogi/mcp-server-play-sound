package platform

import (
	"context"
	"testing"
)

func TestMacOSPlayer_PlatformName(t *testing.T) {
	player := NewMacOSPlayer()
	if player.PlatformName() != "darwin" {
		t.Errorf("Expected platform name 'darwin', got: %s", player.PlatformName())
	}
}

func TestMacOSPlayer_SetVolume_InvalidRange(t *testing.T) {
	player := NewMacOSPlayer()
	ctx := context.Background()

	tests := []struct {
		name   string
		volume int
	}{
		{"negative volume", -5},
		{"too high volume", 150},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := player.SetVolume(ctx, tt.volume)
			if err == nil {
				t.Errorf("Expected error for volume %d, got nil", tt.volume)
			}
		})
	}
}

// Note: Integration tests that actually call osascript and afplay are skipped
// in unit tests. They should be in tests/integration/ directory.
