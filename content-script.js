
var myPort = browser.runtime.connect({name:"port-from-cs"});
var idGenerator = 1;

myPort.onMessage.addListener(function(m) {
  console.log("In content script, received message from background script: ");
  console.log(m.greeting);
  var messageBox = document.getElementById("tiddlyfox-message-box");
	if(!messageBox) {
		messageBox = document.createElement("div");
		messageBox.id = "tiddlyfox-message-box";
		messageBox.style.display = "none";
		document.body.appendChild(messageBox);
	}
	// Attach the event handler to the message box
	messageBox.addEventListener("tiddlyfox-save-file",onSaveTiddlyWiki,false);
	function onSaveTiddlyWiki(event) {
	// Get the details from the message
	var messageElement = event.target,
		path = messageElement.getAttribute("data-tiddlyfox-path"),
		content = messageElement.getAttribute("data-tiddlyfox-content"),
		backupPath = messageElement.getAttribute("data-tiddlyfox-backup-path"),
		messageId = "tiddlywiki-save-file-response-" + idGenerator++;
		console.log("yay");
		console.log(path);
		myPort.postMessage({path: path, messageId: messageId, content: content, backupPath:backupPath});
		}
});