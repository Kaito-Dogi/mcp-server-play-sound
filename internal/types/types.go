package types

// Input represents the input parameters for the play_glass tool.
// Currently, no parameters are required.
type Input struct{}

// Output represents the result of a play_glass tool invocation.
type Output struct {
	// Status indicates the result of the playback operation.
	// Possible values: "played", "error", "unsupported os"
	Status string `json:"status" jsonschema:"playback status"`
}

// Config holds configuration for the sound playback system.
// This allows customization of sound file, volume, and restoration behavior.
type Config struct {
	// SoundFile is the path to the sound file to play.
	// Default: "/System/Library/Sounds/Glass.aiff"
	SoundFile string

	// Volume is the playback volume (0-100).
	// A value of -1 means use maximum volume (100).
	// Default: -1 (maximum)
	Volume int

	// RestoreVolume indicates whether to restore the original volume after playback.
	// Default: true
	RestoreVolume bool
}

// DefaultConfig returns the default configuration.
func DefaultConfig() *Config {
	return &Config{
		SoundFile:     "/System/Library/Sounds/Glass.aiff",
		Volume:        -1, // Maximum volume
		RestoreVolume: true,
	}
}
