#!/usr/bin/env bash
pushd /home/richie/Repos/Timimi/Linux/firefox
./helperScripts/json2msg.js < ./helperScripts/test-working.json | go run timimi.go | ./helperScripts/msg2json.js   
popd