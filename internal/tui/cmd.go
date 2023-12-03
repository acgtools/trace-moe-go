package tui

import (
	"github.com/acgtools/trace-moe-go/internal/search"
	tea "github.com/charmbracelet/bubbletea"
)

func traceMoe(dataSrc string, typ searchType) tea.Cmd {
	switch typ {
	case File:
		return searchFile(dataSrc)
	case Link:
		// TODO:
	}

	return nil
}

func searchFile(dataSrc string) tea.Cmd {
	res, err := search.File(dataSrc)
	if err != nil {
		return func() tea.Msg {
			return errMsg{err: err}
		}
	}

	return func() tea.Msg {
		return resultMsg{result: res}
	}
}
