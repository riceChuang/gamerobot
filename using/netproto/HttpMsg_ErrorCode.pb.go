// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: HttpMsg_ErrorCode.proto

package netproto

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Http_ErrorCodeID int32

const (
	Http_ErrorCodeID_API_SUCCESS                      Http_ErrorCodeID = 0    //成功
	Http_ErrorCodeID_API_TOKEN_MISS_ERROR             Http_ErrorCodeID = 1    //TOKEN 丢失
	Http_ErrorCodeID_API_AGENT_MISS_ERROR             Http_ErrorCodeID = 2    //渠道不存在
	Http_ErrorCodeID_API_TIMESTAMP_ERROR              Http_ErrorCodeID = 3    //验证时间超时
	Http_ErrorCodeID_API_KEY_ERROR                    Http_ErrorCodeID = 4    //验证错误
	Http_ErrorCodeID_API_AREA_WHITEIP_ERROR           Http_ErrorCodeID = 5    //渠道白名單錯誤（請聯繫客服添加服務器白名單)
	Http_ErrorCodeID_API_KEY_MISS_ERROR               Http_ErrorCodeID = 6    //驗證字段丟失
	Http_ErrorCodeID_API_ROUTER_ERROR                 Http_ErrorCodeID = 8    //不存在的請求
	Http_ErrorCodeID_API_AREA_KEY_ERROR               Http_ErrorCodeID = 15   //渠道驗證錯誤
	Http_ErrorCodeID_API_BET_DATA_MISS_ERROR          Http_ErrorCodeID = 16   //數據不存在（當前沒有注單）
	Http_ErrorCodeID_API_ACCOUNT_FORBIDEN_ERROR       Http_ErrorCodeID = 20   //帳號禁用
	Http_ErrorCodeID_API_AES_ERROR                    Http_ErrorCodeID = 22   //AES解密失敗
	Http_ErrorCodeID_API_AREA_QUERY_OUT_OF_RANGE      Http_ErrorCodeID = 24   //渠道拉取數據超過時間範圍
	Http_ErrorCodeID_API_BET_ORDER_MISS               Http_ErrorCodeID = 26   //訂單號不存在
	Http_ErrorCodeID_API_DB_ERROR                     Http_ErrorCodeID = 27   //數據庫異常
	Http_ErrorCodeID_API_IP_FORBIDEN                  Http_ErrorCodeID = 28   //IP禁用
	Http_ErrorCodeID_API_BET_ORDER_REGULER_ERROR      Http_ErrorCodeID = 29   //訂單號與訂單規則不符
	Http_ErrorCodeID_API_USER_ONLINE_INFO_ERROR       Http_ErrorCodeID = 30   //獲取玩家在線狀態失敗
	Http_ErrorCodeID_API_UPDATE_SCORE_OUT_OF_RANGE    Http_ErrorCodeID = 31   //更新的玩家分數小於或者等於0
	Http_ErrorCodeID_API_UPDATE_USER_INFO_ERROR       Http_ErrorCodeID = 32   //更新玩家信息失敗
	Http_ErrorCodeID_API_UPDATE_USER_SCORE_ERROR      Http_ErrorCodeID = 33   //更新玩家金幣失敗
	Http_ErrorCodeID_API_REPEAT_BET_ORDER             Http_ErrorCodeID = 34   //訂單重複
	Http_ErrorCodeID_API_USER_INFO_ERROR              Http_ErrorCodeID = 35   //獲取玩家訊息失敗
	Http_ErrorCodeID_API_KINDID_ERROR                 Http_ErrorCodeID = 36   //KindID不存在
	Http_ErrorCodeID_API_LOGIN_NO_WITHDRAWAL          Http_ErrorCodeID = 37   //登入瞬間禁止下分, 導致下分失敗
	Http_ErrorCodeID_API_WITHDRAWAL_NO_MONEY          Http_ErrorCodeID = 38   //餘額不足導致下分失敗
	Http_ErrorCodeID_API_CONCURRENCY_SYN_MONEY_ERROR  Http_ErrorCodeID = 39   //禁止同一帳號登入帶分, 上分, 下分並發請求
	Http_ErrorCodeID_API_SYN_SCORE_LIMIT              Http_ErrorCodeID = 40   //單次上下分數量不能超過一千萬
	Http_ErrorCodeID_API_REPORT_TIME_RANGE_ERROR      Http_ErrorCodeID = 41   //拉取對局匯總時間範圍有誤
	Http_ErrorCodeID_API_AGENT_FORBIDEN               Http_ErrorCodeID = 42   //代理被禁用
	Http_ErrorCodeID_API_SYNC_TIMES_LIMIT             Http_ErrorCodeID = 43   //拉單過於頻繁(兩次拉單時間間隔必須大於5秒)
	Http_ErrorCodeID_API_ORDER_PROCESSING             Http_ErrorCodeID = 44   //訂單正在處理中
	Http_ErrorCodeID_API_PARAM_ERROR                  Http_ErrorCodeID = 45   //參數錯誤
	Http_ErrorCodeID_API_REGISTER_MEMBER_SYSTEM_ERROR Http_ErrorCodeID = 1001 //註冊會員帳號系統異常
	Http_ErrorCodeID_API_AGENT_MONEY_ERROR            Http_ErrorCodeID = 1002 //代理商金額不足
)

