# alfred-gists

Alfred workflow for GitHub gists

## Developer Notes

### Todos

The list of project todos is maintained in the `todo.taskpaper` file.

### Go

Go ([golang](https://golang.org)) is used to create the core gist utility. This
allows for distribution of a binary instead of relying on bash, or Mac's
grossly out dated versions of python or ruby. If you want to make changes to
the code, install golang (1.11+) via homebrew, and then utilize the makefile to
do binary builds.

### Bumpversion

[bumpversion](https://github.com/peritus/bumpversion) is used to maintain
semantic versioning.

Example: `bumpversion revision --allow-dirty`

### Issues

If you have issues with this workflow, or recommendations, please open a
GitHub issue [here](https://github.com/mattmc3/alfred-gists/issues).
