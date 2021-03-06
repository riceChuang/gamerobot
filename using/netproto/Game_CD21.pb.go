// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: Game_CD21.proto

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

type CD21_GameMessageClassID int32

const (
	CD21_GameMessageClassID_CD21SettingID      CD21_GameMessageClassID = 10001  //设置相关  S2C
	CD21_GameMessageClassID_CD21StateID        CD21_GameMessageClassID = 10002  //阶段ID S2C
	CD21_GameMessageClassID_CD21BetID          CD21_GameMessageClassID = 10003  //下注ID C2S
	CD21_GameMessageClassID_CD21BroadCastBetID CD21_GameMessageClassID = 10004  //广播下注 S2C
	CD21_GameMessageClassID_CD21GameSceneID    CD21_GameMessageClassID = 10005  //游戏场景 断线重连 C2S (CD21_UserInfoArr)
	CD21_GameMessageClassID_CD21BetInfoID      CD21_GameMessageClassID = 100006 //用户断线重连已下注的消息 S2C
	CD21_GameMessageClassID_CD21BetRetID       CD21_GameMessageClassID = 100007 //下注结果 S2C
	//	CD21CutpaiID=10008;				//切牌 S2C
	//	CD21BroadCastCutpaiRetID=10009;	//切哪張牌 C2S
	//	CD21BroadCastCutpaiID=10010;	//切牌廣播 S2C
	CD21_GameMessageClassID_CD21BroadCastFapaiID CD21_GameMessageClassID = 10011 //發牌 S2C
	CD21_GameMessageClassID_CD21UserOptID        CD21_GameMessageClassID = 10012 //玩家可操作行為 S2C
	//	CD21UserOptRetID=10013;			//玩家操作行為 C2S
	CD21_GameMessageClassID_CD21GameResultWinID CD21_GameMessageClassID = 10014 //游戏结果赢钱 S2C
	//	CD21UserSurrender=10017;			//玩家投降 C2S
	CD21_GameMessageClassID_CD21UserInsurence    CD21_GameMessageClassID = 10018 //玩家保險 C2S
	CD21_GameMessageClassID_CD21UserHit          CD21_GameMessageClassID = 10019 //玩家要牌 C2S
	CD21_GameMessageClassID_CD21UserStand        CD21_GameMessageClassID = 10020 //玩家停牌 C2S
	CD21_GameMessageClassID_CD21UserSplit        CD21_GameMessageClassID = 10021 //玩家分牌 C2S
	CD21_GameMessageClassID_CD21UserDouble       CD21_GameMessageClassID = 10022 //玩家加倍 C2S
	CD21_GameMessageClassID_CD21UserCancel       CD21_GameMessageClassID = 10023 //玩家不投降不保險 C2S
	CD21_GameMessageClassID_CD21CurrOptID        CD21_GameMessageClassID = 10024 //当前操作的椅子号 S2C
	CD21_GameMessageClassID_CD21StandBroadCastID CD21_GameMessageClassID = 10025 //停牌广播 S2C
	//	CD21UserSplitHit=10026;					//玩家要牌 C2S
	//	CD21UserSplitStand=10027;				//玩家停牌 C2S
	//	CD21UserSplitDouble=10028;				//玩家加倍 C2S
	CD21_GameMessageClassID_CD21SplitCardBroacast    CD21_GameMessageClassID = 10029 //玩家分牌 S2C (CD21_SplitCard)
	CD21_GameMessageClassID_CD21BroadCastYaopaiID    CD21_GameMessageClassID = 10030 //要牌廣播 S2C (CD21_Yaopai)
	CD21_GameMessageClassID_CD21BroadCastFanPaiID    CD21_GameMessageClassID = 10031 //莊家翻牌廣播 S2C (CD21_FanPai)
	CD21_GameMessageClassID_CD21ZhuangJiaYaoPaiID    CD21_GameMessageClassID = 10032 //莊家要牌廣播 S2C (CD21_ZhuangYaoPai)
	CD21_GameMessageClassID_CD21ZhuangTuoBaoID       CD21_GameMessageClassID = 10033 //莊家保險廣播 S2C (CD21_UserCheck)
	CD21_GameMessageClassID_CD21UserRecoverBetID     CD21_GameMessageClassID = 10034 //斷線重連下注消息 S2C (CD21_RecoverScene)
	CD21_GameMessageClassID_CD21UserSpecialBetID     CD21_GameMessageClassID = 10035 //廣播特殊下注 S2C (CD21_BroadCastSpecialBet)
	CD21_GameMessageClassID_CD21GameEndReadyTimeID   CD21_GameMessageClassID = 10036 //游戏结束准备时间 S2C (CD21_ReadyTime)
	CD21_GameMessageClassID_CD21UserPokerRecoverID   CD21_GameMessageClassID = 10037 //斷線重連後的牌 S2C (CD21_UserPokerRecover)
	CD21_GameMessageClassID_CD21UserWinLoseRecoverID CD21_GameMessageClassID = 10038 //斷線重連後的籌碼 S2C (CD21_WinLoseRecover)
	//	CD21UserWatchDataID=10039;				//觀戰資訊S2C (CD21_WatchData)
	//	CD21CheckTimeUpID=10040;				//觀戰資訊C2S
	CD21_GameMessageClassID_Cd21SysPoolID            CD21_GameMessageClassID = 10041  //21點水池載入
	CD21_GameMessageClassID_Cd21SysPoolRetID         CD21_GameMessageClassID = 10042  //21點水池結果
	CD21_GameMessageClassID_CD21GetAllUserBetID      CD21_GameMessageClassID = 10043  //玩家已下注多少金額 S2C
	CD21_GameMessageClassID_CD21GameNoInfo           CD21_GameMessageClassID = 100044 //牌局 S2C
	CD21_GameMessageClassID_CD21BroadCastDoublePaiID CD21_GameMessageClassID = 100045 //加倍后牌型广播 S2C
	CD21_GameMessageClassID_CD21BroadCastBaoXianRet  CD21_GameMessageClassID = 100046 //保险结果广播 S2C
	CD21_GameMessageClassID_UserStopBet              CD21_GameMessageClassID = 100047 //用户结束下注 C2S
	CD21_GameMessageClassID_CD21BroadCastStopUserBet CD21_GameMessageClassID = 100048 //用户结束下注广播 S2C
	CD21_GameMessageClassID_CD21BroadCastStopBaoXian CD21_GameMessageClassID = 100049 //用户结束保险广播 S2C
)

