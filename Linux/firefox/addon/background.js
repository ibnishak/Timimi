var port = browser.runtime.connectNative("timimi");

// port.onMessage.addListener((response) => {
//     console.log("Received: " + response.content);
// });

function handleMessage(request, sender, sendResponse) {
    console.log("Sending native message");
    browser.runtime.sendNativeMessage("timimi", request);
    // port.postMessage(request);
}


browser.runtime.onMessage.addListener(handleMessage);
