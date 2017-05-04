package main

const (
	indexHtml = `<html>
<head>
  <title>File watcher terminal</title>
</head>

<style>
  body {
    background: #000000;
    color: #efefef;
    padding: 16px;
  }

  #container {
    text-align: left;
    line-height: 1.4;
    width: 100%;
    height: 100%;
    font-size: 16px;
    overflow-y: scroll;
    font-family: monospace;
    white-space: nowrap;
  }

  #text {
    background: #000000;
    color: #efefef;
    width: 100%;
    padding: 10px 0 0;
    font-size: 20px;
    border: none;
  }

  #text:focus,
  #text:active {
    border: none;
    outline: none;
  }
</style>

<body>
<div id="container">
  <div id="terminal"></div>
  <input autofocus placeholder="$" id="text" type="text">
</div>

<script>
  var url = "ws://" + window.location.host + "/watcher";
  var ws = new WebSocket(url);
  var entered = [];
  var offset = -1;
  var container = document.getElementById("container");
  var terminal = document.getElementById("terminal");
  var text = document.getElementById("text");

  var add = function (text, from) {
    var sender = from ? from + " " : "";

    terminal.innerText += sender + text + "\n";
    container.scrollTop = container.scrollHeight;
  };

  var save = function(cmd) {
    entered = [cmd].concat(entered);
  };

  var prev = function() {
    if (offset === entered.length - 1) {
      return
    }

    offset += 1;
    text.value = entered[offset];
  };

  var next = function() {
    if (offset === -1) {
      text.value = "";
      return
    }

    offset -= 1;
    text.value = entered[offset] || "";
  };

  var exec = function() {
    save(text.value);
    add(text.value, "$");
    ws.send(text.value);
    text.value = "";
    offset = -1;
  };

  ws.onmessage = function (msg) {
    add(msg.data);
  };

  text.onkeydown = function (e) {
    if (e.keyCode === 38) {
      e.preventDefault();
      prev();
    }

    if (e.keyCode === 40) {
      e.preventDefault();
      next();
    }

    if (e.keyCode === 13 && text.value !== "") {
      exec();
    }
  };

</script>
</body>
</html>`
)
