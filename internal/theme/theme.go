package theme

import "github.com/charmbracelet/lipgloss"

type Theme struct {
	AppFrameStyle        lipgloss.Style
	PaneStyle            lipgloss.Style
	FocusedPaneStyle     lipgloss.Style
	TitleStyle           lipgloss.Style
	SubtleStyle          lipgloss.Style
	BadgeStyle           lipgloss.Style
	SelectedRowStyle     lipgloss.Style
	NowPlayingTitleStyle lipgloss.Style
	StatusBarStyle       lipgloss.Style
	ErrorStyle           lipgloss.Style
}

func DefaultTheme() Theme {
	return Theme{
		AppFrameStyle:        lipgloss.NewStyle().Margin(1, 2),
		PaneStyle:            lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("240")).Padding(1, 2),
		FocusedPaneStyle:     lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("62")).Padding(1, 2),
		TitleStyle:           lipgloss.NewStyle().Foreground(lipgloss.Color("62")).Bold(true),
		SubtleStyle:          lipgloss.NewStyle().Foreground(lipgloss.Color("241")),
		BadgeStyle:           lipgloss.NewStyle().Foreground(lipgloss.Color("0")).Background(lipgloss.Color("62")).Padding(0, 1).MarginRight(1),
		SelectedRowStyle:     lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Background(lipgloss.Color("62")),
		NowPlayingTitleStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true).Underline(true),
		StatusBarStyle:       lipgloss.NewStyle().Foreground(lipgloss.Color("255")).Background(lipgloss.Color("235")).Padding(0, 1),
		ErrorStyle:           lipgloss.NewStyle().Foreground(lipgloss.Color("196")).Bold(true),
	}
}
