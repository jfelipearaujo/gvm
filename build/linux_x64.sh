#!/bin/bash

GOOS=linux GOARCH=amd64 go build -tags linux -o ./output/gvm-"$1".linux-x64 ../main.go