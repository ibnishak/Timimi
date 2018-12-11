#!/usr/bin/env bash

if [ -f timimi ]; then
   rm -f ./timimi
fi
go build ./timimi.go
go run install.go
