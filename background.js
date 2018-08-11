var portFromCS;

function connected(p) {
  portFromCS = p;
  portFromCS.postMessage({greeting: "hi there content script!"});
  portFromCS.onMessage.addListener(function(m) {
    console.log("In background script, received message from content script")
    console.log(m.greeting);
  });
}

browser.runtime.onConnect.addListener(connected);


// On connect  -> bg posts hi there, bg receieves a reply - bg consoles it
// then the power moves to content script - it consoles and posts message
// once bg receives the message it posted - it consoles it