#!/usr/bin/env bash 

if [ ! -d "../pkg/linux-firefox" ]; then
  mkdir -p ../pkg/mac-firefox
  fi 

echo "Building timimi"
env GOOS=darwin GOARCH=amd64 go build timimi.go
echo "Moving timimi"
mv timimi ../pkg/mac-firefox
echo "Building install"
pushd install-mac-firefox
go build install-mac-firefox.go
echo "Moving install"
mv install-mac-firefox ../../pkg/mac-firefox
popd

pushd ../pkg
echo "Zipping"
zip -r mac-firefox.zip mac-firefox
rm -rf mac-firefox
popd
echo "Finished"
