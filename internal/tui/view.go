package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const searchSuffix = "Searching..."

var (
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	var sb strings.Builder

	if !m.done {
		sb.WriteString(fmt.Sprintf("%s %s\n", m.spinner.View(), searchSuffix))
	}

	if m.searchResult.result != nil {
		sb.WriteString(fmt.Sprintf("%+v", m.searchResult.result))
	}

	help := fmt.Sprint("\n ctrl+c: exit\n")
	sb.WriteString(helpStyle.Render(help))

	return sb.String()
}
