//var twport = browser.runtime.connect({ name: "port-from-cs" });
var idGenerator = 1;
var backup, bpath, bstrategy, tohrecent, tohlevel, psint;
var getting = browser.storage.sync.get();
getting.then(onGot, syncError);

function syncError(error) {
  console.log(`Error in getting values from browser storage: ${error}`);
}

function onGot(item) {
  backup = item.backup;
  bpath = item.bpath;
  bstrategy = item.bstrategy;
  tohrecent = item.tohrecent;
  tohlevel = item.tohlevel;
  psint = item.psint;
  if (backup == "true") {
    console.log("Timimi: Backups enabled");
    console.log("Timimi: Backup method -" + bstrategy);
  }
}

// Checking if the active tab is a  local tiddlywiki file
function checkTW() {
  var results = {};
  // Test for TiddlyWiki Classic
  var versionArea = document.getElementById("versionArea");
  results.isTiddlyWikiClassic =
    document.location.protocol === "file:" &&
    document.getElementById("storeArea") &&
    (versionArea && /TiddlyWiki/.test(versionArea.textContent));
  // Test for TiddlyWiki 5
  var metaTags = document.getElementsByTagName("meta");
  for (var t = 0; t < metaTags.length; t++) {
    if (
      metaTags[t].name === "application-name" &&
      metaTags[t].content === "TiddlyWiki"
    ) {
      results.isTiddlyWiki5 = true;
    }
  }
  results.isTiddlyWiki = results.isTiddlyWikiClassic || results.isTiddlyWiki5;
  // Test for file URI
  if (document.location.protocol === "file:") {
    results.isLocalFile = true;
  }
  return results;
}

var checkTWResults = checkTW();

if (checkTWResults.isTiddlyWiki5 && checkTWResults.isLocalFile) {
  var messageBox = document.getElementById("tiddlyfox-message-box");
  if (!messageBox) {
    messageBox = document.createElement("div");
    messageBox.id = "tiddlyfox-message-box";
    messageBox.style.display = "none";
    document.body.appendChild(messageBox);
  }
  // Listen to initiate message from background script
  // Attach the event handler to the message box
  messageBox.addEventListener("tiddlyfox-save-file", onSaveTiddlyWiki, false);

  function onSaveTiddlyWiki(event) {
    // Get the details from the message
    var messageElement = event.target,
      path = messageElement.getAttribute("data-tiddlyfox-path"),
      content = messageElement.getAttribute("data-tiddlyfox-content"),
      backupPath = messageElement.getAttribute("data-tiddlyfox-backup-path"),
      messageId = "tiddlywiki-save-file-response-" + idGenerator++;
    // Send the details to background script. Not using port because we need a promise and port.postMessage is not a promise
    var sending = browser.runtime.sendMessage({
      path: path,
      messageId: messageId,
      content: content,
      backupPath: backupPath,
      backup: backup,
      bpath: bpath,
      bstrategy: bstrategy,
      tohrecent: tohrecent,
      tohlevel: tohlevel,
      psint: psint
    });
    sending.then(handleResponse, handleError);

    function handleResponse(message) {
      messageElement.parentNode.removeChild(messageElement);
      console.log("Message ID is " + messageId);
      var event = document.createEvent("Events");
      event.initEvent("tiddlyfox-have-saved-file", true, false);
      event.savedFilePath = path;
      messageElement.dispatchEvent(event);
    }

    function handleError() {
      console.log(`Error: ${error}`);
    }
  }
}