// Enum value maps for Http_ErrorCodeID.
var (
	Http_ErrorCodeID_name = map[int32]string{
		0:    "API_SUCCESS",
		1:    "API_TOKEN_MISS_ERROR",
		2:    "API_AGENT_MISS_ERROR",
		3:    "API_TIMESTAMP_ERROR",
		4:    "API_KEY_ERROR",
		5:    "API_AREA_WHITEIP_ERROR",
		6:    "API_KEY_MISS_ERROR",
		8:    "API_ROUTER_ERROR",
		15:   "API_AREA_KEY_ERROR",
		16:   "API_BET_DATA_MISS_ERROR",
		20:   "API_ACCOUNT_FORBIDEN_ERROR",
		22:   "API_AES_ERROR",
		24:   "API_AREA_QUERY_OUT_OF_RANGE",
		26:   "API_BET_ORDER_MISS",
		27:   "API_DB_ERROR",
		28:   "API_IP_FORBIDEN",
		29:   "API_BET_ORDER_REGULER_ERROR",
		30:   "API_USER_ONLINE_INFO_ERROR",
		31:   "API_UPDATE_SCORE_OUT_OF_RANGE",
		32:   "API_UPDATE_USER_INFO_ERROR",
		33:   "API_UPDATE_USER_SCORE_ERROR",
		34:   "API_REPEAT_BET_ORDER",
		35:   "API_USER_INFO_ERROR",
		36:   "API_KINDID_ERROR",
		37:   "API_LOGIN_NO_WITHDRAWAL",
		38:   "API_WITHDRAWAL_NO_MONEY",
		39:   "API_CONCURRENCY_SYN_MONEY_ERROR",
		40:   "API_SYN_SCORE_LIMIT",
		41:   "API_REPORT_TIME_RANGE_ERROR",
		42:   "API_AGENT_FORBIDEN",
		43:   "API_SYNC_TIMES_LIMIT",
		44:   "API_ORDER_PROCESSING",
		45:   "API_PARAM_ERROR",
		1001: "API_REGISTER_MEMBER_SYSTEM_ERROR",
		1002: "API_AGENT_MONEY_ERROR",
	}
	Http_ErrorCodeID_value = map[string]int32{
		"API_SUCCESS":                      0,
		"API_TOKEN_MISS_ERROR":             1,
		"API_AGENT_MISS_ERROR":             2,
		"API_TIMESTAMP_ERROR":              3,
		"API_KEY_ERROR":                    4,
		"API_AREA_WHITEIP_ERROR":           5,
		"API_KEY_MISS_ERROR":               6,
		"API_ROUTER_ERROR":                 8,
		"API_AREA_KEY_ERROR":               15,
		"API_BET_DATA_MISS_ERROR":          16,
		"API_ACCOUNT_FORBIDEN_ERROR":       20,
		"API_AES_ERROR":                    22,
		"API_AREA_QUERY_OUT_OF_RANGE":      24,
		"API_BET_ORDER_MISS":               26,
		"API_DB_ERROR":                     27,
		"API_IP_FORBIDEN":                  28,
		"API_BET_ORDER_REGULER_ERROR":      29,
		"API_USER_ONLINE_INFO_ERROR":       30,
		"API_UPDATE_SCORE_OUT_OF_RANGE":    31,
		"API_UPDATE_USER_INFO_ERROR":       32,
		"API_UPDATE_USER_SCORE_ERROR":      33,
		"API_REPEAT_BET_ORDER":             34,
		"API_USER_INFO_ERROR":              35,
		"API_KINDID_ERROR":                 36,
		"API_LOGIN_NO_WITHDRAWAL":          37,
		"API_WITHDRAWAL_NO_MONEY":          38,
		"API_CONCURRENCY_SYN_MONEY_ERROR":  39,
		"API_SYN_SCORE_LIMIT":              40,
		"API_REPORT_TIME_RANGE_ERROR":      41,
		"API_AGENT_FORBIDEN":               42,
		"API_SYNC_TIMES_LIMIT":             43,
		"API_ORDER_PROCESSING":             44,
		"API_PARAM_ERROR":                  45,
		"API_REGISTER_MEMBER_SYSTEM_ERROR": 1001,
		"API_AGENT_MONEY_ERROR":            1002,
	}
)

