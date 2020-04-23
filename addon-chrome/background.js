function onResponse(response) {
  if (chrome.runtime.lastError) {
    onError(chrome.runtime.lastError.message)
  } else {
    console.log("Timimi: Native Host Responded without errors");
  }
}

function onError(error) {
  console.log(`Timimi: Native Host Error: ${error}`);
  chrome.notifications.create({
    "type": "basic",
    "title": "Timimi save FAILED",
    "iconUrl": chrome.runtime.getURL("icons/icon16.png"),
    "message": "Error on contacting timimi host"
  });
}

function handleMessage(request) {
  console.log("Timimi: Sending native message");
  chrome.runtime.sendNativeMessage("timimi", request, onResponse);
}


chrome.runtime.onMessage.addListener(handleMessage);