chrome.runtime.onMessage.addListener(connect);


function connect(request,_,sendResponse) {
  var hostName = "timimi";
  port = chrome.runtime.connectNative(hostName);
  port.postMessage(request);
  sendResponse("Save handed over to native host");
}

