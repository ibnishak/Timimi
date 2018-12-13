function saveOptions(e) {
  e.preventDefault();
  browser.storage.sync.set({
    bstrategy: document.querySelector("#bstrategy").value
  });
}

function restoreOptions() {
  function setCurrentChoice(result) {
    document.querySelector("#bstrategy").value = result.bstrategy || "toh";
  }

  function onError(error) {
    console.log(`Error: ${error}`);
  }

  var getting = browser.storage.sync.get("bstrategy");
  getting.then(setCurrentChoice, onError);
}

document.addEventListener("DOMContentLoaded", restoreOptions);
document.querySelector("form").addEventListener("submit", saveOptions);
