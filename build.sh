#!/bin/sh

cp -r ./images ./releases

GOOS=darwin GOARCH=amd64 go build && mv ./acolytes ./releases/acolytes-darwin \
  && GOOS=linux GOARCH=amd64 go build && mv ./acolytes ./releases/acolytes-linux \
  && GOOS=windows GOARCH=amd64 go build && mv ./acolytes.exe ./releases/acolytes.exe
