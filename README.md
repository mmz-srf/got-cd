# got-cd

Go implementation of https://github.com/gitcd-io/gitcd for learning purposes

## Usage
### Create Config file
```
cp example.config.json <UserHomeDir>/.got-cd/config.json
```

### Run git-cd

```
git-cd is a command-line tool that simplifies the usage of git commands by providing a more user-friendly interface.

Usage:
  git-cd [flags]
  git-cd [command]

Available Commands:
  clean       Clean up local branches
  completion  Generate the autocompletion script for the specified shell
  finish      Merge the feature branch into main
  help        Help about any command
  login       Login to GitHub
  open        Open the current feature branch in the browser
  release     Create a new release
  review      Create a pull request from feature branch to main
  start       Start a new feature branch
  status      Get the status of an open pull request
  test        Merge feature branch into test
  version     Print the version of got-cd

Flags:
  -h, --help        help for git-cd
  -s, --short-tag   Use short tag format (e.g., 1.0.0 instead of v1.0.0)
  -v, --verbose     Enable verbose mode

Use "git-cd [command] --help" for more information about a command.
```


## Build your own
1. Build binary

```
 make build
```
2. Run binary
```
./bin/git-cd -h
```
