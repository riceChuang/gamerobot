class socket{
    /* websocket實例 */
    ws = null

    /*'#'爲私有屬性，外部不可調用 */

    /* 狀態 */
    //連接狀態
    #alive = false
    //把類的參數傳入這裏，方便調用
    #params = null

    /* 計時器 */
    //重連計時器
    #reconnect_timer = null
    //心跳計時器
    #heart_timer = null

    /* 參數 */
    //心跳時間 50秒一次
    heartBeat = 50000
    //心跳信息：默認爲‘ping’隨便改，看後臺
    heartMsg = 'ping'
    //是否自動重連
    reconnect = false
    //重連間隔時間
    reconnectTime = 5000
    //重連次數
    reconnectTimes = 10
    //帶入參數
    urlParam = ""
    constructor(params) {
        this.#params = params
        this.init()
    }

    /* 初始化 */
    init() {
        //重中之重，不然重連的時候會越來越快
        clearInterval(this.#reconnect_timer)
        clearInterval(this.#heart_timer)

        //取出所有參數
        let params = this.#params
        //設置連接路徑
        let {url, port} = params
        let global_params = ['heartBeat','heartMsg','reconnect','reconnectTime','reconnectTimes','urlParam']

        //定義全局變量
        Object.keys(params).forEach(key=>{
            if (global_params.indexOf(key) !== -1) {
                this[key] = params[key]
            }
        })

        let ws_url = port ? url + ':' + port + '?' + this.urlParam: url + '?' + this.urlParam
        // let ws_url = port ? url + ':' + port : url
        this.ws = null
        this.ws = new WebSocket(ws_url)

        //默認綁定事件
        this.ws.onopen = () => {
            //設置狀態爲開啓
            this.#alive = true
            clearInterval(this.#reconnect_timer)
            //連接後進入心跳狀態
            this.onheartbeat()
        }

        this.ws.onclose = () => {
            //設置狀態爲斷開
            this.#alive = false

            clearInterval(this.#heart_timer)

            //自動重連開啓  +  不在重連狀態下
            if (true == this.reconnect) {
                /* 斷開後立刻重連 */
                this.onreconnect()
            }
        }
    }


    /*
     *
     * 新增‘心跳事件’和‘重連事件’
     *
     */

    /* 心跳事件 */
    onheartbeat(func) {
        //在連接狀態下
        if (true == this.#alive) {
            /* 心跳計時器 */
            this.#heart_timer = setInterval(()=>{
                //發送心跳信息
                this.send(this.heartMsg)
                func ? func(this) :false

            },this.heartBeat)
        }
    }

    /* 重連事件 */
    onreconnect(func) {
        /* 重連間隔計時器 */
        this.#reconnect_timer = setInterval(()=>{
            //限制重連次數
            if (this.reconnectTimes <= 0) {
                //關閉定時器
                // this.#isReconnect = false
                clearInterval(this.#reconnect_timer)
                //跳出函數之間的循環
                return;
            }else{
                //重連一次-1
                this.reconnectTimes--
            }
            //進入初始狀態
            this.init()
            func ? func(this) : false
        },this.reconnectTime)
    }

    /*
     *
     * 對原生方法和事件進行封裝
     *
     */

    // 發送消息
    send(text) {
        if (true == this.#alive) {
            text = typeof text == 'string' ?  text : JSON.stringify(text)
            this.ws.send(text)
        }
    }

    // 斷開連接
    close() {
        if (true == this.#alive) {
            this.ws.close()
        }
    }

    //接受消息
    onmessage(func,all = false) {

        this.ws.onmessage = data => func(!all ? data.data : data)
    }

    //websocket連接成功事件
    onopen(func) {
        this.ws.onopen = event => {

            this.#alive = true
            func ? func(event) : false

        }
    }
    //websocket關閉事件
    onclose(func) {
        this.ws.onclose = event => {

            //設置狀態爲斷開
            this.#alive = false

            clearInterval(this.#heart_timer)

            //自動重連開啓  +  不在重連狀態下
            if (true == this.reconnect) {
                /* 斷開後立刻重連 */
                this.onreconnect()
            }
            func ? func(event) : false
        }
    }
    //websocket錯誤事件
    onerror(func) {
        this.ws.onerror = event => {
            func ? func(event) : false
        }
    }
}

export {
    socket
}