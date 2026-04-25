package app

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/user/radiodrift/internal/player"
	"github.com/user/radiodrift/internal/radio"
	"github.com/user/radiodrift/internal/store"
	"github.com/user/radiodrift/internal/theme"
	"github.com/user/radiodrift/internal/ui/components"
)

type Mode int

const (
	ModeSearch Mode = iota
	ModeList
)

type Model struct {
	theme          theme.Theme
	search         components.Search
	list           components.List
	detail         components.Detail
	allStations    []radio.Station
	mode           Mode
	width          int
	height         int
	player         player.Player
	playingStation *radio.Station
	favorites      map[string]bool
}

func NewModel() Model {
	t := theme.DefaultTheme()
	stations := radio.MockStations()

	favs, _ := store.LoadFavorites()
	if favs == nil {
		favs = make(map[string]bool)
	}

	return Model{
		theme:       t,
		search:      components.NewSearch(),
		list:        components.NewList(stations),
		detail:      components.NewDetail(t),
		allStations: stations,
		mode:        ModeSearch,
		player:      player.NewMPVPlayer(),
		favorites:   favs,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, m.search.Init(), m.list.Init())
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			m.player.Stop()
			return m, tea.Quit
		case "tab":
			if m.mode == ModeSearch {
				m.mode = ModeList
				m.search.Input.Blur()
			} else {
				m.mode = ModeSearch
				m.search.Input.Focus()
			}
			return m, nil
		case "enter":
			if m.mode == ModeList {
				selected := m.list.SelectedItem()
				if selected != nil {
					// Stop playing if it's the same station
					if m.playingStation != nil && m.playingStation.ID == selected.ID {
						m.player.Stop()
						m.playingStation = nil
					} else {
						m.player.Play(selected.URL)
						m.playingStation = selected
					}
				}
			}
		case "f":
			if m.mode == ModeList {
				selected := m.list.SelectedItem()
				if selected != nil {
					m.favorites[selected.ID] = !m.favorites[selected.ID]
					_ = store.SaveFavorites(m.favorites)
				}
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		// Adjust sizes
		h, v := m.theme.AppFrameStyle.GetFrameSize()
		contentWidth := m.width - h
		contentHeight := m.height - v

		// Give list roughly half width, search small height
		listWidth := contentWidth / 2
		m.list.List.SetSize(listWidth-4, contentHeight-6) // Adjust for padding and search box
	}

	if m.mode == ModeSearch {
		var cmd tea.Cmd
		m.search, cmd = m.search.Update(msg)
		cmds = append(cmds, cmd)

		// Filter stations
		query := strings.ToLower(m.search.Input.Value())
		var filtered []radio.Station
		for _, s := range m.allStations {
			if query == "" || strings.Contains(strings.ToLower(s.Name), query) || strings.Contains(strings.ToLower(s.Country), query) || strings.Contains(strings.ToLower(strings.Join(s.Tags, " ")), query) {
				filtered = append(filtered, s)
			}
		}
		m.list.SetStations(filtered)

	} else if m.mode == ModeList {
		var cmd tea.Cmd
		m.list, cmd = m.list.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	if m.width == 0 {
		return "Initializing..."
	}

	// Layout parts
	searchView := m.search.View()
	listView := m.list.View()

	selected := m.list.SelectedItem()
	isPlaying := m.playingStation != nil && selected != nil && m.playingStation.ID == selected.ID
	isFavorite := selected != nil && m.favorites[selected.ID]

	detailView := m.detail.View(selected, isPlaying, isFavorite)

	// Styles
	searchStyle := m.theme.PaneStyle
	if m.mode == ModeSearch {
		searchStyle = m.theme.FocusedPaneStyle
	}

	listStyle := m.theme.PaneStyle
	if m.mode == ModeList {
		listStyle = m.theme.FocusedPaneStyle
	}

	detailStyle := m.theme.PaneStyle

	// Combine
	leftColumn := lipgloss.JoinVertical(lipgloss.Left,
		searchStyle.Render(searchView),
		listStyle.Render(listView),
	)

	rightColumn := detailStyle.Render(detailView)

	content := lipgloss.JoinHorizontal(lipgloss.Top, leftColumn, rightColumn)

	footerStr := "TAB: switch focus • ENTER: play/stop • F: favorite • CTRL+C: quit"
	if m.playingStation != nil {
		footerStr = m.theme.NowPlayingTitleStyle.Render("Now Playing: " + m.playingStation.Name) + " | " + footerStr
	}

	footer := m.theme.StatusBarStyle.Render(footerStr)

	fullView := lipgloss.JoinVertical(lipgloss.Left, content, footer)

	return m.theme.AppFrameStyle.Render(fullView)
}
