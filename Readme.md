
[![code style: prettier](https://img.shields.io/badge/code_style-prettier-ff69b4.svg?style=flat-square)](https://github.com/prettier/prettier)


Tiddlyfox addon for firefox browser stopped working for post 57 versions of firefox when mozilla switched over to sandbox model, restricting access to the file system. However webextension addons provide APIs which interface with filesystem. This is an addon using native messaging, essentially handing over the contents to a webextension host which does the actual saving. For the end user it means a couple of extra steps to install the addon but once installed, you can save the standalone TW from anywhere in your harddrive without any more interactions, like the original Tiddlyfox addon did.

The set up is currently complete for Linux and Mac users. Windows workflow is slightly different and since I do not have a windows machine to test, I cannot provide assistance. However core of the application remains the same, and the work to port it for chrome browser and windows platform would be minimal. The principle followed can be used to develop addons for Edge browser too, for those so inclined.

### Requirements:

* Firefox >57
* Python - usually installed by default in linux systems



### Installation:

#### Linux

Debian derivatives(Debian, Ubuntu, Linux Mint etc), Arch Derivatives (Arch, Antergos, Manjaro)

* Step 1: Run

```
sh -c "$(wget https://raw.githubusercontent.com/ibnishak/Timimi/master/InstallScript/linux.sh -O -)"
```

* Step 2 : In your home folder you will find **timimi.xpi**. Drag and drop it to Firefox addon manager page (about:addons).



If you do not have git and wget installed in your system, or if the install script is throwing errors, follow the steps given below.

* Download the repo from **https://github.com/ibnishak/Timimi**
* Unzip the downloaded file and open the folder.
* Open **native-messaging-hosts/timimi.json** and change the name `richie` in the path key to your username. 
  * Find your username by typing `echo $USER` in your terminal
* Copy **native-messaging-hosts** folder to ~/.mozilla. 
  * If you already have a **native-messaging-hosts** folder in ~/.mozilla, you can safely merge them.
* In **Timimi/addons/web-ext-artifacts** you will find **timimi.xpi**. Drag and drop it to Firefox addon manager page (about:addons).


### Mac - untested

* Download the repo from **https://github.com/ibnishak/Timimi**
* Unzip the downloaded file and open the folder.
* Open **native-messaging-hosts/timimi.json** and change the path to **~/Library/Application Support/Mozilla/NativeMessagingHosts/timimi.py**
* From **native-messaging-hosts** folder, copy timimi.json and timimi.py to **~/Library/Application Support/Mozilla/NativeMessagingHosts/**
* In **Timimi/addons/web-ext-artifacts** you will find **timimi.xpi**. Drag and drop it to Firefox addon manager page (about:addons).



#### Windows:
Need more time.



### Status

1. Presently it is a barebones addon, as in it just does one thing - saving. Other features of old tiddlyfox plugin, like disabling, showing save status etc can be easily added. I leave it to those who are interested and familiar with Tfox code.

2. The webextension host is currently written in python under the assumption that linux systems usually have it installed by default. It is possible to write the same in nodejs or other language for different platforms to minimize number of installations.


### Pros

1. Webextension host can be expanded for different purposes depending on personal preferences. For eg:
 * Saving time stamped backups instead of overwrting.
 * Monitor changes to a tiddler and enter it to a local relational database like sqlite.

2. It can be easily ported over to chrome
3. No need to run a server, does not require continued user interactions

### Cons
1. Cannot be used in android

