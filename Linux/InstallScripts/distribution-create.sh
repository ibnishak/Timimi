#!/bin/sh
DIR="/home/richie/Repos/Webextensions/Timimi"
TARGET="/home/richie/Repos/Webextensions/Timimi/za-distribution/"

mkdir $TARGET && cd $TARGET
mkdir -p windows/chrome windows/chromium windows/firefox linux-mac/chrome linux-mac/chromium linux-mac/firefox
parallel ln -s $DIR/chrome-addon/ {} ::: windows/chrome windows/chromium linux-mac/chrome linux-mac/chromium
parallel ln -s $DIR/firefox-addon/ {} ::: windows/firefox linux-mac/firefox
parallel ln -s $DIR/timimi-1.4-an+fx.xpi {} ::: windows/firefox linux-mac/firefox
parallel ln -s $DIR/timimi.exe {} ::: windows/chrome windows/chromium windows/firefox
parallel ln -s $DIR/timimi.py {} ::: linux-mac/chrome linux-mac/chromium linux-mac/firefox
ln -s $DIR/windows/chrome/install.bat windows/chrome
ln -s $DIR/windows/chrome/uninstall.bat windows/chrome
ln -s $DIR/windows/chrome/timimi.json windows/chrome
ln -s $DIR/windows/chromium/install.bat windows/chromium
ln -s $DIR/windows/chromium/uninstall.bat windows/chromium
ln -s $DIR/windows/chromium/timimi.json windows/chromium
ln -s $DIR/windows/firefox/install.bat windows/firefox
ln -s $DIR/windows/firefox/uninstall.bat windows/firefox
ln -s $DIR/windows/firefox/timimi.json windows/firefox

ln -s $DIR/linux-mac/chrome/install-linux-mac-chrome.sh linux-mac/chrome
ln -s $DIR/linux-mac/chrome/uninstall-linux-mac-chrome.sh linux-mac/chrome
ln -s $DIR/linux-mac/chrome/timimi.json linux-mac/chrome
ln -s $DIR/linux-mac/chromium/install-linux-mac-chromium.sh linux-mac/chromium
ln -s $DIR/linux-mac/chromium/uninstall-linux-mac-chromium.sh linux-mac/chromium
ln -s $DIR/linux-mac/chromium/timimi.json linux-mac/chromium
ln -s $DIR/linux-mac/firefox/install-linux-mac-firefox.sh linux-mac/firefox
ln -s $DIR/linux-mac/firefox/uninstall-linux-mac-firefox.sh linux-mac/firefox
ln -s $DIR/linux-mac/firefox/timimi.json linux-mac/firefox