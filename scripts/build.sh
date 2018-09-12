#!/usr/bin/env bash

curdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
projdir="$(cd "$curdir/.." && pwd)"

# set the gopath
export GOPATH="$HOME/Projects/golang"
mkdir -p "$GOPATH"

# -u updates, -v verbose
pkgs=(
    golang.org/x/oauth2
    github.com/google/go-github/github
    github.com/deanishe/awgo
)

for p in "${pkgs[@]}"; do
    if [[ ! -d "${GOPATH}/src/$p" ]]; then
        # go get the package and output that command
        (set -x; go get "$p")
    fi
done

go build -o "$projdir/workflow/alfred-gists" "$projdir/src/main.go"
bumpversion revision --allow-dirty
