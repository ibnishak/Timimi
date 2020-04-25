var idGenerator = 1;
var timimisettings = {};
var tlast = new Date();

// If and only if the local file is TW-Classic(old) file, inject the patch enabling it to be saved with tiddlyfox-derived savers
function injectExtensionScript(path) {
  var script = document.createElement('script');
  script.src = chrome.extension.getURL(path); 
  (document.head || document.documentElement).appendChild(script);
  script.onload = script.remove;
}

// Checking if the active tab is a local tiddlywiki file
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

// Loading settings from chrome storage
function getSettings() {
  chrome.storage.sync.get(["bpath", "bstrategy", "tohrecent", "tohlevel", "psint", "tint", "fint" ], function(result) {
    timimisettings = result;
  });
}

// ----- START HERE -------- //

// Check if the local file is a TW
var checkTWResults = checkTW();

// If it is a TW-Classic, then do @YakovL magic to save it
if (checkTWResults.isTiddlyWikiClassic && checkTWResults.isLocalFile) {
  getSettings();
  injectExtensionScript('patch-classic-io.js');
}

// If it is a TW-5, then do extract settings and save it.
if (checkTWResults.isTiddlyWiki && checkTWResults.isLocalFile) {
  // Extract settings
  getSettings();

  // Build an area in the TW5 where TW5 can notify a save event. Code credits to @Jermolene
  var messageBox = document.getElementById("tiddlyfox-message-box");
  if (!messageBox) {
    messageBox = document.createElement("div");
    messageBox.id = "tiddlyfox-message-box";
    messageBox.style.display = "none";
    document.body.appendChild(messageBox);
  }
  // Attach the event handler to the message box
  messageBox.addEventListener("tiddlyfox-save-file", onSaveTiddlyWiki, false);

  function onSaveTiddlyWiki(event) {
    tbackup = "false";
    if (timimisettings.bstrategy == "timed") {
      var now = new Date();
      var diffMs = now - tlast;
      var diffMins = Math.round(((diffMs % 86400000) % 3600000) / 60000);
      if (diffMins >= timimisettings.tint) {
        tbackup = "true";
        tlast = now;
        console.log("Timimi: Creating Timed Backup");
      }
    }

    // Get the details from the message
    var messageElement = event.target,
      path = messageElement.getAttribute("data-tiddlyfox-path"),
      content = messageElement.getAttribute("data-tiddlyfox-content"),
      backupPath = messageElement.getAttribute("data-tiddlyfox-backup-path"),
      messageId = "tiddlywiki-save-file-response-" + idGenerator++;

    // Send the details to background script. Not using port because we need a promise and port.postMessage is not a promise
    chrome.runtime.sendMessage({
      path: path,
      messageId: idGenerator,
      content: content,
      backupPath: backupPath,
      backup: timimisettings.backup,
      bpath: timimisettings.bpath,
      bstrategy: timimisettings.bstrategy,
      tohrecent: timimisettings.tohrecent,
      tohlevel: timimisettings.tohlevel,
      psint: timimisettings.psint,
      fint: timimisettings.fint,
      tbackup: tbackup
    }, handleSent);
    

    function handleSent(message) {
      messageElement.parentNode.removeChild(messageElement);
      console.log("Message ID is " + messageId);
      var event = document.createEvent("Events");
      event.initEvent("tiddlyfox-have-saved-file", true, false);
      event.savedFilePath = path;
      messageElement.dispatchEvent(event);
    }
  
  }
}


