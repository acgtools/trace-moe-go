package cmd

import (
	"errors"
	"log/slog"
	"os"

	"github.com/acgtools/trace-moe-go/internal/search"

	"github.com/spf13/cobra"

	log "github.com/acgtools/trace-moe-go/pkg/log"
)

const minArgNum = 1

var fileCmd = &cobra.Command{
	Use:   "file",
	Short: "",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < minArgNum {
			return errors.New("not enough arguments")
		}

		config, err := NewConfig()
		if err != nil {
			return err
		}

		logLevel, err := log.ParseLevel(config.Log.Level)
		if err != nil {
			return err //nolint:wrapcheck
		}
		logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
			ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
				if a.Key == slog.TimeKey {
					return slog.Attr{}
				}
				return a
			},
		}))
		slog.SetDefault(logger)

		return search.File(args[0])
	},
}

func init() {

	rootCmd.AddCommand(fileCmd)
}
