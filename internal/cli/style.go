package cli

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/kunchenguid/no-mistakes/internal/types"
)

// Gruvbox Material Dark (soft) truecolor palette, matching internal/tui/theme.go.
var (
	sRed    = lipgloss.NewStyle().Foreground(lipgloss.Color("#ea6962"))
	sGreen  = lipgloss.NewStyle().Foreground(lipgloss.Color("#a9b665"))
	sYellow = lipgloss.NewStyle().Foreground(lipgloss.Color("#d8a657"))
	sBlue   = lipgloss.NewStyle().Foreground(lipgloss.Color("#7daea3"))
	sCyan   = lipgloss.NewStyle().Foreground(lipgloss.Color("#89b482"))
	sDim    = lipgloss.NewStyle().Foreground(lipgloss.Color("#32302f"))
	sBold   = lipgloss.NewStyle().Bold(true)
)

// runStatusStyle returns a styled string for the given run status.
func runStatusStyle(status types.RunStatus) string {
	s := string(status)
	switch status {
	case types.RunCompleted:
		return sGreen.Render(s)
	case types.RunFailed:
		return sRed.Render(s)
	case types.RunRunning:
		return sBlue.Render(s)
	case types.RunCIMonitorInterrupted:
		return sYellow.Render(s)
	default:
		return sDim.Render(s)
	}
}