func (x Http_ErrorCodeID) Enum() *Http_ErrorCodeID {
	p := new(Http_ErrorCodeID)
	*p = x
	return p
}

func (x Http_ErrorCodeID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Http_ErrorCodeID) Descriptor() protoreflect.EnumDescriptor {
	return file_HttpMsg_ErrorCode_proto_enumTypes[0].Descriptor()
}

func (Http_ErrorCodeID) Type() protoreflect.EnumType {
	return &file_HttpMsg_ErrorCode_proto_enumTypes[0]
}

func (x Http_ErrorCodeID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Http_ErrorCodeID) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Http_ErrorCodeID(num)
	return nil
}

// Deprecated: Use Http_ErrorCodeID.Descriptor instead.
func (Http_ErrorCodeID) EnumDescriptor() ([]byte, []int) {
	return file_HttpMsg_ErrorCode_proto_rawDescGZIP(), []int{0}
}

var File_HttpMsg_ErrorCode_proto protoreflect.FileDescriptor

var file_HttpMsg_ErrorCode_proto_rawDesc = []byte{
	0x0a, 0x17, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x73, 0x67, 0x5f, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x43,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x65, 0x74, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2a, 0xc0, 0x07, 0x0a, 0x10, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x43, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x50, 0x49, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x50, 0x49,
	0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x4d, 0x49, 0x53, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x50, 0x49, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54,
	0x5f, 0x4d, 0x49, 0x53, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x17, 0x0a,
	0x13, 0x41, 0x50, 0x49, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x54, 0x41, 0x4d, 0x50, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x50, 0x49, 0x5f, 0x4b, 0x45,
	0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x04, 0x12, 0x1a, 0x0a, 0x16, 0x41, 0x50, 0x49,
	0x5f, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x57, 0x48, 0x49, 0x54, 0x45, 0x49, 0x50, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0x05, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x50, 0x49, 0x5f, 0x4b, 0x45, 0x59,
	0x5f, 0x4d, 0x49, 0x53, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x14, 0x0a,
	0x10, 0x41, 0x50, 0x49, 0x5f, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x08, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x50, 0x49, 0x5f, 0x41, 0x52, 0x45, 0x41, 0x5f,
	0x4b, 0x45, 0x59, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x0f, 0x12, 0x1b, 0x0a, 0x17, 0x41,
	0x50, 0x49, 0x5f, 0x42, 0x45, 0x54, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x4d, 0x49, 0x53, 0x53,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x10, 0x12, 0x1e, 0x0a, 0x1a, 0x41, 0x50, 0x49, 0x5f,
	0x41, 0x43, 0x43, 0x4f, 0x55, 0x4e, 0x54, 0x5f, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x45, 0x4e,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x14, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x50, 0x49, 0x5f,
	0x41, 0x45, 0x53, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x16, 0x12, 0x1f, 0x0a, 0x1b, 0x41,
	0x50, 0x49, 0x5f, 0x41, 0x52, 0x45, 0x41, 0x5f, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x4f, 0x55,
	0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x18, 0x12, 0x16, 0x0a, 0x12,
	0x41, 0x50, 0x49, 0x5f, 0x42, 0x45, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x4d, 0x49,
	0x53, 0x53, 0x10, 0x1a, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x50, 0x49, 0x5f, 0x44, 0x42, 0x5f, 0x45,
	0x52, 0x52, 0x4f, 0x52, 0x10, 0x1b, 0x12, 0x13, 0x0a, 0x0f, 0x41, 0x50, 0x49, 0x5f, 0x49, 0x50,
	0x5f, 0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x45, 0x4e, 0x10, 0x1c, 0x12, 0x1f, 0x0a, 0x1b, 0x41,
	0x50, 0x49, 0x5f, 0x42, 0x45, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x5f, 0x52, 0x45, 0x47,
	0x55, 0x4c, 0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x1d, 0x12, 0x1e, 0x0a, 0x1a,
	0x41, 0x50, 0x49, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x4f, 0x4e, 0x4c, 0x49, 0x4e, 0x45, 0x5f,
	0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x1e, 0x12, 0x21, 0x0a, 0x1d,
	0x41, 0x50, 0x49, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45,
	0x5f, 0x4f, 0x55, 0x54, 0x5f, 0x4f, 0x46, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x1f, 0x12,
	0x1e, 0x0a, 0x1a, 0x41, 0x50, 0x49, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x20, 0x12,
	0x1f, 0x0a, 0x1b, 0x41, 0x50, 0x49, 0x5f, 0x55, 0x50, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x53,
	0x45, 0x52, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x21,
	0x12, 0x18, 0x0a, 0x14, 0x41, 0x50, 0x49, 0x5f, 0x52, 0x45, 0x50, 0x45, 0x41, 0x54, 0x5f, 0x42,
	0x45, 0x54, 0x5f, 0x4f, 0x52, 0x44, 0x45, 0x52, 0x10, 0x22, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x50,
	0x49, 0x5f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x49, 0x4e, 0x46, 0x4f, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x23, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x50, 0x49, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x49,
	0x44, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x24, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x50, 0x49,
	0x5f, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x4e, 0x4f, 0x5f, 0x57, 0x49, 0x54, 0x48, 0x44, 0x52,
	0x41, 0x57, 0x41, 0x4c, 0x10, 0x25, 0x12, 0x1b, 0x0a, 0x17, 0x41, 0x50, 0x49, 0x5f, 0x57, 0x49,
	0x54, 0x48, 0x44, 0x52, 0x41, 0x57, 0x41, 0x4c, 0x5f, 0x4e, 0x4f, 0x5f, 0x4d, 0x4f, 0x4e, 0x45,
	0x59, 0x10, 0x26, 0x12, 0x23, 0x0a, 0x1f, 0x41, 0x50, 0x49, 0x5f, 0x43, 0x4f, 0x4e, 0x43, 0x55,
	0x52, 0x52, 0x45, 0x4e, 0x43, 0x59, 0x5f, 0x53, 0x59, 0x4e, 0x5f, 0x4d, 0x4f, 0x4e, 0x45, 0x59,
	0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x27, 0x12, 0x17, 0x0a, 0x13, 0x41, 0x50, 0x49, 0x5f,
	0x53, 0x59, 0x4e, 0x5f, 0x53, 0x43, 0x4f, 0x52, 0x45, 0x5f, 0x4c, 0x49, 0x4d, 0x49, 0x54, 0x10,
	0x28, 0x12, 0x1f, 0x0a, 0x1b, 0x41, 0x50, 0x49, 0x5f, 0x52, 0x45, 0x50, 0x4f, 0x52, 0x54, 0x5f,
	0x54, 0x49, 0x4d, 0x45, 0x5f, 0x52, 0x41, 0x4e, 0x47, 0x45, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x29, 0x12, 0x16, 0x0a, 0x12, 0x41, 0x50, 0x49, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f,
	0x46, 0x4f, 0x52, 0x42, 0x49, 0x44, 0x45, 0x4e, 0x10, 0x2a, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x50,
	0x49, 0x5f, 0x53, 0x59, 0x4e, 0x43, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x53, 0x5f, 0x4c, 0x49, 0x4d,
	0x49, 0x54, 0x10, 0x2b, 0x12, 0x18, 0x0a, 0x14, 0x41, 0x50, 0x49, 0x5f, 0x4f, 0x52, 0x44, 0x45,
	0x52, 0x5f, 0x50, 0x52, 0x4f, 0x43, 0x45, 0x53, 0x53, 0x49, 0x4e, 0x47, 0x10, 0x2c, 0x12, 0x13,
	0x0a, 0x0f, 0x41, 0x50, 0x49, 0x5f, 0x50, 0x41, 0x52, 0x41, 0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x2d, 0x12, 0x25, 0x0a, 0x20, 0x41, 0x50, 0x49, 0x5f, 0x52, 0x45, 0x47, 0x49, 0x53,
	0x54, 0x45, 0x52, 0x5f, 0x4d, 0x45, 0x4d, 0x42, 0x45, 0x52, 0x5f, 0x53, 0x59, 0x53, 0x54, 0x45,
	0x4d, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0xe9, 0x07, 0x12, 0x1a, 0x0a, 0x15, 0x41, 0x50,
	0x49, 0x5f, 0x41, 0x47, 0x45, 0x4e, 0x54, 0x5f, 0x4d, 0x4f, 0x4e, 0x45, 0x59, 0x5f, 0x45, 0x52,
	0x52, 0x4f, 0x52, 0x10, 0xea, 0x07, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6e, 0x65, 0x74, 0x70,
	0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_HttpMsg_ErrorCode_proto_rawDescOnce sync.Once
	file_HttpMsg_ErrorCode_proto_rawDescData = file_HttpMsg_ErrorCode_proto_rawDesc
)

