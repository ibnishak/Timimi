if [ -f timimi.exe ]
then
    rm -f timimi.exe
fi
echo "Building go exec"
env GOOS=windows GOARCH=amd64 go build timimi.go 

echo "Changing unix text files to windows format"

awk 'sub("$", "\r")' unix-readme.txt > readme.txt
awk 'sub("$", "\r")' unix-license-agpl.txt > license.txt
if [ -d "windows-temp-package" ] 
then
rm -rf windows-temp-package
fi
mkdir windows-temp-package
echo "Moving files"
mv timimi.exe windows-temp-package/
mv license.txt windows-temp-package/
mv readme.txt windows-temp-package/
cp static/logo.ico windows-temp-package/
cp static/scripts.ico windows-temp-package/
cp install-windows-firefox.nsi windows-temp-package/
cp timimi.json windows-temp-package/

echo "Windows build finished"