
var myPort = browser.runtime.connect({name:"port-from-cs"});
var idGenerator = 1;

var messageBox = document.getElementById("tiddlyfox-message-box");
if(!messageBox) {
	messageBox = document.createElement("div");
	messageBox.id = "tiddlyfox-message-box";
	messageBox.style.display = "none";
	document.body.appendChild(messageBox);
}

myPort.onMessage.addListener(function(m) {
  console.log(m.message);
	// Attach the event handler to the message box
	messageBox.addEventListener("tiddlyfox-save-file",onSaveTiddlyWiki,false);

	function onSaveTiddlyWiki(event) {
	// Get the details from the message
	var messageElement = event.target,
		path = messageElement.getAttribute("data-tiddlyfox-path"),
		content = messageElement.getAttribute("data-tiddlyfox-content"),
		backupPath = messageElement.getAttribute("data-tiddlyfox-backup-path"),
		messageId = "tiddlywiki-save-file-response-" + idGenerator++;
		console.log("Saving to " + path);
		myPort.postMessage({path: path, messageId: messageId, content: content, backupPath:backupPath});
		messageElement.parentNode.removeChild(messageElement);
		console.log("Saving huuuge success");
    	var event = document.createEvent("Events");
		event.initEvent("tiddlyfox-have-saved-file",true,false);
		event.savedFilePath = path;
		messageElement.dispatchEvent(event);
		}
});