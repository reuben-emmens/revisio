set dotenv-load

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

set-control: 
  mkdir -pv $ROOT/DEBIAN
  echo "package: $PACKAGE" > $CONTROL
  echo "section: $SECTION" >> $CONTROL
  echo "version: $VERSION" >> $CONTROL
  echo "priority: $PRIORITY" >> $CONTROL
  echo "architecture: $ARCHITECTURE" >> $CONTROL
  echo "maintainer: $MAINTAINER" >> $CONTROL
  echo "description: $DESCRIPTION" >> $CONTROL

deb-pkg: build set-control
  mkdir -pv $BIN 
  mv revisio $BIN
  dpkg-deb --build $ROOT $PACKAGES
  sha256sum $PACKAGES/$DEBPKG > $PACKAGES/$DEBSHA

var:
  echo $DEBPKG