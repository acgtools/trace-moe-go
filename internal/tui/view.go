package tui

import (
	"fmt"
	"strings"
)

func (m Model) View() string {
	if m.err != nil {
		return m.err.Error()
	}

	var sb strings.Builder

	if !m.done {
		sb.WriteString(fmt.Sprintf("%s %s\n", "Searching...", m.spinner.View()))
		sb.WriteString("help - q: exit")
	} else {
		sb.WriteString(m.resList.View())
	}

	return sb.String()
}
