var sock=null;
var wsuri ="ws://127.0.0.1:8080/ws"; //這裏的IP假設是局域測試的話。須要換成自己的
window.onload = function(){
    console.log("onload");
    sock = new WebSocket(wsuri);
    sock.onopen=function(e){
        console.log("connected to "+wsuri);
    }
    sock.onclose=function(e){
        console.log("connection closed (" + e.code + ")");
    }
    sock.onmessage=function(e){
        console.log("message received: " + e.data);
        var div = document.getElementById("list")

        var root=  document.createElement("div")
        var imgPhot = document.createElement("img")
        var msggg = document.createElement("p")
        var time = document.createElement("span")
        root.className = "container"
        imgPhot.src = "img/robot.png"
        imgPhot.className = "right"
        imgPhot.alt = "System"
        msggg.innerText = e.data
        time.className = "time-left"
        time.innerText = "12:00"

        root.appendChild(imgPhot)
        root.appendChild(msggg)
        root.appendChild(time)


        div.appendChild(root)

    }
}

// <div className="container">
//     <img src="img/robot.png" alt="System">
//     <p>廣播訊息： 階段1 下注</p>
//     <span className="time-right">11:00</span>
// </div>
