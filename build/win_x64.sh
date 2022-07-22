#!/bin/bash

GOOS=windows GOARCH=amd64 go build -tags windows -o ./output/gvm-"$1".windows-x64.exe ../main.go