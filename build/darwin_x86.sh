#!/bin/bash

GOOS=darwin GOARCH=386 go build -tags darwin -o ./output/gvm-"$1".darwin-x86 ../main.go