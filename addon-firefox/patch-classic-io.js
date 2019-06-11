window.mozillaLoadFile = function(path) {
  try {
    // Just read the file synchronously
    var xhReq = new XMLHttpRequest();
    xhReq.open("GET", "file:///" + escape(path), false);
    xhReq.send(null);
    return xhReq.responseText;
  } catch(ex) {
    return false;
  }
};

window.mozillaSaveFile = function(path, content) {
  var messageBox = document.getElementById("tiddlyfox-message-box");
  if(!messageBox) return false;

  // Create the message element and put it into the message box
  var message = document.createElement("div");
  message.setAttribute("data-tiddlyfox-path", path);
  message.setAttribute("data-tiddlyfox-content", content);
  messageBox.appendChild(message);

  // Create and dispatch the custom event to the extension
  var event = document.createEvent("Events");
  event.initEvent("tiddlyfox-save-file", true, false);
  message.dispatchEvent(event);

  return true;
};

// not necessary to track the exact version, but important to change when the overwritten methods are updated
window.timimiClassicVersion = "2.0.2";