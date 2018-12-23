#!/data/data/com.termux/files/usr/bin/env bash

#ln -s ../os-gnulinux/timimi-test-termux.go timimi-test-termux.go

echo "Removing Previous logs if any"
rm -f ./demoLog/*

echo "Copying demo scripts to ~/.timimi"
if [ ! -d ~/.timimi ]; then
mkdir ~/.timimi
fi
# cp ./demoScripts/* ~/.timimi

echo "Running tests"
for i in $( ls ./demoDataTermux); do
		echo $i >  ./demoLog/$i.txt
		echo "------------------------------" >>  ./demoLog/$i.txt
            go run demoTest/json2msg.go < demoDataTermux/$i | go run timimi-test-termux.go | go run demoTest/msg2json.go >> ./demoLog/$i.txt
        done
        echo "Concatenating logs"
        cat ./demoLog/* >> ./demoLog/all-cat-log.txt
        echo "Cleaning up"
        rm -f ./demoPath/alternateBpath/*
        pushd ./demoPath
        find . ! -name 'a.txt' -type f -exec rm -f {} +
        popd
#rm -f ./timimi-test-termux.go
sed '/^Errors/!d' ./demoLog/all-cat-log.txt > ./demoLog/Errors.txt
cat ./demoLog/Errors.txt
