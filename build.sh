#!/bin/sh

GOOS=darwin GOARCH=amd64 go build -o acolytes && zip -r ./releases/acolytes-darwin ./images ./acolytes && rm acolytes \
  && GOOS=linux GOARCH=amd64 go build -o acolytes && zip -r ./releases/acolytes-linux ./images ./acolytes && rm acolytes \
  && GOOS=windows GOARCH=amd64 go build -o acolytes.exe && zip -r ./releases/acolytes-win ./images ./acolytes.exe && rm acolytes.exe
