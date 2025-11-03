// Package types defines common types and errors used across the mcp-server-play-sound project.
package types

import "errors"

// Sentinel errors for categorizing failures in the sound playback system.
var (
	// ErrUnsupportedPlatform indicates the current OS platform is not supported.
	// Currently, only macOS (darwin) is supported.
	ErrUnsupportedPlatform = errors.New("unsupported platform")

	// ErrVolumeControl indicates a failure to get or set system volume.
	// This may occur if osascript is unavailable or permissions are insufficient.
	ErrVolumeControl = errors.New("volume control failed")

	// ErrSoundPlayback indicates a failure during sound playback.
	// This may occur if afplay is unavailable or the sound file doesn't exist.
	ErrSoundPlayback = errors.New("sound playback failed")
)
