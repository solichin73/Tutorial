const root = document.querySelector("#root");

function detik(){
  const element = (
   <div>
     <h1>Jam sekarang</h1>
     {new Date().toLocaleTimeString()}
    </div>
  )

  ReactDOM.render(element, root);
}

detik();

setInterval(function(){
  detik();
},1000);
