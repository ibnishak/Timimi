#!/usr/bin/env bash

pushd /home/richie/Repos/Timimi/Linux/firefox
if [ -f timimi ]; then
   rm -f ./timimi
fi
echo "Building timimi"
go build ./timimi.go
echo "Installing timimi"
go run ./helperScripts/install.go
popd
