var portFromCS;

function connected(p) {
    portFromCS = p;
    portFromCS.postMessage({ message: "Initiate" });
    portFromCS.onDisconnect.addListener(disconnected);
}

function disconnected() {
    console.log("Not a local Tiddlywiki file");
}

function onResponse(response) {
    sendResponse({ response: "Saved Successfully" });
    //console.log("Background script sends success response to content script")
}

function onError(error) {
    console.log(`Error: ${error}`);
}

function handleMessage(request, sender, sendResponse) {
    console.log("Sending native message");
    var sending = browser.runtime.sendNativeMessage(
        "keeptw", request);
    sending.then(onResponse, onError);
}


browser.runtime.onConnect.addListener(connected);
browser.runtime.onMessage.addListener(handleMessage);
