window.onload = function () {
  const upButton = document.getElementById("up");
  const downButton = document.getElementById("down");
  upButton.onclick = function () {
    var oReq = new XMLHttpRequest();
    oReq.addEventListener("load", reqListener);
    oReq.open("GET", "/up");
    oReq.send();
  }

  downButton.onclick = function () {
    var oReq = new XMLHttpRequest();
    oReq.addEventListener("load", reqListener);
    oReq.open("GET", "/down");
    oReq.send();
  }

  window.setInterval(function() {
    var oReq = new XMLHttpRequest();
    oReq.addEventListener("load", reqListener);
    oReq.open("GET", "/get");
    oReq.send();
  }, 400);
}

function reqListener() {
  const volumeDisplay = document.getElementById("volume")
  volumeDisplay.innerHTML = this.responseText
}
