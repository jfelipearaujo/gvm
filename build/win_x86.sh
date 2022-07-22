#!/bin/bash

GOOS=windows GOARCH=386 go build -tags windows -o ./output/gvm-"$1".windows-x86.exe ../main.go