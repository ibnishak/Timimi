function saveOptions(e) {
  e.preventDefault();
  browser.storage.sync.set({
    backup: document.querySelector("#backup").value,
    bpath: document.querySelector("#bpath").value,
    bpath: document.querySelector("#bpath").value,
    bstrategy: document.querySelector("#bstrategy").value,
    tohrecent: document.querySelector("#tohrecent").value,
    tohlevel: document.querySelector("#tohlevel").value,
    psint: document.querySelector("#psint").value,
    exec: document.querySelector("#exec").value,
    tint: document.querySelector("#tint").value
  });
}

function restoreOptions() {
  var getting = browser.storage.sync.get();
  getting.then(setCurrentChoice, onError);

  function setCurrentChoice(result) {
    document.querySelector("#backup").value = result.backup || "yes";
    document.querySelector("#bpath").value = result.bpath || "";
    document.querySelector("#bstrategy").value = result.bstrategy || "toh";
    document.querySelector("#tohrecent").value = result.tohrecent || "5";
    document.querySelector("#tohlevel").value = result.tohlevel || "8";
    document.querySelector("#psint").value = result.psint || "10";
    document.querySelector("#exec").value = result.exec || "yes";
    document.querySelector("#tint").value = result.tint || "4";
  }

  function onError(error) {
    console.log(`Error: ${error}`);
  }
}

document.addEventListener("DOMContentLoaded", restoreOptions);
// document.addEventListener("DOMContentLoaded", selectCh);
document.querySelector("form").addEventListener("submit", saveOptions);

var slider = document.getElementById("tohlevel");
var output = document.getElementById("rangeout");
output.innerHTML = slider.value;

slider.oninput = function() {
  output.innerHTML = this.value;
};

// function selectCh() {
//   var sel = document.getElementById("bstrategy");
//   document.getElementById("toh").style.display = "none";
//   document.getElementById("psave").style.display = "none";
//   document.getElementById(sel.value).style.display = "block";
// }

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
