#!/bin/bash

rm -rf bin
mkdir -p bin

echo "Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o bin/docker-monitor-windows-amd64.exe

echo "Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o bin/docker-monitor-linux-amd64

echo "Building for Mac (Intel)..."
GOOS=darwin GOARCH=amd64 go build -o bin/docker-monitor-darwin-amd64

echo "Building for Mac (Apple Silicon - M1/M2)..."
GOOS=darwin GOARCH=arm64 go build -o bin/docker-monitor-darwin-arm64

echo "Done! Check the 'bin' folder."