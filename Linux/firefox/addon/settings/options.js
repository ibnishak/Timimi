function saveOptions(e) {
  e.preventDefault();
  browser.storage.sync.set({
    backup: document.querySelector("#backup").value,
    bpath: document.querySelector("#bpath").value,
    bstrategy: document.querySelector("#bstrategy").value,
    tohrecent: document.querySelector("#tohrecent").value,
    tohlevel: document.querySelector("#tohlevel").value,
    psint: document.querySelector("#psint").value
  });
}

function restoreOptions() {
  function setCurrentChoice(result) {
    document.querySelector("#backup").value = result.backup || "false";
    document.querySelector("#bpath").value = result.bpath || ".";
    document.querySelector("#bstrategy").value = result.bstrategy || "toh";
    document.querySelector("#tohrecent").value = result.tohrecent || "5";
    document.querySelector("#tohlevel").value = result.tohlevel || "8";
    document.querySelector("#psint").value = result.psint || "10";
  }

  function onError(error) {
    console.log(`Error: ${error}`);
  }

  var getting = browser.storage.sync.get();
  getting.then(setCurrentChoice, onError);
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
