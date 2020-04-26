function onResponse(response) {
  console.log("Timimi: Native Host Responded without errors");
}

function onError(error) {
  console.log(`Timimi: Native Host Error: ${error}`);
  browser.notifications.create({
    "type": "basic",
    "title": "Timimi save FAILED",
    "iconUrl": browser.runtime.getURL("icons/index.svg"),
    "message": error.toString()
  });
}

function handleMessage(request, sender, sendResponse) {
  console.log("Timimi: Sending native message");
  var sending = browser.runtime.sendNativeMessage("timimi", request);
  sending.then(onResponse, onError);
}


browser.runtime.onMessage.addListener(handleMessage);

function handleInstalled(details) {
  console.log(details.reason);
  browser.tabs.create({
    url: "https://ibnishak.github.io/Timimi/#Important%3A%20Post%20Update%2FInstallation%20instructions"
  });
}

browser.runtime.onInstalled.addListener(handleInstalled);

function openPage() {
  browser.tabs.create({
    url: "https://ibnishak.github.io/Timimi"
  });
}

browser.browserAction.onClicked.addListener(openPage);