func file_HttpMsg_ErrorCode_proto_rawDescGZIP() []byte {
	file_HttpMsg_ErrorCode_proto_rawDescOnce.Do(func() {
		file_HttpMsg_ErrorCode_proto_rawDescData = protoimpl.X.CompressGZIP(file_HttpMsg_ErrorCode_proto_rawDescData)
	})
	return file_HttpMsg_ErrorCode_proto_rawDescData
}

var file_HttpMsg_ErrorCode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_HttpMsg_ErrorCode_proto_goTypes = []interface{}{
	(Http_ErrorCodeID)(0), // 0: netproto.Http_ErrorCodeID
}
var file_HttpMsg_ErrorCode_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_HttpMsg_ErrorCode_proto_init() }
func file_HttpMsg_ErrorCode_proto_init() {
	if File_HttpMsg_ErrorCode_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HttpMsg_ErrorCode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HttpMsg_ErrorCode_proto_goTypes,
		DependencyIndexes: file_HttpMsg_ErrorCode_proto_depIdxs,
		EnumInfos:         file_HttpMsg_ErrorCode_proto_enumTypes,
	}.Build()
	File_HttpMsg_ErrorCode_proto = out.File
	file_HttpMsg_ErrorCode_proto_rawDesc = nil
	file_HttpMsg_ErrorCode_proto_goTypes = nil
	file_HttpMsg_ErrorCode_proto_depIdxs = nil
}
