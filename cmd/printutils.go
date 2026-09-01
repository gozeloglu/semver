package cmd

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/gozeloglu/semver/pkg/git"
)

var (
	cyan  = color.New(color.FgCyan).SprintFunc()
	green = color.New(color.FgGreen).SprintFunc()
	red   = color.New(color.FgRed).SprintFunc()
)

// PrintVersionBump prints the version transition with colors
func PrintVersionBump(oldVersion, newVersion string) {
	fmt.Printf("%s → %s\n", cyan(oldVersion), green(newVersion))
}

func PrintSuccessTagCreation(newVersion string) {
	fmt.Printf("Successfully created tag %s\n", green(newVersion))
}

func PrintSuccessTagPush(newVersion string) {
	fmt.Printf("Successfully pushed tag %s\n", green(newVersion))
}

func PrintCreatingTag(newVersion string) {
	fmt.Printf("Creating tag %s...\n", newVersion)
}

func PrintPushingTag(newVersion string) {
	fmt.Printf("Pushing tag %s...\n", newVersion)
}
func PrintErrorTagCreation(err error) {
	fmt.Println("Error creating tag:", err)
}

func PrintErrorTagPush(err error) {
	fmt.Println("Error pushing tag:", err)
}

func processFlags(oldVersion string, newVersion string, tagFlag bool, pushFlag bool) {
	PrintVersionBump(oldVersion, newVersion)

	if tagFlag {
		PrintCreatingTag(newVersion)
		err := git.Tag(newVersion)
		if err != nil {
			PrintErrorTagCreation(err)
			return
		}
		PrintSuccessTagCreation(newVersion)
	}

	if pushFlag {
		PrintPushingTag(newVersion)
		err := git.Push(newVersion)
		if err != nil {
			PrintErrorTagPush(err)
			return
		}
		PrintSuccessTagPush(newVersion)
	}
}
