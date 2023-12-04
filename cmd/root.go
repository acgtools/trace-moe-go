package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "dev"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "moe-go",
	Short:   "A TUI app for finding anime scene by image",
	Long:    `A TUI app for finding anime scene by image, using trace.moe api `,
	Version: version,
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
