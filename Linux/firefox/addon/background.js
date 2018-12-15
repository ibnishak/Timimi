// var port = browser.runtime.connectNative("timimi");

// port.onMessage.addListener((response) => {
//     console.log("Received: " + response.content);
// });
function onResponse(response) {
  console.log("TImimi: Native Host: " + response.content);
}

function onError(error) {
  console.log(`Error: ${error}`);
}

function handleMessage(request, sender, sendResponse) {
  console.log("Sending native message");
  var sending = browser.runtime.sendNativeMessage("timimi", request);
  sending.then(onResponse, onError);
  // port.postMessage(request);
}

browser.runtime.onMessage.addListener(handleMessage);
