package tui

import (
	"image"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/acgtools/trace-moe-go/internal/search"
	"github.com/acgtools/trace-moe-go/internal/util"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/disintegration/imaging"
	"github.com/lucasb-eyer/go-colorful"
)

func traceMoe(dataSrc string, typ searchType) tea.Cmd {
	return func() tea.Msg {
		var (
			res *search.TraceMoeResponse
			err error
		)

		if typ == File {
			res, err = search.File(dataSrc)
		}

		if err != nil {
			return errMsg{err}
		}

		return successMsg{resp: res}
	}
}

func imgToString(width int, dataSrc string, typ searchType) (string, error) {
	var (
		imageContent io.ReadCloser
		err          error
	)

	if typ == File {
		imageContent, err = os.Open(filepath.Clean(dataSrc))
	} else {
		imageContent, err = fetchCoverImage(dataSrc)
	}

	if err != nil {
		return "", err
	}
	defer util.Close(imageContent, &err)

	img, _, err := image.Decode(imageContent)
	if err != nil {
		return "", err
	}

	return toString(width, img), err
}

func toString(width int, img image.Image) string {
	img = imaging.Resize(img, width, 0, imaging.Lanczos)
	b := img.Bounds()
	imageWidth := b.Max.X
	h := b.Max.Y
	str := strings.Builder{}

	for heightCounter := 0; heightCounter < h; heightCounter += 2 {
		for x := imageWidth; x < width; x += 2 {
			str.WriteString(" ")
		}

		for x := 0; x < imageWidth; x++ {
			c1, _ := colorful.MakeColor(img.At(x, heightCounter))
			color1 := lipgloss.Color(c1.Hex())
			c2, _ := colorful.MakeColor(img.At(x, heightCounter+1))
			color2 := lipgloss.Color(c2.Hex())
			str.WriteString(lipgloss.NewStyle().Foreground(color1).
				Background(color2).Render("▀"))
		}

		str.WriteString("\n")
	}

	return str.String()
}

func fetchCoverImage(imgURL string) (io.ReadCloser, error) {
	resp, err := http.Get(imgURL) //nolint:gosec,noctx
	if err != nil {
		return nil, err
	}

	return resp.Body, err
}

func errCmd(err error) tea.Cmd {
	return func() tea.Msg {
		return errMsg{err}
	}
}
