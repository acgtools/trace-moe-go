package tui

import (
	"fmt"
	"strings"

	"github.com/acgtools/trace-moe-go/internal/search"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"

	tea "github.com/charmbracelet/bubbletea"
)

type searchType uint8

const (
	File searchType = iota
	Link
)

type Model struct {
	spinner    spinner.Model
	dataSrc    string
	searchType searchType
	done       bool
	err        error
	resList    list.Model
	size       winSize
}

type winSize struct {
	width  int
	height int
}

var _ tea.Model = Model{}

func New(dataSrc string, typ searchType) Model {
	return Model{
		spinner:    spinner.New(),
		dataSrc:    dataSrc,
		searchType: typ,
	}
}

type resultItem struct {
	index      int
	episode    int
	name       string
	from       float64
	to         float64
	similarity float64
}

func (i resultItem) Title() string { return i.name }

func (i resultItem) Description() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Episode: %d\n", i.episode))
	sb.WriteString(fmt.Sprintf("From %f to %f\n", i.from, i.to))
	sb.WriteString(fmt.Sprintf("Similarity: %f", i.similarity))

	return sb.String()
}

func (i resultItem) FilterValue() string { return i.name }

func newList(results []*search.Result, size winSize) list.Model {
	items := make([]list.Item, 0, len(results))
	for i, res := range results {
		const minSim = 0.9
		if res.Similarity < minSim {
			continue
		}

		item := resultItem{
			index:      i,
			episode:    res.Episode,
			name:       res.Anilist.Title.Native,
			from:       res.From,
			to:         res.To,
			similarity: res.Similarity,
		}
		items = append(items, item)
	}

	delegate := list.NewDefaultDelegate()
	delegate.SetHeight(4)

	l := list.New(items, delegate, size.width, size.height)
	l.Title = "Result list"

	return l
}
