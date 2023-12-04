package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/acgtools/trace-moe-go/internal/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

const minArgNum = 1

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "search image by file",
	Long:  `file [flags] [image file path]`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < minArgNum {
			return errors.New("not enough arguments")
		}

		filePath := args[0]
		info, err := os.Lstat(filePath)
		if err != nil {
			return fmt.Errorf("invalid path: %w", err)
		}

		if info.IsDir() {
			return errors.New("image cannot be a directory")
		}

		m := tui.New(filePath, tui.File)
		var opts []tea.ProgramOption

		// Always append alt screen program option.
		opts = append(opts, tea.WithAltScreen())

		if _, err := tea.NewProgram(m, opts...).Run(); err != nil {
			return err
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(fileCmd)
}
