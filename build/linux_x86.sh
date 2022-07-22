#!/bin/bash

GOOS=linux GOARCH=386 go build -tags linux -o ./output/gvm-"$1".linux-x86 ../main.go