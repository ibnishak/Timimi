 #!/usr/bin/env bash

pushd ~/.timimi
if [ -f mytouch.sh ]; then
   rm -f mytouch.sh
fi
echo "#!/usr/bin/env bash" > mytouch.sh
echo "touch ~/ExecTestSuccess" >> mytouch.sh
echo "tr -dc A-Za-z0-9 </dev/urandom | head -c 1024 > ~/ExecTestSuccess" >> mytouch.sh
echo "cat ~/ExecTestSuccess" >> mytouch.sh
chmod +x mytouch.sh
popd
pushd /home/richie/Repos/Timimi/Linux/firefox
./helperScripts/json2msg.js < ./helperScripts/exec-test-linux.json | go run timimi.go | ./helperScripts/msg2json.js  
popd