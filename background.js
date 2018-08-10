// background-script.js

var portFromCS;


// This function is executed on connection
function connected(p) {
  portFromCS = p;
  browser.browserAction.onClicked.addListener(function() {
	  portFromCS.postMessage({greeting: "hi there content script!"});
	  portFromCS.onMessage.addListener(function(m) {
	    console.log("In background script, received message from content script")
	    console.log(m.greeting);
	  });
  });
}

// Actual command to execute the connected function
browser.runtime.onConnect.addListener(connected);
// Pass message to content script on clicking the button. Need response here
// browser.browserAction.onClicked.addListener(function() {
//   portFromCS.postMessage({greeting: "they clicked the button!"});
// });