#!/bin/sh

DIR="/home/richie/Repos/Webextensions/Timimi"
TARGET="/home/richie/Repos/Webextensions/Timimi/zb-packages/"


mkdir $TARGET && cd $TARGET
mkdir -p windows/chrome-chromium windows/firefox linux-mac/chrome-chromium linux-mac/firefox
parallel cp -r $DIR/chrome-addon/ {} ::: windows/chrome-chromium linux-mac/chrome-chromium
parallel cp $DIR/timimi.exe {} ::: windows/chrome-chromium windows/firefox
parallel cp $DIR/timimi-1.4-an+fx.xpi {} ::: windows/firefox linux-mac/firefox
parallel cp $DIR/timimi.py {} ::: linux-mac/chrome-chromium linux-mac/firefox

cp $DIR/windows/chrome/timimi.json windows/chrome-chromium/
cp $DIR/windows/chrome/install.bat windows/chrome-chromium/windows-chrome-install.bat
cp $DIR/windows/chrome/uninstall.bat windows/chrome-chromium/windows-chrome-uninstall.bat
cp $DIR/windows/chromium/install.bat windows/chrome-chromium/windows-chromium-install.bat
cp $DIR/windows/chromium/uninstall.bat windows/chrome-chromium/windows-chromium-uninstall.bat
cp $DIR/windows/firefox/timimi.json windows/firefox/
cp $DIR/windows/firefox/install.bat windows/firefox/
cp $DIR/windows/firefox/uninstall.bat windows/firefox/

cp $DIR/linux-mac/chrome/install-linux-mac-chrome.sh linux-mac/chrome-chromium/
cp $DIR/linux-mac/chrome/uninstall-linux-mac-chrome.sh linux-mac/chrome-chromium/
cp $DIR/linux-mac/chrome/timimi.json linux-mac/chrome-chromium/
cp $DIR/linux-mac/chromium/install-linux-mac-chromium.sh linux-mac/chrome-chromium/
cp $DIR/linux-mac/chromium/uninstall-linux-mac-chromium.sh linux-mac/chrome-chromium/
cp $DIR/linux-mac/firefox/install-linux-mac-firefox.sh linux-mac/firefox/
cp $DIR/linux-mac/firefox/uninstall-linux-mac-firefox.sh linux-mac/firefox/
cp $DIR/linux-mac/firefox/timimi.json linux-mac/firefox/

zip -r win64-chrome-chromium.zip windows/chrome-chromium
zip -r win64-firefox.zip windows/firefox
tar -zcvf linux-mac-chrome-chromium.tar.gz linux-mac/chrome-chromium
tar -zcvf linux-mac-firefox.tar.gz linux-mac/firefox

rm -rf windows/
rm -rf linux-mac/
