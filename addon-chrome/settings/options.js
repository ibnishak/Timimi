function saveOptions(e) {
  e.preventDefault();
  browser.storage.sync.set({
    bpath: document.querySelector("#bpath").value,
    bstrategy: document.querySelector("#bstrategy").value,
    tohrecent: document.querySelector("#tohrecent").value,
    tohlevel: document.querySelector("#tohlevel").value,
    psint: document.querySelector("#psint").value,
    tint: document.querySelector("#tint").value
  });
  
  browser.notifications.create({
    "type": "basic",
    "title": "Timimi preferences SAVED",
    "iconUrl": browser.runtime.getURL("icons/index.svg"),
    "message": "Please reload the TW5 browser tabs for the new preferences to take effect"
  });
}

function restoreOptions() {
  var getting = browser.storage.sync.get();
  getting.then(setCurrentChoice, onError);

  function setCurrentChoice(result) {
    document.querySelector("#bpath").value = result.bpath || "";
    document.querySelector("#bstrategy").value = result.bstrategy || "none";
    document.querySelector("#tohrecent").value = result.tohrecent || "5";
    document.querySelector("#tohlevel").value = result.tohlevel || "8";
    document.querySelector("#psint").value = result.psint || "10";
    document.querySelector("#tint").value = result.tint || "4";
  }

  function onError(error) {
    console.log(`Error: ${error}`);
  }
}

document.addEventListener("DOMContentLoaded", restoreOptions);
document.querySelector("form").addEventListener("submit", saveOptions);

var slider = document.getElementById("tohlevel");
var output = document.getElementById("rangeout");
output.innerHTML = slider.value;

slider.oninput = function() {
  output.innerHTML = this.value;
};



// Ensuring that only integers are entered to tohrecent and psint

function setInputFilter(textbox, inputFilter) {
  [
    "input",
    "keydown",
    "keyup",
    "mousedown",
    "mouseup",
    "select",
    "contextmenu",
    "drop"
  ].forEach(function(event) {
    textbox.addEventListener(event, function() {
      if (inputFilter(this.value)) {
        this.oldValue = this.value;
        this.oldSelectionStart = this.selectionStart;
        this.oldSelectionEnd = this.selectionEnd;
      } else if (this.hasOwnProperty("oldValue")) {
        this.value = this.oldValue;
        this.setSelectionRange(this.oldSelectionStart, this.oldSelectionEnd);
      }
    });
  });
}

setInputFilter(document.getElementById("tohrecent"), function(value) {
  return /^-?\d*$/.test(value);
});

setInputFilter(document.getElementById("psint"), function(value) {
  return /^-?\d*$/.test(value);
});
