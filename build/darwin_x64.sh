#!/bin/bash

GOOS=darwin GOARCH=amd64 go build -tags darwin -o ./output/gvm-"$1".darwin-x64 ../main.go