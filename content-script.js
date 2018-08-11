
var myPort = browser.runtime.connect({name:"port-from-cs"});

myPort.onMessage.addListener(function(m) {
  console.log("In content script, received message from background script: ");
  console.log(m.greeting);
  myPort.postMessage({greeting: "hello from content script"});
});