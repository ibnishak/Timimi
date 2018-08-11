// background-script.js

var portFromCS;


// This function is executed on connection
function connected(p) {
  portFromCS = p;
  // browser.browserAction.onClicked.addListener(function() {
	 //  portFromCS.postMessage({greeting: "hi there content script!"});
	 //  portFromCS.onMessage.addListener(function(m) {
	 //    console.log("In background script, received message from content script")
	 //    console.log(m.greeting);
	 //  });
  // });
}


function onResponse(response) {
    console.log("Received " + response);
    browser.notifications.create({
	"type": "basic",
	"title": response,
	"message": "Link added to shiori",
	"iconUrl": browser.extension.getURL("icons/shiori.png")
  });
}

function onError(error) {
    console.log(`Error: ${error}`);
}

// Actual command to execute the connected function
browser.runtime.onConnect.addListener(connected);
// Pass message to content script on clicking the button. Need response here
browser.browserAction.onClicked.addListener(function() {
  portFromCS.postMessage({msg: "Calling save modules"});
	portFromCS.onMessage.addListener(function(m) {
    console.log("In background script, received message from content script")
    console.log(m);
    // console.log(m.docbody);
    // console.log(m.dochead);
    // console.log(m.docurl);
 //    var sending = browser.runtime.sendNativeMessage(
	//      "keeptw",
	//      m);
	// sending.then(onResponse, onError);
  });
});

