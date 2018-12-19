#!/usr/bin/env bash

pushd ../os-GNULinux
if [ -f timimi ]; then
   rm -f ./timimi
fi
echo "Building timimi"
go build ./timimi.go
echo "Installing timimi"
go run ./install-linux-firefox.go
popd
