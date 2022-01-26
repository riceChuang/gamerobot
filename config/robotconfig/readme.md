## 設置config 介紹

- "name": "config1"

用於機器人帳號的index 目前沒有使用 但還是先設置

- "rooms":

設定多間房間的config, 請由體驗房依序往下新增

"roomIndex": 1   // 房間index

"robotNumber": 50  // 此房間最多可使用的機器人數量, 若壓測時數入超過則會錯誤

"minMoney": 100 // 設定至少攜帶入得最小金額

！！！！需注意！！！！！

1. 此金額為元 ex. 100 = 100元人民幣
2. 假如遊戲的最低準入金額為 10元人民幣 請將 minMoney設定超過此金額多一些 ex. 50
3. 補充2點, 因為若低於準入金額時遊戲會直接踢除玩家 服務就無法偵測做自動加值

"storedMoney": 1000 //當身上金額低於minMoney時候單次存進去的錢 單位 元

"robotName": "bydh_robot1"  // 機器人的名字前面的index
e.x. bydh_robot1_0000 ~ bydh_robot1_0050