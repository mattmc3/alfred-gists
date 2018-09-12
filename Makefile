.DEFAULT_GOAL := help

build:
	./scripts/build.sh

run:
	go run ./src/main.go

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
