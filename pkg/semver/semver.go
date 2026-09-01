package semver

import (
	"fmt"
	"strconv"
	"strings"
)

type semVer struct {
	major int
	minor int
	patch int
}

func BumpMajor(v string) string {
	sv := parseVersion(v)
	sv.major++
	sv.minor = 0
	sv.patch = 0
	return newVersion(sv)
}
func BumpMinor(v string) string {
	sv := parseVersion(v)
	sv.minor += 1
	sv.patch = 0
	return newVersion(sv)
}

func BumpPatch(v string) string {
	sv := parseVersion(v)
	sv.patch += 1
	return newVersion(sv)
}

func DefaultTag() string {
	return newVersion(semVer{
		major: 0,
		minor: 0,
		patch: 0,
	})
}

func parseVersion(v string) semVer {
	vv := strings.Split(v[1:], ".")
	return semVer{
		major: convert(vv[0]),
		minor: convert(vv[1]),
		patch: convert(vv[2]),
	}
}

func convert(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

func newVersion(v semVer) string {
	return fmt.Sprintf("v%d.%d.%d", v.major, v.minor, v.patch)
}
