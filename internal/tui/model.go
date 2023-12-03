package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

type searchType uint8

const (
	File searchType = iota
	Link
)

type Model struct {
	spinner      spinner.Model
	dataSrc      string
	searchType   searchType
	searchResult resultMsg
	done         bool
	err          error
}

var _ tea.Model = Model{}

func NewModel(dataSrc string, typ searchType) Model {
	return Model{
		spinner:    spinner.New(),
		dataSrc:    dataSrc,
		searchType: typ,
	}
}
