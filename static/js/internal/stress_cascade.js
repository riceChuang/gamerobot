import * as common from "../common/index.js";
import {socket} from "./websocket.js";

var settings = new Map();
var showSettings = new Map();
var sock = null;
var count = 1;
//啟動壓測
function startStress() {
    if (settings.size === 0) {
        common.addFail("no data")
        return
    }

    var settingsObj = mapToObj(settings)
    // console.log(JSON.stringify(settingsObj))

    post(common.router.startRobotUrl, settingsObj).then(function (resp){
        if (resp.status == 200){
            resp.json().then(data => {
                console.log(data)
                connectWS(data)
                common.addSuccess()
            })
        }else {
            resp.json().then(data => {
                common.addFail(JSON.stringify(data))
            })
        }
    })
}

//關閉壓測機器人
function stopRobot() {
    get(common.router.stopRobotUrl).then(function (resp){
        if (resp.status == 200){
            common.addSuccess()
        }else {
            resp.json().then(data => {
                common.addFail(JSON.stringify(data))
            })
        }
    })
}

//增加設定檔
function addSetting() {
    var envReq = document.getElementById(`envSelect`).value;
    var gameID = Number(document.getElementById(`gameSelect`).value);
    var roomType = document.getElementById(`roomSelect`).value;
    var robotNum = Number(document.getElementById(`robotnum`).value);

    if (envReq === "" || envReq === "Please Select") {
        common.addFail("please Select env")
        return
    }
    if (isNaN(gameID)){
        common.addFail("please Select gameID")
        return
    }
    if (roomType === "" || roomType === "Please Select"){
        common.addFail("please Select roomType")
        return
    }
    if (isNaN(robotNum)){
        common.addFail("please Select robotNum")
        return
    }

    var cfg = document.getElementById(`config`).value;
    // console.log("我的CFG:"+cfg)
    var gameName = document.getElementById(gameID).innerText;

    var setting = {
        id: count,
        gameid: gameID,
        roomtype: roomType,
        env: envReq,
        robotcount: robotNum,
        strategy: cfg,
    }
    var showSetting = {
        "id": count,
        "遊戲id": gameID,
        "遊戲名稱": gameName,
        "遊戲房間":   roomType,
        "機器人數量": robotNum,
    }
    settings.set(count, setting)
    showSettings.set(count, showSetting)
    count++
    refreshParam()
}

//刪除設定擋
function deleteSetting() {
    var deleteid = Number(document.getElementById(`deleteSetting`).value);
    if (isNaN(deleteid)) {
        common.addFail()
    } else {
        // console.log(deleteid)
        var bool = settings.delete(deleteid)
        var bool2 = showSettings.delete(deleteid)
        if (bool && bool2) {
            document.getElementById(`deleteSetting`).disabled = true;
            document.getElementById(`config`).value = "";
            common.addSuccess()
            refreshParam()
        }else {
            common.addFail()
        }
    }
}

//刪除所有設定
function deleteAllSetting() {
    settings = new Map();
    showSettings = new Map();
    refreshParam()
    common.addSuccess()
}

//點擊遊戲列表
function clickStressGame(select) {
    var game = gameList[select.value];
    var roomContainer = document.getElementById("roomSelect")
    while (roomContainer.lastElementChild){
        roomContainer.removeChild(roomContainer.lastElementChild)
    }

    Object.values(game["RoomType"]).forEach((value,index) => {
        var room = document.createElement("option");
        room.innerText = value
        room.value = value
        roomContainer.appendChild(room);
        if (index === 0) {
            reloadLeftSideConfig(game, game["RoomType"][0],false)
        }
    })
}

function clickStressRoom(select) {
    var roomName = select.value
    var roomID = document.getElementById("gameSelect").value
    reloadLeftSideConfig(gameList[roomID], roomName,false)
}


function  connectWS(data) {

    sock = new socket({
        //網址（端口是我下面的服務器的端口）
        'url': common.router.stressWSUrl ,
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



//重整設定
function refreshParam() {
    var param = document.getElementById("param")

    //全部先清除
    while (param.lastElementChild) {
        param.removeChild(param.lastElementChild)
    }
    //重新塞
    var rr = document.createElement("div");
    rr.className = "row"
    showSettings.forEach((value, key) => {
        var root =  document.createElement("div");
        var setting =  document.createElement("button");
        setting.innerText = JSON.stringify(value, null, 2).replace('{\n','').replace('\n}','');
        setting.className = "areasize"
        setting.onclick = function () {
            reloadLeftSideConfig(gameList[value["遊戲id"]],value["遊戲房間"],true, value.id)
        }
        root.className = "col-sm"
        root.appendChild(setting)
        rr.appendChild(root)
    })
    param.appendChild(rr)


}

function mapToObj(inputMap) {
    let obj = [];

    inputMap.forEach(function(value, key){
        obj.push(value)
    });

    return obj;
}

//重新reload config
function reloadLeftSideConfig(game, roomCfgID, isButton, deleteId) {
    // console.log("game: "+game+ " roomCfig: "+ roomCfgID)
    //先把delete ture
    var deleteBtn = document.getElementById(`deleteSetting`);
    deleteBtn.disabled= true;

    var cfgAreaText = document.getElementById(`config`);
    var cfg = new Map(Object.entries(game["RoomStrategy"]))
    var strategy = cfg.get(roomCfgID)
    cfgAreaText.value = JSON.stringify(strategy)

    if (isButton) {
        var st =settings.get(deleteId)
        strategy = st.strategy
        cfgAreaText.value = strategy
        deleteBtn.disabled = false;
        deleteBtn.value = deleteId;
    }
}


export {
    stopRobot,
    addSetting,
    deleteSetting,
    startStress,
    deleteAllSetting,
    clickStressGame,
    clickStressRoom
}