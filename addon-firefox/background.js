function onResponse(response) {
  console.log("Timimi: Native Host Responded without errors");
}

function onError(error) {
  console.log(`Timimi: Native Host Error: ${error}`);
  browser.notifications.create({
    "type": "basic",
    "title": "Timimi save FAILED",
    "message": error.toString()
  });
}

function handleMessage(request, sender, sendResponse) {
  console.log("Timimi: Sending native message");
  var sending = browser.runtime.sendNativeMessage("timimi", request);
  sending.then(onResponse, onError);
}


browser.runtime.onMessage.addListener(handleMessage);
