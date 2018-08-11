var portFromCS;

function connected(p) {
    portFromCS = p;
    portFromCS.postMessage({ message: "Initiate" });
    portFromCS.onMessage.addListener(function(m) {
        console.log("Saving to " + m.path);
        var sending = browser.runtime.sendNativeMessage(
            "keeptw", m);
        sending.then(onResponse, onError);
    });
    portFromCS.onDisconnect.addListener(disconnected);
}

function disconnected() {
    console.log("Not a local Tiddlywiki file");
}


browser.runtime.onConnect.addListener(connected);

function onResponse(response) {
    console.log("Saved Successfully");
}

function onError(error) {
    console.log(`Error: ${error}`);
}