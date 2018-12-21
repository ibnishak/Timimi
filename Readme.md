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
<a href="https://GitHub.com/ibnishak/timimi/stargazers/" class="rich-diff-level-one">
	<img src="https://camo.githubusercontent.com/ea01a1feb6d3db0cbe7c102fe98407f05fc55413/68747470733a2f2f696d672e736869656c64732e696f2f6769746875622f73746172732f69626e697368616b2f74696d696d692e7376673f7374796c653d736f6369616c266c6162656c3d53746172266d61784167653d32353932303030" alt="GitHub stars" data-canonical-src="https://img.shields.io/github/stars/ibnishak/timimi.svg?style=social&amp;label=Star&amp;maxAge=2592000" style="max-width:100%;">
</a>
<a href="http://makeapullrequest.com" rel="nofollow" class="rich-diff-level-one">
	<img src="https://camo.githubusercontent.com/a34cfbf37ba6848362bf2bee0f3915c2e38b1cc1/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f5052732d77656c636f6d652d627269676874677265656e2e7376673f7374796c653d666c61742d737175617265" alt="PRs Welcome" data-canonical-src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" style="max-width:100%;">
</a>
[![HitCount](http://hits.dwyl.com/ibnishak/Timimi.svg)](http://hits.dwyl.com/ibnishak/Timimi)
</p>


<br/><br/><br/>


* [Intro](#intro)
* [Supported browsers](#supported-browsers)
* [Supported OS](#supported-os)
* [Installation](#installation)
* [To do](#to-do)
* [Pros](#pros)
* [Cons](#cons)
* [Credits](#credits)


### Intro

Timimi is a webextension using native messaging API that allows it to save standalone tiddlywiki files.

Tiddlyfox addon for firefox browser stopped working for post 57 versions of firefox when mozilla switched over to sandbox model, restricting access to the file system. However webextension addons provide APIs which interface with filesystem. This is an addon using native messaging, essentially handing over the contents to a webextension host which does the actual saving. For the end user it means a couple of extra steps to install the addon but once installed, you can save the standalone TW from anywhere in your harddrive without any more interactions, like the original Tiddlyfox addon.



### Supported browsers


* Firefox >57

### Supported OS

* Debian based systems - Debian, Ubuntu, Elementary, Mint etc
* Arch based systems - Arch Linux, Antergos, Manjaro etc
* Windows 7 and later.



### Installation

#### Firefox
Documentation Pending.


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




