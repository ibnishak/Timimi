var portFromCS;

function connected(p) {
  portFromCS = p;
  portFromCS.postMessage({greeting: "hi there content script!"});
  portFromCS.onMessage.addListener(function(m) {
    console.log("In background script, received message from content script")
    console.log(m.path);
    console.log(m.messageId);
  var sending = browser.runtime.sendNativeMessage(
       "keeptw",m);
  sending.then(onResponse, onError);
  });
}

browser.runtime.onConnect.addListener(connected);

function onResponse(response) {
    console.log("Received " + response);
}

function onError(error) {
    console.log(`Error: ${error}`);
}