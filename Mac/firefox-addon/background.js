
function handleMessage(request, sender, sendResponse) {
    console.log("Sending native message");
    browser.runtime.sendNativeMessage(
        "timimi", request);
}


browser.runtime.onMessage.addListener(handleMessage);
