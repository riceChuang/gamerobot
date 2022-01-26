const domainName = "127.0.0.1:8888"

const router = {
    loginUrl: "login",
    msgUrl: "msg",
    stopRobotUrl: "stopstress",
    startRobotUrl: "startstress",
    gameListUrl: "gamelist",
    stressWSUrl : "ws://127.0.0.1:8888/stressws",
    wsUrl: "ws://127.0.0.1:8888/ws",
}

router.stressWSUrl = domain + "/stressws"
router.wsUrl = domain + "/ws"

console.log(router.stressWSUrl)
console.log(router.wsUrl)

window.post = function(url, data) {
    return fetch(url, {method: "POST", headers: {'Content-Type': 'application/json'}, body: JSON.stringify(data)});
}

window.get = function(url, data) {
    return fetch(url, {method: "GET", headers: {'Content-Type': 'application/json'}, body: JSON.stringify(data)});
}

function addFail(data) {
    var bd = document.body
    var root =  document.createElement("div")
    var pp = document.createElement("p")
    root.role = "alart"
    root.className = "fail-msg message"
    root.style = "top: 20px; z-index: 2014;"
    pp.innerText = "操作失敗失敗失敗失敗失敗失敗" + data
    pp.style="text-align:center"
    pp.className = "msg-conetnt"
    pp.style="text-align:center"
    root.appendChild(pp)
    bd.appendChild(root)
    hide(root)
}

function addSuccess() {
    var bd = document.body
    var root =  document.createElement("div")
    var pp = document.createElement("p")
    root.role = "alart"
    root.className = "success-msg message"
    root.style = "top: 20px; z-index: 2014;"
    pp.innerText = "操作成功成功成功成功成功成功"
    pp.className = "msg-conetnt"
    pp.style="text-align:center"
    root.appendChild(pp)
    bd.appendChild(root)
    hide(root)
}

function hide(value) {
    var count = (function() {
        var timer;
        var i = 0;
        function change(tar) {
            i++;
            var num = 1-i/45;
            value.style.opacity=num;
            if (num == 0) {
                value.remove()
            }
            if (i === tar) {
                clearTimeout(timer);
                return false;
            }
            timer = setTimeout(function() {
                change(tar)
            }, 45)
        }
        return change;
    })()
    count(45)
}




export {
    router,
    addFail,
    addSuccess
}
