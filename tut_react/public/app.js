const root = document.querySelector("#root");

function detik() {
  const element = /*#__PURE__*/React.createElement("div", null, /*#__PURE__*/React.createElement("h1", null, "Jam sekarang"), new Date().toLocaleTimeString());
  ReactDOM.render(element, root);
}

detik();
setInterval(function () {
  detik();
}, 1000);