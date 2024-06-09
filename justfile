# list all recipes
list:
  @just --list --unsorted

fmt:
  @go fmt ./...

vet: fmt
  @go vet ./...

lint: vet
  @golangci-lint run ./...

test: lint
  @go test ./...

alias b := build
build: lint
  @go build -o revisio cmd/revisio/main.go

path: build
  @mkdir -p ~/bin
  @cp revisio ~/bin
  export PATH=$PATH:$HOME/bin