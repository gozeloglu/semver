package cmd

import (
	"fmt"

	"github.com/gozeloglu/semver/pkg/git"
	"github.com/gozeloglu/semver/pkg/semver"
	"github.com/spf13/cobra"
	gosemver "golang.org/x/mod/semver"
)

var majorCmd = &cobra.Command{
	Use:   "major",
	Short: "Increment the major version",
	Run: func(cmd *cobra.Command, args []string) {
		tagFlag, _ := cmd.Flags().GetBool("tag")
		pushFlag, _ := cmd.Flags().GetBool("push")

		oldVersion := git.LatestTag()
		if ok := gosemver.IsValid(oldVersion); !ok {
			fmt.Println("latest tag is not valid")
			return
		}

		newVersion := semver.BumpMajor(oldVersion)
		processFlags(oldVersion, newVersion, tagFlag, pushFlag)
	},
}

func init() {
	rootCmd.AddCommand(majorCmd)
	majorCmd.Flags().BoolP("tag", "t", false, "Create a git tag")
	majorCmd.Flags().BoolP("push", "p", false, "Push the git tag to remote origin")
}
