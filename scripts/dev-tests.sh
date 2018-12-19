#!/usr/bin/env bash

ln -s ../os-GNULinux/timimi.go timimi.go

echo "Removing Previous logs if any"
rm -f ./demoLog/*

echo "Copying demo scripts to ~/.timimi"
if [ ! -d ~/.timimi ]; then
mkdir ~/.timimi
fi
cp ./demoScripts/* ~/.timimi

echo "Running tests"
for i in $( ls ./demoData); do
		echo $i >  ./demoLog/$i.txt
		echo "------------------------------" >>  ./demoLog/$i.txt
            go run demoTest/json2msg.go < demoData/$i | go run timimi.go | go run demoTest/msg2json.go >> ./demoLog/$i.txt
        done
        echo "Concatenating logs"
        cat ./demoLog/* >> ./demoLog/all-cat-log.txt
        echo "Cleaning up"
        rm -f ./demoPath/alternateBpath/*
        pushd ./demoPath
        find . ! -name 'a.txt' -type f -exec rm -f {} +
        popd
rm -f ./timimi.go
