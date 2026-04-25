package components

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Search struct {
	Input textinput.Model
}

func NewSearch() Search {
	ti := textinput.New()
	ti.Placeholder = "Search for a station..."
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 30

	return Search{
		Input: ti,
	}
}

func (s Search) Init() tea.Cmd {
	return textinput.Blink
}

func (s Search) Update(msg tea.Msg) (Search, tea.Cmd) {
	var cmd tea.Cmd
	s.Input, cmd = s.Input.Update(msg)
	return s, cmd
}

func (s Search) View() string {
	return s.Input.View()
}
