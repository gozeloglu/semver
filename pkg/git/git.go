package git

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/gozeloglu/semver/pkg/semver"
)

func LatestTag() string {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	out, err := cmd.Output()
	if err != nil {
		return semver.DefaultTag()
	}
	return strings.TrimSpace(string(out))
}

func Tag(v string) error {
	cmd := exec.Command("git", "tag", "-a", v, "-m", fmt.Sprintf("Release %s", v))
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to create tag: %w", err)
	}
	return nil
}

func Push(v string) error {
	cmd := exec.Command("git", "push", "origin", v)
	_, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("failed to push new tag: %w", err)
	}
	return nil
}
