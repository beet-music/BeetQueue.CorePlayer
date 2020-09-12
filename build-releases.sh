#!/bin/bash

mkdir -p releases

GOOS=windows GOARCH=386 go build -o releases/core-player-windows-x86.exe
GOOS=windows GOARCH=amd64 go build -o releases/core-player-windows-x64.exe
GOOS=windows GOARCH=arm go build -o releases/core-player-windows-arm.exe
GOOS=darwin GOARCH=amd64 go build -o releases/core-player-macos-intel
#GOOS=darwin GOARCH=arm go build -o releases/core-player-macos-silicon
GOOS=linux GOARCH=386 go build -o releases/core-player-linux-i386
GOOS=linux GOARCH=amd64 go build -o releases/core-player-linux-amd64
