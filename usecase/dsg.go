package usecase

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"github.com/riceChuang/gamerobot/model"
	"github.com/riceChuang/gamerobot/util/crypt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

var dsgApiService *DSGApiService
var dsgApiServiceOnce = &sync.Once{}

type DSGApiService struct{}

func NewDSGService() DsgAPIBase {
	dsgApiServiceOnce.Do(func() {
		dsgApiService = &DSGApiService{}
	})
	return dsgApiService
}

// test usage DEV
var (
	loginAction = "0"
	storeAction = "2"
	password    = "123qwe"
	ip          = "127.0.0.1"
	AESKey      = "xyz123!QAZ2wsxGE"
)

// Login to get token
func (ds *DSGApiService) Login(loginReq *model.DSGLoginReq) (loginResp *model.DSGLoginResp, err error) {
	timestamp := time.Now().Unix() // s
	param := ds.getLoginParam(loginReq.Account, password, loginReq.GameID)
	postURL := ds.getPostURL(loginReq.LoginDomain, 8200, loginReq.AgentID, param, timestamp)

	resp, err := http.Post(postURL,
		"application/x-www-form-urlencoded",
		strings.NewReader("test"),
	)
	if err != nil {
		return nil, fmt.Errorf("err: %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("err: %s", err.Error())
	}

	var jsonObj map[string]interface{}
	if err := json.Unmarshal(body, &jsonObj); err != nil {
		return nil, fmt.Errorf("err: %s", err.Error())
	}

	if s, ok := jsonObj["s"]; ok {
		if d, ok := jsonObj["d"]; s.(float64) == 100 && ok {
			fmt.Printf("Get Login Res %+v\n", jsonObj)
			if hallUrl, ok := d.(map[string]interface{})["url"]; ok {

				loginResp = &model.DSGLoginResp{
					URL: hallUrl.(string),
				}

				tokens := strings.Split(hallUrl.(string), "token=")
				if len(tokens) >= 2 {
					loginResp.Token = tokens[1]
				}

				acoounts := strings.Split(hallUrl.(string), "account=")
				if len(acoounts) >= 2 {
					loginResp.UserName = acoounts[1]
				}
				return
			} else {
				return nil, fmt.Errorf("LoginFail: res data %s", string(body))
			}
		}
	}
	return nil, fmt.Errorf("LoginFail: res data %s", string(body))
}

// Login to get token
func (ds *DSGApiService) StoreMoney(loginDomain string, agentID int, account string, money int32) (int32, error) {
	timestamp := time.Now().Unix() // s
	score := money * 100
	param := ds.getStoreParam(account, score, timestamp)
	postURL := ds.getPostURL(loginDomain, 8200, agentID, param, timestamp)

	resp, err := http.Post(postURL,
		"application/x-www-form-urlencoded",
		strings.NewReader("test"),
	)
	if err != nil {
		return 0, fmt.Errorf("err: %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("err: %s", err.Error())
	}

	var jsonObj map[string]interface{}
	if err := json.Unmarshal(body, &jsonObj); err != nil {
		return 0, fmt.Errorf("err: %s", err.Error())
	}

	if s, ok := jsonObj["s"]; ok {
		if d, ok := jsonObj["d"]; s.(float64) == 102 && ok {
			fmt.Printf("Get Store Res %+v\n", jsonObj)
			if money, ok := d.(map[string]interface{})["money"]; ok {
				return int32(money.(float64)), nil
			}
			return 0, fmt.Errorf("StoreFail: res data %s", string(body))
		}
	}
	return 0, fmt.Errorf("StoreFail: res data %s", string(body))
}

func (ds *DSGApiService) getStoreParam(account string, money int32, timestamp int64) string {
	orderNo := fmt.Sprintf("%s%d%s", storeAction, timestamp, account)
	param := fmt.Sprintf("account=%s&money=%d&orderNo=%s&s=%s", account, money, orderNo, storeAction)
	ciphertext := crypt.AesEncrypt(param, AESKey)
	return ciphertext
}

func (ds *DSGApiService) getLoginParam(account string, password string, gameID int32) string {
	param := fmt.Sprintf("account=%s&gameId=%d&ip=%s&password=%s&s=%s", account, gameID, ip, password, loginAction)
	ciphertext := crypt.AesEncrypt(param, AESKey)
	return ciphertext
}

func (ds *DSGApiService) getPostURL(domain string, port int, agentID int, param string, timestamp int64) string {
	params := url.Values{}
	params.Add("param", param)
	escapeParam := params.Encode()
	md5Key := ds.getMD5Key(agentID, timestamp)
	URL := fmt.Sprintf("http://%s:%d/channelHandle?agent=%d&timestamp=%d&%s&key=%s", domain, port, agentID, timestamp, escapeParam, md5Key)
	return URL
}

func (ds *DSGApiService) getMD5Key(agentID int, timestamp int64) string {
	h := md5.New()
	sKey := "xyz123"
	s := fmt.Sprintf("%d%d%s", agentID, timestamp, sKey)
	io.WriteString(h, s)
	key := fmt.Sprintf("%x", h.Sum((nil)))
	return key
}
