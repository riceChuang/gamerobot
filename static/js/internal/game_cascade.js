import * as common from "../common/index.js";
import {socket} from "./websocket.js";


var sock= null;

function clickGame(select) {
    var game = gameList[select.value];
    var container = document.getElementById("param")
    var roomContainer = document.getElementById("roomSelect")
    while (container.lastElementChild) {
        container.removeChild(container.lastElementChild)
    }
    while (roomContainer.lastElementChild){
        roomContainer.removeChild(roomContainer.lastElementChild)
    }
    new Map(Object.entries(game["Buttons"])).forEach((value, key) => {
        var button = document.createElement("button");
        button.innerText = key
        button.className = "btn btn-primary"
        button.data = value
        button.setAttribute("onclick", "window.index.sendMessageBtn(this.data)"); // added line
        container.appendChild(button);
    })

    Object.values(game["RoomType"]).forEach(value => {
        var room = document.createElement("option");
        room.innerText = value
        room.value = value
        roomContainer.appendChild(room);
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
    var today = new Date
    time.innerText = today.getTime()
    root.appendChild(imgPhot)
    root.appendChild(msggg)
    root.appendChild(time)
    div.appendChild(root)

    //clear
    document.getElementById(`msg`).value="";
    var data = who + "  say:  " + msg;
    sock.send(data);
}


function sendMessageBtn(data) {
    var sendMsg = JSON.stringify({SClassID: data.SClassID, BClassID: data.BClassID})
    var div = document.getElementById("list")
    var root=  document.createElement("div")
    var imgPhot = document.createElement("img")
    var msggg = document.createElement("p")
    var time = document.createElement("span")
    root.className = "container"
    imgPhot.src = "img/client.png"
    imgPhot.alt = "System"
    msggg.innerText = sendMsg
    time.className = "time-right"
    var today = new Date
    time.innerText = today.getDate()
    root.appendChild(imgPhot)
    root.appendChild(msggg)
    root.appendChild(time)
    div.appendChild(root)

    sock.send(sendMsg);
}

function loginRequest() {

    var envReq = document.getElementById(`envSelect`).value;
    var gameID = Number(document.getElementById(`gameSelect`).value);
    var roomType = document.getElementById(`roomSelect`).value;

    post(common.router.loginUrl, {
        gameid: gameID,
        env: envReq,
        roomtype: roomType,
    }).then(function (resp){
        if (resp.status == 200){
            resp.json().then(data => {
                console.log(data)
                connectWS(data.data)
                common.addSuccess()
            })
        }else {
            resp.json().then(data => {
                common.addFail(JSON.stringify(data))
            })
        }
    })
}

function  connectWS(data) {

    var accountID = document.getElementById(`right-account`);
    accountID.innerText = data.account
    var evnID = document.getElementById(`right-env`);
    evnID.innerText = data.env
    var gameID = document.getElementById(`right-game`);
    gameID.innerText = data.gamename

    sock = new socket({
        //網址（端口是我下面的服務器的端口）
        'url': common.router.wsUrl ,
        'urlParam': "connID=" + data.conn,
    })

    console.log(sock)
    //關閉事件
    sock.onclose((event)=>{
        console.log("connection closed (" + e.code + ")");
    })
    //心跳事件
    sock.onheartbeat(()=>{
        console.log('heartbeat')
    })
    //接收信息
    sock.onmessage((data)=>{
        if (data == "pong") {
            console.log("ping pong msg: " + data);
            return
        }

        console.log("message received: " + data);
        var div = document.getElementById("list")
        var root=  document.createElement("div")
        var imgPhot = document.createElement("img")
        var msggg = document.createElement("p")
        var time = document.createElement("span")
        root.className = "container"
        imgPhot.src = "img/robot.png"
        imgPhot.className = "right"
        imgPhot.alt = "System"
        msggg.innerText = data
        time.className = "time-left"
        time.innerText = "12:00"

        root.appendChild(imgPhot)
        root.appendChild(msggg)
        root.appendChild(time)
        div.appendChild(root)
        div.scrollTop = div.scrollHeight;
    })
    //心跳事件
    sock.onreconnect(()=>{
        console.log('reconnect')
    })
}


function unpack(str) {
    var bytes = [];
    for(var i = 0; i < str.length; i++) {
        var char = str.charCodeAt(i);
        bytes.push(char >>> 8);
        bytes.push(char & 0xFF);
    }
    return bytes;
}

export {
    loginRequest,
    send,
    clickGame,
    sendMessageBtn
}