// Enum value maps for CD21_GameMessageClassID.
var (
	CD21_GameMessageClassID_name = map[int32]string{
		10001:  "CD21SettingID",
		10002:  "CD21StateID",
		10003:  "CD21BetID",
		10004:  "CD21BroadCastBetID",
		10005:  "CD21GameSceneID",
		100006: "CD21BetInfoID",
		100007: "CD21BetRetID",
		10011:  "CD21BroadCastFapaiID",
		10012:  "CD21UserOptID",
		10014:  "CD21GameResultWinID",
		10018:  "CD21UserInsurence",
		10019:  "CD21UserHit",
		10020:  "CD21UserStand",
		10021:  "CD21UserSplit",
		10022:  "CD21UserDouble",
		10023:  "CD21UserCancel",
		10024:  "CD21CurrOptID",
		10025:  "CD21StandBroadCastID",
		10029:  "CD21SplitCardBroacast",
		10030:  "CD21BroadCastYaopaiID",
		10031:  "CD21BroadCastFanPaiID",
		10032:  "CD21ZhuangJiaYaoPaiID",
		10033:  "CD21ZhuangTuoBaoID",
		10034:  "CD21UserRecoverBetID",
		10035:  "CD21UserSpecialBetID",
		10036:  "CD21GameEndReadyTimeID",
		10037:  "CD21UserPokerRecoverID",
		10038:  "CD21UserWinLoseRecoverID",
		10041:  "Cd21SysPoolID",
		10042:  "Cd21SysPoolRetID",
		10043:  "CD21GetAllUserBetID",
		100044: "CD21GameNoInfo",
		100045: "CD21BroadCastDoublePaiID",
		100046: "CD21BroadCastBaoXianRet",
		100047: "UserStopBet",
		100048: "CD21BroadCastStopUserBet",
		100049: "CD21BroadCastStopBaoXian",
	}
	CD21_GameMessageClassID_value = map[string]int32{
		"CD21SettingID":            10001,
		"CD21StateID":              10002,
		"CD21BetID":                10003,
		"CD21BroadCastBetID":       10004,
		"CD21GameSceneID":          10005,
		"CD21BetInfoID":            100006,
		"CD21BetRetID":             100007,
		"CD21BroadCastFapaiID":     10011,
		"CD21UserOptID":            10012,
		"CD21GameResultWinID":      10014,
		"CD21UserInsurence":        10018,
		"CD21UserHit":              10019,
		"CD21UserStand":            10020,
		"CD21UserSplit":            10021,
		"CD21UserDouble":           10022,
		"CD21UserCancel":           10023,
		"CD21CurrOptID":            10024,
		"CD21StandBroadCastID":     10025,
		"CD21SplitCardBroacast":    10029,
		"CD21BroadCastYaopaiID":    10030,
		"CD21BroadCastFanPaiID":    10031,
		"CD21ZhuangJiaYaoPaiID":    10032,
		"CD21ZhuangTuoBaoID":       10033,
		"CD21UserRecoverBetID":     10034,
		"CD21UserSpecialBetID":     10035,
		"CD21GameEndReadyTimeID":   10036,
		"CD21UserPokerRecoverID":   10037,
		"CD21UserWinLoseRecoverID": 10038,
		"Cd21SysPoolID":            10041,
		"Cd21SysPoolRetID":         10042,
		"CD21GetAllUserBetID":      10043,
		"CD21GameNoInfo":           100044,
		"CD21BroadCastDoublePaiID": 100045,
		"CD21BroadCastBaoXianRet":  100046,
		"UserStopBet":              100047,
		"CD21BroadCastStopUserBet": 100048,
		"CD21BroadCastStopBaoXian": 100049,
	}
)

