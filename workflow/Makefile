.DEFAULT_GOAL := help

build:
	go build -o ./bin/alfred-gists ./src/main.go
	bumpversion revision --allow-dirty

run:
	go run ./src/main.go

get:
	./go-get-pkgs.sh

help:
	@echo "help"
	@echo "    Show this message"
	@echo ""
	@echo "build"
	@echo "    go build bin"
	@echo ""
	@echo "run"
	@echo "    go run"
	@echo ""
	@echo "get"
	@echo "    go get packages"
	@echo ""
