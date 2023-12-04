package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	layoutStyle     = lipgloss.NewStyle()
	imageStyle      = lipgloss.NewStyle()
	previewImgStyle = lipgloss.NewStyle()
	listStyle       = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())
	videoStyle      = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())
	infoStyle       = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())
	errStyle        = lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())
)

func (m Model) View() string {
	if m.err != nil {
		return errStyle.Render(fmt.Sprintf("%v\npress q to exit", m.err))
	}

	var sb strings.Builder

	if !m.ready {
		sb.WriteString(fmt.Sprintf("%s %s\n", "Searching...", m.spinner.View()))
		sb.WriteString("help - q: exit")
	} else {
		infoWidth, infoHeight := m.size.width/4, m.resList.Height()

		image := imageStyle.Render("Your search image\n" + m.image)
		previewImg := previewImgStyle.Render("Matched image preview\n" + m.matchPreview)

		list := listStyle.Render(m.resList.View())

		videoStyle = videoStyle.Width(infoWidth).Height(10).Align(lipgloss.Center, lipgloss.Center)
		video := videoStyle.Render("Assuming there is a video here. >_<")

		infoStyle = infoStyle.Width(infoWidth).Height(infoHeight - 12)
		info := infoStyle.Render(formatAnimeInfo(m.info))

		output := layout(image, previewImg, list, video, info)
		layoutStyle = layoutStyle.MaxWidth(m.size.width).MaxHeight(m.size.height)

		sb.WriteString(layoutStyle.Render(output))
	}

	return sb.String()
}

func layout(image, previewImg, list, video, info string) string {
	right := lipgloss.JoinVertical(lipgloss.Top, video, info)
	images := lipgloss.JoinVertical(lipgloss.Top, image, previewImg)
	return lipgloss.JoinHorizontal(lipgloss.Top, list, right, images)
}

func formatAnimeInfo(info animeInfo) string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Anilist ID: %d\n\n", info.anilistID))
	sb.WriteString(fmt.Sprintf("Is adult: %t\n\n", info.isAdult))

	for _, name := range info.names {
		sb.WriteString(name + "\n\n")
	}

	return sb.String()
}
