<p align="center">
<img src="https://raw.githubusercontent.com/ibnishak/Timimi/master/resources/fish.png" width=400/>
</p>


<h2 align="center">Webextensions for Tiddlywiki</h2>
<p align="center">
<a href="https://github.com/prettier/prettier" class="rich-diff-level-one">
	<img src="https://camo.githubusercontent.com/687a8ae8d15f9409617d2cc5a30292a884f6813a/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f636f64655f7374796c652d70726574746965722d6666363962342e7376673f7374796c653d666c61742d737175617265" alt="code style: prettier" data-canonical-src="https://img.shields.io/badge/code_style-prettier-ff69b4.svg?style=flat-square" style="max-width:100%;">
</a>
<a href="https://GitHub.com/ibnishak/timimi/graphs/commit-activity" class="rich-diff-level-one">
	<img src="https://camo.githubusercontent.com/0e6a3f975d68b438efec82fef1f9491600606df8/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f4d61696e7461696e65642533462d7965732d677265656e2e737667" alt="Maintenance" data-canonical-src="https://img.shields.io/badge/Maintained%3F-yes-green.svg" style="max-width:100%;">
</a>
	<img src="resources/golang.svg" style="max-width:100%;">
	<img src="resources/licence.svg" style="max-width:100%;">
	<a href="https://tiddlywiki.com/" target="_blank"><img src="resources/tw.svg" style="max-width:100%;"></a>
</a>
</p>


<br/><br/><br/>

### Announcement


Owing to less than optimal conditions in personal life, porting Timimi 2.0 to chrome and other browsers are indefinitely postponed. Timimi 1.0 will still work in Windows and Linux in Firefox as well as chrome, while Timimi 2.0 will work in Windows, Linux and Mac but only in Firefox browser. So if you are a Firefox user, you are strongly urged to update to lastest Timimi release.

Sincerely,
Riz


* [Updates](#updates)
* [Intro](#intro)
* [Supported browsers](#supported-browsers)
* [Supported OS](#supported-os)
* [Installation](#installation)
* [To do](#to-do)
* [Pros](#pros)
* [Cons](#cons)
* [Credits](#credits)


### Updates

July 20, 2019: Thanks to the efforts by [@YakovL](https://github.com/YakovL), timimi has basic support for saving Tiddlywiki Classic too!

### Intro

Timimi is a webextension using native messaging API that allows it to save standalone tiddlywiki files.

Tiddlyfox addon for firefox browser stopped working for post 57 versions of firefox when mozilla switched over to sandbox model, restricting access to the file system. However webextension addons provide APIs which interface with filesystem. This is an addon using native messaging, essentially handing over the contents to a webextension host which does the actual saving. For the end user it means a couple of extra steps to install the addon but once installed, you can save the standalone TW from anywhere in your harddrive without any more interactions, like the original Tiddlyfox addon.

**Note:** This extention only works on tiddlywiki html files. (ie. It will not work with the nodejs server, or similar.)

### Supported browsers


* Firefox >57


### Supported OS

* Debian based systems - Debian, Ubuntu, Elementary, Mint etc
* Arch based systems - Arch Linux, Antergos, Manjaro etc
* Windows 7 and later.



### Installation

Please see: https://ibnishak.github.io/Timimi/

### Usage


Please see: https://ibnishak.github.io/Timimi/

### To do

1. DONE Relative paths for backups
2. DONE Rewrite webextension host in golang
3. DONE Launch simple scripts


### Pros

* Webextension host can be expanded for different purposes depending on personal preferences. For eg:
  * Saving time stamped backups instead of overwrting.
  * Monitor changes to a tiddler and enter it to a local relational database like sqlite.

* 3 different backup strategies
* No need to run a server, does not require continued user interactions

### Cons

1. Cannot be used in android

### Credits

#### Logo credits

Icons made by [Freepik](http://www.freepik.com) from <a href="https://www.flaticon.com/" title="Flaticon">www.flaticon.com</a> is licensed by <a href="http://creativecommons.org/licenses/by/3.0/" title="Creative Commons BY 3.0" target="_blank">CC 3.0 BY</a>




