// var port = browser.runtime.connectNative("timimi");

// port.onMessage.addListener((response) => {
//     console.log("Received: " + response.content);
// });
function onResponse(response) {
  handleArray(response.Resp);
  if (response.Errors != null) {
    handleArray(response.Errors);
  }
  if (response.Stdout != "") {
    console.log(response.Stdout);
    stdout = response.Stdout;
    browser.tabs
      .query({
        currentWindow: true,
        active: true
      })
      .then(sendMessageToTabs)
      .catch(onError);
    function sendMessageToTabs(tabs) {
      for (let tab of tabs) {
        browser.tabs.sendMessage(tab.id, { stdout: stdout });
      }
      console.log("Timimi: Sending stdout to contentscript");
    }
  }
}

function onError(error) {
  console.log(`Timimi: Native Host Error: ${error}`);
  browser.notifications.create({
    "type": "basic",
    "iconUrl": browser.extension.getURL("../resources/fish.png"),
    "title": "Could not save the Tiddlywiki file",
    "message": `Have you installed the Timimi native application?\nNative Host Error: ${error}`
  });
}

function handleMessage(request, sender, sendResponse) {
  console.log("Timimi: Sending native message");
  var sending = browser.runtime.sendNativeMessage("timimi", request);
  sending.then(onResponse, onError);
  // port.postMessage(request);
}

function handleArray(IncomingArray) {
  dataArray = IncomingArray.map(function(e) {
    return JSON.stringify(e);
  });
  dataString = dataArray.join("\n");
  console.log(dataString);
}

browser.runtime.onMessage.addListener(handleMessage);
