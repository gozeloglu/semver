package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "semver",
	Short: "A CLI tool to manage semantic versioning",
	Long:  `A CLI tool to manage semantic versioning with tags and pushes.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Enable built-in completion command (it's enabled by default, but this makes it explicit)
	rootCmd.CompletionOptions.DisableDefaultCmd = false
}
