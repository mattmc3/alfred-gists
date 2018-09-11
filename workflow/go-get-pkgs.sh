#!/usr/bin/env bash

curdir="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
# gp="$(cd $curdir/../../../.. && pwd)"
# if [[ $gp != "$HOME"* ]]; then
#     echo "Expecting \$GOPATH to be off your \$HOME."
#     exit 1
# fi
# export GOPATH="$gp"
export GOPATH="$HOME/Projects/golang"
mkdir -p "$GOPATH"

# -u updates, -v verbose
go get golang.org/x/oauth2
go get github.com/google/go-github/github
go get github.com/deanishe/awgo
