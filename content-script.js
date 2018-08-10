// content-script.js
// Message posted as soon as connected
var myPort = browser.runtime.connect({name:"port-from-cs"});
myPort.postMessage({greeting: "hello from content script"});
// This will be carried out in response to message from background js. Including onclick
myPort.onMessage.addListener(function(m) {
  console.log("In content script, received message from background script: ");
  console.log(m.greeting);
  myPort.postMessage({greeting: "Now we are sending messages back"});
});

// document.body.addEventListener("click", function() {
//   myPort.postMessage({greeting: "they clicked the page!"});
// });