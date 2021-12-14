
function clickGame(select) {
    var game = gameList[select.value];
    var container = document.getElementById("abc")

    while (container.lastElementChild) {
        container.removeChild(container.lastElementChild)
    }

    Object.keys(game["Buttons"]).forEach(e => {
        console.log(e);
        var button = document.createElement("button");
        button.innerText = e
        container.appendChild(button);
    })

}

function send() {
    var who = "Client"
    var msg = document.getElementById(`msg`).value;
    if (msg.length == 0){
        document.getElementById(`msg`).focus();
        return
    }

    var div = document.getElementById("list")

    var root=  document.createElement("div")
    var imgPhot = document.createElement("img")
    var msggg = document.createElement("p")
    var time = document.createElement("span")
    root.className = "container"
    imgPhot.src = "img/client.png"
    imgPhot.alt = "System"
    msggg.innerText = msg
    time.className = "time-right"
    time.innerText = "12:00"


    root.appendChild(imgPhot)
    root.appendChild(msggg)
    root.appendChild(time)

    div.appendChild(root)

    //clear
    document.getElementById(`msg`).value="";
    var data = who + "  say:  " + msg + "<br/><br/>";
    sock.send(data);
}

