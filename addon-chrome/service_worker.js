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

function handleMessage(request, sender, sendResponse) {
  console.log("Timimi: Sending native message");
  sendResponse({message: "Timimi: Data received in background.js"}); // Necessary to keep Port closed erro away
  chrome.runtime.sendNativeMessage("timimi", request, onResponse);

}

chrome.runtime.onMessage.addListener(handleMessage);

chrome.runtime.onInstalled.addListener(details => {
  if (details.reason == "chrome_update") return;
  chrome.tabs.create({
    url: "https://ibnishak.github.io/Timimi/#Important%3A%20Post%20Update%2FInstallation%20instructions"
  });
});

chrome.action.onClicked.addListener(function() {
  chrome.runtime.openOptionsPage();
});