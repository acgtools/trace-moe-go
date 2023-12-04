package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var version = "dev"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "moe-go",
	Short:   "A TUI app for finding anime scene by image",
	Long:    `A TUI app for finding anime scene by image, using trace.moe api `,
	Version: version,
}

func init() {
	rootCmd.PersistentFlags().String("log-level", "debug", "log level, options: debug, info, warn, error")

	_ = viper.BindPFlag("log.level", rootCmd.PersistentFlags().Lookup("log-level"))
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
