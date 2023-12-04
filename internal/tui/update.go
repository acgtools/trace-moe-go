package tui

import (
	"github.com/acgtools/trace-moe-go/internal/search"
	tea "github.com/charmbracelet/bubbletea"
)

type successMsg struct {
	resp *search.TraceMoeResponse
}

type errMsg struct {
	err error
}

var _ error = errMsg{}

func (e errMsg) Error() string {
	return e.err.Error()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) { //nolint:cyclop
	var (
		cmd  tea.Cmd
		cmds []tea.Cmd
	)

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.size = size{
			width:  msg.Width,
			height: msg.Height,
		}
		m.imgWidth = msg.Width / 3
		img, err := imgToString(m.imgWidth, m.dataSrc, m.searchType)
		if err != nil {
			return m, errCmd(err)
		}
		m.image = img
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case " ", "enter":
			if m.ready {
				res := m.resp.Result[m.resList.Index()]
				m.setInfo(res)
				err := m.setPreview(res)
				if err != nil {
					return m, errCmd(err)
				}
			}
		}
	case successMsg:
		m.ready = true
		m.resp = msg.resp
		m.resList = newList(msg.resp.Result, m.size)
		if len(msg.resp.Result) > 0 {
			res := m.resp.Result[0]
			m.setInfo(res)
			err := m.setPreview(res)
			if err != nil {
				return m, errCmd(err)
			}
		}
	case errMsg:
		m.err = msg.err
	}

	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)

	if m.ready {
		m.resList, cmd = m.resList.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) setInfo(res *search.Result) {
	m.info = newAnimeInfo(res.Anilist)
}

func (m *Model) setPreview(res *search.Result) error {
	img, err := imgToString(m.imgWidth, res.Image, Link)
	if err != nil {
		return err
	}
	m.matchPreview = img

	return nil
}
