var port = browser.runtime.connectNative("timimi");
var bpath;
var getting = browser.storage.sync.get("color");
getting.then(onGot, onError);
// port.onMessage.addListener((response) => {
//     console.log("Received: " + response.content);
// });
function onError(error) {
  console.log(`Error: ${error}`);
}

function onGot(item) {
  if (item.bpath) {
    bpath = item.bpath;
  }
}

function handleMessage(request, sender, sendResponse) {
  console.log("Sending native message");
  request.bpath = bpath;
  browser.runtime.sendNativeMessage("timimi", request);
  // port.postMessage(request);
}

browser.runtime.onMessage.addListener(handleMessage);
