// Package platform provides platform-specific sound playback implementations.
// It defines the SoundPlayer interface and provides implementations for different operating systems.
package platform

import (
	"context"
)

// SoundPlayer defines the interface for platform-specific sound playback operations.
// Different operating systems should implement this interface to provide
// volume control and sound playback capabilities.
type SoundPlayer interface {
	// Play plays a sound file at the specified volume (0-100).
	// A volume of -1 indicates maximum volume.
	// Returns an error if playback fails.
	Play(ctx context.Context, soundFile string, volume int) error

	// GetVolume returns the current system volume (0-100).
	// Returns an error if volume retrieval fails.
	GetVolume(ctx context.Context) (int, error)

	// SetVolume sets the system volume (0-100).
	// Returns an error if volume setting fails.
	SetVolume(ctx context.Context, volume int) error

	// PlatformName returns the name of the platform (e.g., "darwin", "linux", "windows").
	PlatformName() string
}
