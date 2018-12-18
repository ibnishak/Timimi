#!/usr/bin/env bash



echo "Removing Previous logs if any"
rm -f ./demoLog/*

echo "Copying demo scripts to ~/.timimi"
if [ ! -d ~/.timimi ]; then
mkdir ~/.timimi
fi
cp ./demoScripts/* ~/.timimi

echo "Running tests"
for i in $( ls ./demoData); do
            ./json2msg.js < demoData/$i | go run timimi.go | ./msg2json.js > ./demoLog/$i
        done
        echo "Concatenating logs"
        cat ./demoLog/* >> ./demoLog/all-cat-log.json
        echo "Cleaning up"
        rm -f ./demoPath/alternateBpath/*
        pushd ./demoPath
        find . ! -name 'a.txt' -type f -exec rm -f {} +
        popd