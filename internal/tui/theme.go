package tui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

// Gruvbox Material Dark (soft) palette, pinned as 24-bit hex values so
// no-mistakes renders the same colors regardless of the terminal's palette.
const (
	gruvboxRed         = "#ea6962"
	gruvboxGreen       = "#a9b665"
	gruvboxYellow      = "#d8a657"
	gruvboxBlue        = "#7daea3"
	gruvboxCyan        = "#89b482"
	gruvboxBrightBlack = "#32302f"
)

func init() {
	configureTUIColors()
}

// configureTUIColors forces the TrueColor profile so the pinned gruvbox hex
// colors above emit as 24-bit escapes, keeping the TUI self-contained and
// independent of the terminal's configured ANSI palette.
func configureTUIColors() {
	lipgloss.SetColorProfile(termenv.TrueColor)
}
