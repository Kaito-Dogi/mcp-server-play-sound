package platform

import (
	"context"
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"mcp-server-play-sound/internal/types"
)

// MacOSPlayer implements the SoundPlayer interface for macOS.
// It uses osascript for volume control and afplay for sound playback.
type MacOSPlayer struct{}

// NewMacOSPlayer creates a new macOS sound player.
func NewMacOSPlayer() *MacOSPlayer {
	return &MacOSPlayer{}
}

// Play plays a sound file at the specified volume on macOS.
// It uses afplay to play the sound file synchronously.
func (m *MacOSPlayer) Play(ctx context.Context, soundFile string, volume int) error {
	afplay := "/usr/bin/afplay"
	cmd := exec.CommandContext(ctx, afplay, soundFile)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w: %v", types.ErrSoundPlayback, err)
	}
	return nil
}

// GetVolume returns the current system volume on macOS (0-100).
// It uses osascript to query the system volume settings.
func (m *MacOSPlayer) GetVolume(ctx context.Context) (int, error) {
	cmd := exec.CommandContext(ctx, "osascript", "-e", "output volume of (get volume settings)")
	output, err := cmd.Output()
	if err != nil {
		return 0, fmt.Errorf("%w: failed to get volume: %v", types.ErrVolumeControl, err)
	}

	volumeStr := strings.TrimSpace(string(output))
	volume, err := strconv.Atoi(volumeStr)
	if err != nil {
		return 0, fmt.Errorf("%w: invalid volume value '%s': %v", types.ErrVolumeControl, volumeStr, err)
	}

	return volume, nil
}

// SetVolume sets the system volume on macOS (0-100).
// It uses osascript to set the output volume.
func (m *MacOSPlayer) SetVolume(ctx context.Context, volume int) error {
	if volume < 0 || volume > 100 {
		return fmt.Errorf("%w: volume must be between 0 and 100, got %d", types.ErrVolumeControl, volume)
	}

	cmd := exec.CommandContext(ctx, "osascript", "-e", fmt.Sprintf("set volume output volume %d", volume))
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%w: failed to set volume to %d: %v", types.ErrVolumeControl, volume, err)
	}

	return nil
}

// PlatformName returns "darwin" for macOS.
func (m *MacOSPlayer) PlatformName() string {
	return "darwin"
}
