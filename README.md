# semver

`semver` is a command-line tool designed to manage semantic versioning operations on git tags.

It reads the current git tag and creates a new version by incrementing the specified version type (major, minor, or patch). Optionally, it can create a new git tag for this version and push it to the remote origin.

## Installation

You can install it directly if you have Go installed on your system:

```bash
go install github.com/gozeloglu/semver@latest
```

Alternatively, you can clone the project and build it manually:

```bash
git clone https://github.com/gozeloglu/semver.git
cd semver
go build -o semver main.go
# Move the resulting binary to a directory in your PATH
mv semver /usr/local/bin/
```

## Usage

The tool has three main subcommands: `major`, `minor`, and `patch`.

Assuming your current version is tagged as `v1.2.3`:

### Incrementing Major Version

When you make incompatible API changes (results in v2.0.0):

```bash
semver major
```

### Incrementing Minor Version

When you add functionality in a backward-compatible manner (results in v1.3.0):

```bash
semver minor
```

### Incrementing Patch Version

When you make backward-compatible bug fixes (results in v1.2.4):

```bash
semver patch
```

### Flags

All subcommands support the following flags:

- `-t`, `--tag`: Not only prints the new version number to the terminal, but also creates a new git tag with this name in your current repository.
- `-p`, `--push`: Pushes the newly created git tag to the remote server (`origin`) using `git push origin <tag>`. Note that the tag must be created for this to work (typically used in conjunction with `-t`).

**Examples:**

To just calculate and print the next `patch` version:
```bash
semver patch
```

To create a git tag for the next `minor` version:
```bash
semver minor -t
# or
semver minor --tag
```

To create a git tag for the next `major` version and push it to the remote origin:
```bash
semver major -t -p
# or
semver major --tag --push
```
