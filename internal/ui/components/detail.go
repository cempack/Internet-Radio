package components

import (
	"fmt"
	"strings"

	"github.com/user/radiodrift/internal/radio"
	"github.com/user/radiodrift/internal/theme"
)

type Detail struct {
	Theme theme.Theme
}

func NewDetail(t theme.Theme) Detail {
	return Detail{Theme: t}
}

func (d Detail) View(station *radio.Station, isPlaying bool, isFavorite bool) string {
	if station == nil {
		return d.Theme.SubtleStyle.Render("Select a station to view details.")
	}

	var sb strings.Builder

	titleStr := station.Name
	if isPlaying {
		titleStr = d.Theme.NowPlayingTitleStyle.Render("▶ " + station.Name)
	} else {
		titleStr = d.Theme.TitleStyle.Render(station.Name)
	}
	sb.WriteString(titleStr + "\n\n")

	sb.WriteString(fmt.Sprintf("Location: %s\n", station.Country))
	sb.WriteString(fmt.Sprintf("Language: %s\n", station.Language))

	tagsStr := ""
	for _, tag := range station.Tags {
		tagsStr += d.Theme.BadgeStyle.Render(tag) + " "
	}
	sb.WriteString(fmt.Sprintf("Tags: %s\n\n", tagsStr))

	sb.WriteString(fmt.Sprintf("Codec: %s @ %dkbps\n", station.Codec, station.Bitrate))
	sb.WriteString(fmt.Sprintf("Votes: %d\n\n", station.Votes))

	if isFavorite {
		sb.WriteString(d.Theme.TitleStyle.Render("★ Favorited"))
	} else {
		sb.WriteString(d.Theme.SubtleStyle.Render("☆ Not Favorited"))
	}

	return sb.String()
}
