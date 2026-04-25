package components

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/radiodrift/internal/radio"
)

type StationItem struct {
	radio.Station
}

func (i StationItem) Title() string       { return i.Name }
func (i StationItem) Description() string {
	tag := ""
	if len(i.Tags) > 0 {
		tag = i.Tags[0]
	}
	return fmt.Sprintf("%s - %s", i.Country, tag)
}
func (i StationItem) FilterValue() string { return i.Name }

type List struct {
	List list.Model
}

func NewList(stations []radio.Station) List {
	items := make([]list.Item, len(stations))
	for i, s := range stations {
		items[i] = StationItem{Station: s}
	}

	delegate := list.NewDefaultDelegate()
	delegate.ShowDescription = true

	l := list.New(items, delegate, 30, 20)
	l.Title = "Stations"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false) // We'll handle filtering manually with the search component

	return List{List: l}
}

func (l List) Init() tea.Cmd {
	return nil
}

func (l List) Update(msg tea.Msg) (List, tea.Cmd) {
	var cmd tea.Cmd
	l.List, cmd = l.List.Update(msg)
	return l, cmd
}

func (l List) View() string {
	return l.List.View()
}

func (l *List) SetStations(stations []radio.Station) {
	items := make([]list.Item, len(stations))
	for i, s := range stations {
		items[i] = StationItem{Station: s}
	}
	l.List.SetItems(items)
}

func (l List) SelectedItem() *radio.Station {
	item := l.List.SelectedItem()
	if i, ok := item.(StationItem); ok {
		return &i.Station
	}
	return nil
}
