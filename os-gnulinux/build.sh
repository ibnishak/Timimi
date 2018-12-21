#!/usr/bin/env bash 

if [ ! -d "../pkg/linux-firefox" ]; then
  mkdir -p ../pkg/linux-firefox
  fi 

echo "Building timimi"
go build timimi.go
echo "Moving timimi"
mv timimi ../pkg/linux-firefox
echo "Building install"

cd install-linux-firefox
go build install-linux-firefox.go
echo "Moving install"
mv install-linux-firefox ../../pkg/linux-firefox
echo "Finished"