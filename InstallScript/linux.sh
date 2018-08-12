git clone --depth=1 https://github.com/ibnishak/Timimi.git
cp -r Timimi/native-messaging-hosts $HOME/.mozilla
sed -ie "s/richie/$USER/" $HOME/.mozilla/native-messaging-hosts/timimi.json
cp Timimi/addons/web-ext-artifacts/timimi-1.3-an+fx.xpi $HOME