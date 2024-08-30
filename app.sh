#!/bin/bash

set -e

function dev {
    air
}

function deps {
  go mod tidy
  go mod vendor
}

case $1 in
  dev)
    dev
    ;;

  deps)
    deps
    ;;

  build)
    CGO_ENABLED=0 go build -o bin/server main.go
    ;;

  build:dev)
    CGO_ENABLED=0 go build -gcflags="all=-N -l"  -o bin/server main.go
    ;;

  *)
    echo -n "unknown command"
    ;;
esac