package cmd

import (
	"fmt"

	"github.com/gozeloglu/semver/pkg/git"
	"github.com/gozeloglu/semver/pkg/semver"
	"github.com/spf13/cobra"
	gosemver "golang.org/x/mod/semver"
)

var minorCmd = &cobra.Command{
	Use:   "minor",
	Short: "Increment the minor version",
	Run: func(cmd *cobra.Command, args []string) {
		tagFlag, _ := cmd.Flags().GetBool("tag")
		pushFlag, _ := cmd.Flags().GetBool("push")

		oldVersion := git.LatestTag()
		if ok := gosemver.IsValid(oldVersion); !ok {
			fmt.Println("Latest tag is not valid")
			return
		}
		newVersion := semver.BumpMinor(oldVersion)
		processFlags(oldVersion, newVersion, tagFlag, pushFlag)
	},
}

func init() {
	rootCmd.AddCommand(minorCmd)
	minorCmd.Flags().BoolP("tag", "t", false, "Create a git tag")
	minorCmd.Flags().BoolP("push", "p", false, "Push the git tag to remote origin")
}