func (x CD21_GameMessageClassID) Enum() *CD21_GameMessageClassID {
	p := new(CD21_GameMessageClassID)
	*p = x
	return p
}

func (x CD21_GameMessageClassID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CD21_GameMessageClassID) Descriptor() protoreflect.EnumDescriptor {
	return file_Game_CD21_proto_enumTypes[0].Descriptor()
}

func (CD21_GameMessageClassID) Type() protoreflect.EnumType {
	return &file_Game_CD21_proto_enumTypes[0]
}

func (x CD21_GameMessageClassID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *CD21_GameMessageClassID) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = CD21_GameMessageClassID(num)
	return nil
}

// Deprecated: Use CD21_GameMessageClassID.Descriptor instead.
func (CD21_GameMessageClassID) EnumDescriptor() ([]byte, []int) {
	return file_Game_CD21_proto_rawDescGZIP(), []int{0}
}

var File_Game_CD21_proto protoreflect.FileDescriptor

var file_Game_CD21_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x47, 0x61, 0x6d, 0x65, 0x5f, 0x43, 0x44, 0x32, 0x31, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xa0, 0x07, 0x0a, 0x17,
	0x43, 0x44, 0x32, 0x31, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x0d, 0x43, 0x44, 0x32, 0x31, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x10, 0x91, 0x4e, 0x12, 0x10, 0x0a, 0x0b, 0x43,
	0x44, 0x32, 0x31, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x44, 0x10, 0x92, 0x4e, 0x12, 0x0e, 0x0a,
	0x09, 0x43, 0x44, 0x32, 0x31, 0x42, 0x65, 0x74, 0x49, 0x44, 0x10, 0x93, 0x4e, 0x12, 0x17, 0x0a,
	0x12, 0x43, 0x44, 0x32, 0x31, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x42, 0x65,
	0x74, 0x49, 0x44, 0x10, 0x94, 0x4e, 0x12, 0x14, 0x0a, 0x0f, 0x43, 0x44, 0x32, 0x31, 0x47, 0x61,
	0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x44, 0x10, 0x95, 0x4e, 0x12, 0x13, 0x0a, 0x0d,
	0x43, 0x44, 0x32, 0x31, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x10, 0xa6, 0x8d,
	0x06, 0x12, 0x12, 0x0a, 0x0c, 0x43, 0x44, 0x32, 0x31, 0x42, 0x65, 0x74, 0x52, 0x65, 0x74, 0x49,
	0x44, 0x10, 0xa7, 0x8d, 0x06, 0x12, 0x19, 0x0a, 0x14, 0x43, 0x44, 0x32, 0x31, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x46, 0x61, 0x70, 0x61, 0x69, 0x49, 0x44, 0x10, 0x9b, 0x4e,
	0x12, 0x12, 0x0a, 0x0d, 0x43, 0x44, 0x32, 0x31, 0x55, 0x73, 0x65, 0x72, 0x4f, 0x70, 0x74, 0x49,
	0x44, 0x10, 0x9c, 0x4e, 0x12, 0x18, 0x0a, 0x13, 0x43, 0x44, 0x32, 0x31, 0x47, 0x61, 0x6d, 0x65,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x57, 0x69, 0x6e, 0x49, 0x44, 0x10, 0x9e, 0x4e, 0x12, 0x16,
	0x0a, 0x11, 0x43, 0x44, 0x32, 0x31, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x73, 0x75, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x10, 0xa2, 0x4e, 0x12, 0x10, 0x0a, 0x0b, 0x43, 0x44, 0x32, 0x31, 0x55, 0x73,
	0x65, 0x72, 0x48, 0x69, 0x74, 0x10, 0xa3, 0x4e, 0x12, 0x12, 0x0a, 0x0d, 0x43, 0x44, 0x32, 0x31,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x10, 0xa4, 0x4e, 0x12, 0x12, 0x0a, 0x0d,
	0x43, 0x44, 0x32, 0x31, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x10, 0xa5, 0x4e,
	0x12, 0x13, 0x0a, 0x0e, 0x43, 0x44, 0x32, 0x31, 0x55, 0x73, 0x65, 0x72, 0x44, 0x6f, 0x75, 0x62,
	0x6c, 0x65, 0x10, 0xa6, 0x4e, 0x12, 0x13, 0x0a, 0x0e, 0x43, 0x44, 0x32, 0x31, 0x55, 0x73, 0x65,
	0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x10, 0xa7, 0x4e, 0x12, 0x12, 0x0a, 0x0d, 0x43, 0x44,
	0x32, 0x31, 0x43, 0x75, 0x72, 0x72, 0x4f, 0x70, 0x74, 0x49, 0x44, 0x10, 0xa8, 0x4e, 0x12, 0x19,
	0x0a, 0x14, 0x43, 0x44, 0x32, 0x31, 0x53, 0x74, 0x61, 0x6e, 0x64, 0x42, 0x72, 0x6f, 0x61, 0x64,
	0x43, 0x61, 0x73, 0x74, 0x49, 0x44, 0x10, 0xa9, 0x4e, 0x12, 0x1a, 0x0a, 0x15, 0x43, 0x44, 0x32,
	0x31, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x43, 0x61, 0x72, 0x64, 0x42, 0x72, 0x6f, 0x61, 0x63, 0x61,
	0x73, 0x74, 0x10, 0xad, 0x4e, 0x12, 0x1a, 0x0a, 0x15, 0x43, 0x44, 0x32, 0x31, 0x42, 0x72, 0x6f,
	0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x59, 0x61, 0x6f, 0x70, 0x61, 0x69, 0x49, 0x44, 0x10, 0xae,
	0x4e, 0x12, 0x1a, 0x0a, 0x15, 0x43, 0x44, 0x32, 0x31, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61,
	0x73, 0x74, 0x46, 0x61, 0x6e, 0x50, 0x61, 0x69, 0x49, 0x44, 0x10, 0xaf, 0x4e, 0x12, 0x1a, 0x0a,
	0x15, 0x43, 0x44, 0x32, 0x31, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x4a, 0x69, 0x61, 0x59, 0x61,
	0x6f, 0x50, 0x61, 0x69, 0x49, 0x44, 0x10, 0xb0, 0x4e, 0x12, 0x17, 0x0a, 0x12, 0x43, 0x44, 0x32,
	0x31, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x54, 0x75, 0x6f, 0x42, 0x61, 0x6f, 0x49, 0x44, 0x10,
	0xb1, 0x4e, 0x12, 0x19, 0x0a, 0x14, 0x43, 0x44, 0x32, 0x31, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x63, 0x6f, 0x76, 0x65, 0x72, 0x42, 0x65, 0x74, 0x49, 0x44, 0x10, 0xb2, 0x4e, 0x12, 0x19, 0x0a,
	0x14, 0x43, 0x44, 0x32, 0x31, 0x55, 0x73, 0x65, 0x72, 0x53, 0x70, 0x65, 0x63, 0x69, 0x61, 0x6c,
	0x42, 0x65, 0x74, 0x49, 0x44, 0x10, 0xb3, 0x4e, 0x12, 0x1b, 0x0a, 0x16, 0x43, 0x44, 0x32, 0x31,
	0x47, 0x61, 0x6d, 0x65, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x69, 0x6d, 0x65,
	0x49, 0x44, 0x10, 0xb4, 0x4e, 0x12, 0x1b, 0x0a, 0x16, 0x43, 0x44, 0x32, 0x31, 0x55, 0x73, 0x65,
	0x72, 0x50, 0x6f, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x44, 0x10,
	0xb5, 0x4e, 0x12, 0x1d, 0x0a, 0x18, 0x43, 0x44, 0x32, 0x31, 0x55, 0x73, 0x65, 0x72, 0x57, 0x69,
	0x6e, 0x4c, 0x6f, 0x73, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x49, 0x44, 0x10, 0xb6,
	0x4e, 0x12, 0x12, 0x0a, 0x0d, 0x43, 0x64, 0x32, 0x31, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x6f, 0x6c,
	0x49, 0x44, 0x10, 0xb9, 0x4e, 0x12, 0x15, 0x0a, 0x10, 0x43, 0x64, 0x32, 0x31, 0x53, 0x79, 0x73,
	0x50, 0x6f, 0x6f, 0x6c, 0x52, 0x65, 0x74, 0x49, 0x44, 0x10, 0xba, 0x4e, 0x12, 0x18, 0x0a, 0x13,
	0x43, 0x44, 0x32, 0x31, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x42, 0x65,
	0x74, 0x49, 0x44, 0x10, 0xbb, 0x4e, 0x12, 0x14, 0x0a, 0x0e, 0x43, 0x44, 0x32, 0x31, 0x47, 0x61,
	0x6d, 0x65, 0x4e, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0xcc, 0x8d, 0x06, 0x12, 0x1e, 0x0a, 0x18,
	0x43, 0x44, 0x32, 0x31, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x44, 0x6f, 0x75,
	0x62, 0x6c, 0x65, 0x50, 0x61, 0x69, 0x49, 0x44, 0x10, 0xcd, 0x8d, 0x06, 0x12, 0x1d, 0x0a, 0x17,
	0x43, 0x44, 0x32, 0x31, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x42, 0x61, 0x6f,
	0x58, 0x69, 0x61, 0x6e, 0x52, 0x65, 0x74, 0x10, 0xce, 0x8d, 0x06, 0x12, 0x11, 0x0a, 0x0b, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x74, 0x6f, 0x70, 0x42, 0x65, 0x74, 0x10, 0xcf, 0x8d, 0x06, 0x12, 0x1e,
	0x0a, 0x18, 0x43, 0x44, 0x32, 0x31, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x53,
	0x74, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x42, 0x65, 0x74, 0x10, 0xd0, 0x8d, 0x06, 0x12, 0x1e,
	0x0a, 0x18, 0x43, 0x44, 0x32, 0x31, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x53,
	0x74, 0x6f, 0x70, 0x42, 0x61, 0x6f, 0x58, 0x69, 0x61, 0x6e, 0x10, 0xd1, 0x8d, 0x06, 0x42, 0x0c,
	0x5a, 0x0a, 0x2e, 0x2f, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_Game_CD21_proto_rawDescOnce sync.Once
	file_Game_CD21_proto_rawDescData = file_Game_CD21_proto_rawDesc
)

func file_Game_CD21_proto_rawDescGZIP() []byte {
	file_Game_CD21_proto_rawDescOnce.Do(func() {
		file_Game_CD21_proto_rawDescData = protoimpl.X.CompressGZIP(file_Game_CD21_proto_rawDescData)
	})
	return file_Game_CD21_proto_rawDescData
}

var file_Game_CD21_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Game_CD21_proto_goTypes = []interface{}{
	(CD21_GameMessageClassID)(0), // 0: netproto.CD21_GameMessageClassID
}
var file_Game_CD21_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Game_CD21_proto_init() }
func file_Game_CD21_proto_init() {
	if File_Game_CD21_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Game_CD21_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Game_CD21_proto_goTypes,
		DependencyIndexes: file_Game_CD21_proto_depIdxs,
		EnumInfos:         file_Game_CD21_proto_enumTypes,
	}.Build()
	File_Game_CD21_proto = out.File
	file_Game_CD21_proto_rawDesc = nil
	file_Game_CD21_proto_goTypes = nil
	file_Game_CD21_proto_depIdxs = nil
}
