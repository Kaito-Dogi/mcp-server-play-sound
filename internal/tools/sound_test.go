package tools

import (
	"context"
	"errors"
	"testing"

	"mcp-server-play-sound/internal/types"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// mockPlayer is a mock implementation of platform.SoundPlayer for testing.
type mockPlayer struct {
	playErr      error
	getVolumeErr error
	setVolumeErr error
	currentVolume int
	platformName string
}

func (m *mockPlayer) Play(ctx context.Context, soundFile string, volume int) error {
	return m.playErr
}

func (m *mockPlayer) GetVolume(ctx context.Context) (int, error) {
	if m.getVolumeErr != nil {
		return 0, m.getVolumeErr
	}
	return m.currentVolume, nil
}

func (m *mockPlayer) SetVolume(ctx context.Context, volume int) error {
	if m.setVolumeErr == nil {
		m.currentVolume = volume
	}
	return m.setVolumeErr
}

func (m *mockPlayer) PlatformName() string {
	return m.platformName
}

func TestPlayGlass_Success(t *testing.T) {
	player := &mockPlayer{
		platformName: "darwin",
		currentVolume: 50,
	}
	cfg := types.DefaultConfig()
	tool := NewSoundTool(player, cfg)

	ctx := context.Background()
	req := &mcp.CallToolRequest{}
	input := types.Input{}

	result, output, err := tool.PlayGlass(ctx, req, input)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if output.Status != "played" {
		t.Errorf("Expected status 'played', got: %s", output.Status)
	}
	if result != nil {
		t.Errorf("Expected nil result, got: %v", result)
	}
	// Volume should be restored to 50
	if player.currentVolume != 50 {
		t.Errorf("Expected volume restored to 50, got: %d", player.currentVolume)
	}
}

func TestPlayGlass_UnsupportedPlatform(t *testing.T) {
	player := &mockPlayer{
		platformName: "linux",
	}
	cfg := types.DefaultConfig()
	tool := NewSoundTool(player, cfg)

	ctx := context.Background()
	req := &mcp.CallToolRequest{}
	input := types.Input{}

	_, output, err := tool.PlayGlass(ctx, req, input)

	if err != types.ErrUnsupportedPlatform {
		t.Errorf("Expected ErrUnsupportedPlatform, got: %v", err)
	}
	if output.Status != "unsupported os" {
		t.Errorf("Expected status 'unsupported os', got: %s", output.Status)
	}
}

func TestPlayGlass_PlaybackError(t *testing.T) {
	player := &mockPlayer{
		platformName: "darwin",
		playErr:      errors.New("afplay failed"),
	}
	cfg := types.DefaultConfig()
	tool := NewSoundTool(player, cfg)

	ctx := context.Background()
	req := &mcp.CallToolRequest{}
	input := types.Input{}

	_, output, err := tool.PlayGlass(ctx, req, input)

	if err == nil {
		t.Error("Expected error, got nil")
	}
	if output.Status != "error" {
		t.Errorf("Expected status 'error', got: %s", output.Status)
	}
}

func TestPlayGlass_VolumeRestoreDisabled(t *testing.T) {
	player := &mockPlayer{
		platformName: "darwin",
		currentVolume: 50,
	}
	cfg := &types.Config{
		SoundFile:     "/test/sound.aiff",
		Volume:        75,
		RestoreVolume: false, // Disabled
	}
	tool := NewSoundTool(player, cfg)

	ctx := context.Background()
	req := &mcp.CallToolRequest{}
	input := types.Input{}

	_, output, err := tool.PlayGlass(ctx, req, input)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if output.Status != "played" {
		t.Errorf("Expected status 'played', got: %s", output.Status)
	}
	// Volume should remain at 75 (not restored)
	if player.currentVolume != 75 {
		t.Errorf("Expected volume to remain at 75, got: %d", player.currentVolume)
	}
}

func TestPlayGlass_GetVolumeError(t *testing.T) {
	player := &mockPlayer{
		platformName: "darwin",
		getVolumeErr: errors.New("volume retrieval failed"),
	}
	cfg := types.DefaultConfig()
	tool := NewSoundTool(player, cfg)

	ctx := context.Background()
	req := &mcp.CallToolRequest{}
	input := types.Input{}

	// Should succeed despite volume retrieval error
	_, output, err := tool.PlayGlass(ctx, req, input)

	if err != nil {
		t.Errorf("Expected no error (graceful degradation), got: %v", err)
	}
	if output.Status != "played" {
		t.Errorf("Expected status 'played', got: %s", output.Status)
	}
}

func TestPlayGlass_SetVolumeError(t *testing.T) {
	player := &mockPlayer{
		platformName: "darwin",
		currentVolume: 50,
		setVolumeErr: errors.New("volume setting failed"),
	}
	cfg := types.DefaultConfig()
	tool := NewSoundTool(player, cfg)

	ctx := context.Background()
	req := &mcp.CallToolRequest{}
	input := types.Input{}

	// Should succeed despite volume setting error
	_, output, err := tool.PlayGlass(ctx, req, input)

	if err != nil {
		t.Errorf("Expected no error (graceful degradation), got: %v", err)
	}
	if output.Status != "played" {
		t.Errorf("Expected status 'played', got: %s", output.Status)
	}
}
