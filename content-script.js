// content-script.js
// Message posted as soon as connected
var idGenerator = 1;
var myPort = browser.runtime.connect({name:"port-from-cs"});
myPort.postMessage({greeting: "hello from content script"});
// This will be carried out in response to message from background js. Including onclick
myPort.onMessage.addListener(function(m) {
  console.log(m.msg);
  console.log("Responding to save");
  myPort.postMessage({docbody: document.body.outerHTML, docurl: document.URL, dochead: document.head.outerHTML});
});

var messageBox = document.getElementById("tiddlyfox-message-box");
	if(!messageBox) {
		messageBox = document.createElement("div");
		messageBox.id = "tiddlyfox-message-box";
		messageBox.style.display = "none";
		document.body.appendChild(messageBox);
	}
	// Attach the event handler to the message box
messageBox.addEventListener("tiddlyfox-save-file",onSaveTiddlyWiki,false);
// document.body.addEventListener("click", function() {
//   myPort.postMessage({greeting: "they clicked the page!"});
// });


function onSaveTiddlyWiki(event) {
	// Get the details from the message
	var messageElement = event.target,
		path = messageElement.getAttribute("data-tiddlyfox-path"),
		content = messageElement.getAttribute("data-tiddlyfox-content"),
		backupPath = messageElement.getAttribute("data-tiddlyfox-backup-path"),
		messageId = "tiddlywiki-save-file-response-" + idGenerator++;
		console.log("yay");
		console.log(path);
		myPort.postMessage({greeting: "hello from FOX script"});
		}
