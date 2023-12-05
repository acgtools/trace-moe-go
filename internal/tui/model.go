package tui

import (
	"fmt"
	"math"
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
	spinner      spinner.Model
	dataSrc      string
	searchType   searchType
	resp         *search.TraceMoeResponse
	image        string
	matchPreview string
	ready        bool
	err          error
	resList      list.Model
	info         animeInfo
	size         size
	imgWidth     int
}

type size struct {
	width  int
	height int
}

type animeInfo struct {
	anilistID int
	names     []string
	isAdult   bool
}

func newAnimeInfo(ani *search.Anilist) animeInfo {
	var names []string
	names = append(names,
		ani.Title.Native,
		ani.Title.Romaji,
		ani.Title.English)
	names = append(names, ani.Synonyms...)

	return animeInfo{
		anilistID: ani.ID,
		names:     names,
		isAdult:   ani.IsAdult,
	}
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
	episode    search.Episode
	name       string
	from       float64
	to         float64
	similarity float64
}

func (i resultItem) Title() string { return i.name }

func (i resultItem) Description() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("Episode: %s\n", i.episode))
	sb.WriteString(fmt.Sprintf("From %s to %s\n", formatTime(i.from), formatTime(i.to)))
	sb.WriteString(fmt.Sprintf("Similarity: %.3f%s", i.similarity*100, "%"))

	return sb.String()
}

func (i resultItem) FilterValue() string { return i.name }

func newList(results []*search.Result, size size) list.Model {
	items := make([]list.Item, 0, len(results))
	for i, res := range results {
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

	w, h := listStyle.GetFrameSize()
	l := list.New(items, delegate, size.width-w, size.height-h)
	l.Title = "Result list"

	return l
}

func formatTime(secs float64) string {
	seconds := int(math.Round(secs))
	m, s := seconds/60, seconds%60
	return fmt.Sprintf("%d:%d", m, s)
}
