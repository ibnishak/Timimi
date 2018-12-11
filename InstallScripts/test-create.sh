#!/bin/sh
DIR="/home/richie/Repos/Webextensions/Timimi"
TARGET="/home/richie/Repos/Webextensions/Timimi/za-test/"

mkdir $TARGET && cd $TARGET
mkdir -p windows/chrome windows/chromium windows/firefox linux-mac/chrome linux-mac/chromium linux-mac/firefox
parallel cp -r $DIR/chrome-addon/ {} ::: windows/chrome windows/chromium linux-mac/chrome linux-mac/chromium
parallel cp -r $DIR/firefox-addon/ {} ::: windows/firefox linux-mac/firefox
parallel cp $DIR/timimi-1.4-an+fx.xpi {} ::: windows/firefox linux-mac/firefox
parallel cp $DIR/timimi.exe {} ::: windows/chrome windows/chromium windows/firefox
parallel cp $DIR/timimi.py {} ::: linux-mac/chrome linux-mac/chromium linux-mac/firefox
cp $DIR/windows/chrome/install.bat windows/chrome
cp $DIR/windows/chrome/uninstall.bat windows/chrome
cp $DIR/windows/chrome/timimi.json windows/chrome
cp $DIR/windows/chromium/install.bat windows/chromium
cp $DIR/windows/chromium/uninstall.bat windows/chromium
cp $DIR/windows/chromium/timimi.json windows/chromium
cp $DIR/windows/firefox/install.bat windows/firefox
cp $DIR/windows/firefox/uninstall.bat windows/firefox
cp $DIR/windows/firefox/timimi.json windows/firefox

cp $DIR/linux-mac/chrome/install-linux-mac-chrome.sh linux-mac/chrome
cp $DIR/linux-mac/chrome/uninstall-linux-mac-chrome.sh linux-mac/chrome
cp $DIR/linux-mac/chrome/timimi.json linux-mac/chrome
cp $DIR/linux-mac/chromium/install-linux-mac-chromium.sh linux-mac/chromium
cp $DIR/linux-mac/chromium/uninstall-linux-mac-chromium.sh linux-mac/chromium
cp $DIR/linux-mac/chromium/timimi.json linux-mac/chromium
cp $DIR/linux-mac/firefox/install-linux-mac-firefox.sh linux-mac/firefox
cp $DIR/linux-mac/firefox/uninstall-linux-mac-firefox.sh linux-mac/firefox
cp $DIR/linux-mac/firefox/timimi.json linux-mac/firefox