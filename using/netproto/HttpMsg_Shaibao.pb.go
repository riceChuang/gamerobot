// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: HttpMsg_Shaibao.proto

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

type Http_Shaibao_GameInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableList []*Http_Shaibao_TableList `protobuf:"bytes,1,rep,name=TableList" json:"TableList,omitempty"` //桌子信息
}

func (x *Http_Shaibao_GameInfo) Reset() {
	*x = Http_Shaibao_GameInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HttpMsg_Shaibao_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_Shaibao_GameInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_Shaibao_GameInfo) ProtoMessage() {}

func (x *Http_Shaibao_GameInfo) ProtoReflect() protoreflect.Message {
	mi := &file_HttpMsg_Shaibao_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_Shaibao_GameInfo.ProtoReflect.Descriptor instead.
func (*Http_Shaibao_GameInfo) Descriptor() ([]byte, []int) {
	return file_HttpMsg_Shaibao_proto_rawDescGZIP(), []int{0}
}

func (x *Http_Shaibao_GameInfo) GetTableList() []*Http_Shaibao_TableList {
	if x != nil {
		return x.TableList
	}
	return nil
}

type Http_Shaibao_TableList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableNo           *int32 `protobuf:"varint,1,req,name=TableNo" json:"TableNo,omitempty"`                     //桌子号
	Status            *int32 `protobuf:"varint,2,req,name=Status" json:"Status,omitempty"`                       //当前状态 0 下注中 1下注结束动画 2开奖中 3开奖动画 4结算动画 5等待开始游戏
	GameUserCount     *int32 `protobuf:"varint,3,req,name=GameUserCount" json:"GameUserCount,omitempty"`         //玩家人数(包含未下注玩家)
	RobotCount        *int32 `protobuf:"varint,4,req,name=RobotCount" json:"RobotCount,omitempty"`               //机器人数(包含未下注机器人)
	SumEarnAmount     *int64 `protobuf:"varint,5,req,name=SumEarnAmount" json:"SumEarnAmount,omitempty"`         //系统从开服到现在的总收益
	LastEarnAmount    *int32 `protobuf:"varint,6,req,name=LastEarnAmount" json:"LastEarnAmount,omitempty"`       //最后一局收益
	PoolCurrentAmount *int64 `protobuf:"varint,7,req,name=PoolCurrentAmount" json:"PoolCurrentAmount,omitempty"` //当前水池金币数
}

func (x *Http_Shaibao_TableList) Reset() {
	*x = Http_Shaibao_TableList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HttpMsg_Shaibao_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_Shaibao_TableList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_Shaibao_TableList) ProtoMessage() {}

func (x *Http_Shaibao_TableList) ProtoReflect() protoreflect.Message {
	mi := &file_HttpMsg_Shaibao_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_Shaibao_TableList.ProtoReflect.Descriptor instead.
func (*Http_Shaibao_TableList) Descriptor() ([]byte, []int) {
	return file_HttpMsg_Shaibao_proto_rawDescGZIP(), []int{1}
}

func (x *Http_Shaibao_TableList) GetTableNo() int32 {
	if x != nil && x.TableNo != nil {
		return *x.TableNo
	}
	return 0
}

func (x *Http_Shaibao_TableList) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *Http_Shaibao_TableList) GetGameUserCount() int32 {
	if x != nil && x.GameUserCount != nil {
		return *x.GameUserCount
	}
	return 0
}

func (x *Http_Shaibao_TableList) GetRobotCount() int32 {
	if x != nil && x.RobotCount != nil {
		return *x.RobotCount
	}
	return 0
}

func (x *Http_Shaibao_TableList) GetSumEarnAmount() int64 {
	if x != nil && x.SumEarnAmount != nil {
		return *x.SumEarnAmount
	}
	return 0
}

func (x *Http_Shaibao_TableList) GetLastEarnAmount() int32 {
	if x != nil && x.LastEarnAmount != nil {
		return *x.LastEarnAmount
	}
	return 0
}

func (x *Http_Shaibao_TableList) GetPoolCurrentAmount() int64 {
	if x != nil && x.PoolCurrentAmount != nil {
		return *x.PoolCurrentAmount
	}
	return 0
}

type Http_Shaibao_TableDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableNo                  *int32                   `protobuf:"varint,1,req,name=TableNo" json:"TableNo,omitempty"`                                    //桌子号
	Status                   *int32                   `protobuf:"varint,2,req,name=Status" json:"Status,omitempty"`                                      //当前状态 0 下注中 1下注结束动画 2开奖中 3开奖动画 4结算动画 5等待开始游戏
	SumEarnAmount            *int64                   `protobuf:"varint,3,req,name=SumEarnAmount" json:"SumEarnAmount,omitempty"`                        //系统从开服到现在的总收益
	LastEarnAmount           *int32                   `protobuf:"varint,4,req,name=LastEarnAmount" json:"LastEarnAmount,omitempty"`                      //最后一局收益
	PoolCurrentAmount        *int64                   `protobuf:"varint,5,req,name=PoolCurrentAmount" json:"PoolCurrentAmount,omitempty"`                //当前水池金币数
	GameUserCount            *int32                   `protobuf:"varint,6,req,name=GameUserCount" json:"GameUserCount,omitempty"`                        //游戏玩家人数(包含未下注玩家)
	RobotCount               *int32                   `protobuf:"varint,7,req,name=RobotCount" json:"RobotCount,omitempty"`                              //机器人数(包含未下注机器人)
	PopWaterRate             *int32                   `protobuf:"varint,8,req,name=PopWaterRate" json:"PopWaterRate,omitempty"`                          //当前注水率
	ZhuangType               *int32                   `protobuf:"varint,9,req,name=ZhuangType" json:"ZhuangType,omitempty"`                              //庄家类型	0系统 1机器人当庄 2玩家上庄
	ZhuangNickName           *string                  `protobuf:"bytes,10,req,name=ZhuangNickName" json:"ZhuangNickName,omitempty"`                      //庄家昵称
	ZhuangUserID             *int32                   `protobuf:"varint,11,req,name=ZhuangUserID" json:"ZhuangUserID,omitempty"`                         //庄家用户ID
	ZhuangSurplusTotal       *int32                   `protobuf:"varint,12,req,name=ZhuangSurplusTotal" json:"ZhuangSurplusTotal,omitempty"`             //庄家剩余局数
	ZhuangCurrentMoneyAmount *int64                   `protobuf:"varint,13,req,name=ZhuangCurrentMoneyAmount" json:"ZhuangCurrentMoneyAmount,omitempty"` //庄家当前金币数 系统或机器人坐庄则表示水池金币数量
	ZhuangInitMoneyAmount    *int64                   `protobuf:"varint,14,req,name=ZhuangInitMoneyAmount" json:"ZhuangInitMoneyAmount,omitempty"`       //庄家初始上庄金币数 系统或机器人坐庄则用0表示
	OpenResultZhuang         *int64                   `protobuf:"varint,15,req,name=OpenResultZhuang" json:"OpenResultZhuang,omitempty"`                 //当前开奖结果为庄的收益数量
	OpenResultXian           *int64                   `protobuf:"varint,16,req,name=OpenResultXian" json:"OpenResultXian,omitempty"`                     //当前开奖结果为闲的收益数量
	OpenResultHe             *int64                   `protobuf:"varint,17,req,name=OpenResultHe" json:"OpenResultHe,omitempty"`                         //当前开奖结果为和的收益数量
	CurentOpenResult         *string                  `protobuf:"bytes,18,req,name=CurentOpenResult" json:"CurentOpenResult,omitempty"`                  //当前开奖结果,只有在开奖动画状态下有值
	CurentOpenResultEarn     *int64                   `protobuf:"varint,19,req,name=CurentOpenResultEarn" json:"CurentOpenResultEarn,omitempty"`         //当前开奖收益,只有在开奖动画状态下有值
	UserInfo                 []*Http_Shaibao_UserInfo `protobuf:"bytes,20,rep,name=UserInfo" json:"UserInfo,omitempty"`                                  //玩家信息
}

func (x *Http_Shaibao_TableDetail) Reset() {
	*x = Http_Shaibao_TableDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HttpMsg_Shaibao_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_Shaibao_TableDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_Shaibao_TableDetail) ProtoMessage() {}

func (x *Http_Shaibao_TableDetail) ProtoReflect() protoreflect.Message {
	mi := &file_HttpMsg_Shaibao_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_Shaibao_TableDetail.ProtoReflect.Descriptor instead.
func (*Http_Shaibao_TableDetail) Descriptor() ([]byte, []int) {
	return file_HttpMsg_Shaibao_proto_rawDescGZIP(), []int{2}
}

func (x *Http_Shaibao_TableDetail) GetTableNo() int32 {
	if x != nil && x.TableNo != nil {
		return *x.TableNo
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetSumEarnAmount() int64 {
	if x != nil && x.SumEarnAmount != nil {
		return *x.SumEarnAmount
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetLastEarnAmount() int32 {
	if x != nil && x.LastEarnAmount != nil {
		return *x.LastEarnAmount
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetPoolCurrentAmount() int64 {
	if x != nil && x.PoolCurrentAmount != nil {
		return *x.PoolCurrentAmount
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetGameUserCount() int32 {
	if x != nil && x.GameUserCount != nil {
		return *x.GameUserCount
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetRobotCount() int32 {
	if x != nil && x.RobotCount != nil {
		return *x.RobotCount
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetPopWaterRate() int32 {
	if x != nil && x.PopWaterRate != nil {
		return *x.PopWaterRate
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetZhuangType() int32 {
	if x != nil && x.ZhuangType != nil {
		return *x.ZhuangType
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetZhuangNickName() string {
	if x != nil && x.ZhuangNickName != nil {
		return *x.ZhuangNickName
	}
	return ""
}

func (x *Http_Shaibao_TableDetail) GetZhuangUserID() int32 {
	if x != nil && x.ZhuangUserID != nil {
		return *x.ZhuangUserID
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetZhuangSurplusTotal() int32 {
	if x != nil && x.ZhuangSurplusTotal != nil {
		return *x.ZhuangSurplusTotal
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetZhuangCurrentMoneyAmount() int64 {
	if x != nil && x.ZhuangCurrentMoneyAmount != nil {
		return *x.ZhuangCurrentMoneyAmount
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetZhuangInitMoneyAmount() int64 {
	if x != nil && x.ZhuangInitMoneyAmount != nil {
		return *x.ZhuangInitMoneyAmount
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetOpenResultZhuang() int64 {
	if x != nil && x.OpenResultZhuang != nil {
		return *x.OpenResultZhuang
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetOpenResultXian() int64 {
	if x != nil && x.OpenResultXian != nil {
		return *x.OpenResultXian
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetOpenResultHe() int64 {
	if x != nil && x.OpenResultHe != nil {
		return *x.OpenResultHe
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetCurentOpenResult() string {
	if x != nil && x.CurentOpenResult != nil {
		return *x.CurentOpenResult
	}
	return ""
}

func (x *Http_Shaibao_TableDetail) GetCurentOpenResultEarn() int64 {
	if x != nil && x.CurentOpenResultEarn != nil {
		return *x.CurentOpenResultEarn
	}
	return 0
}

func (x *Http_Shaibao_TableDetail) GetUserInfo() []*Http_Shaibao_UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

type Http_Shaibao_UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID          *int32  `protobuf:"varint,1,req,name=UserID" json:"UserID,omitempty"`                    //玩家ID
	UserType        *int32  `protobuf:"varint,2,req,name=UserType" json:"UserType,omitempty"`                //玩家类型 0机器人 1玩家
	NickName        *string `protobuf:"bytes,3,req,name=NickName" json:"NickName,omitempty"`                 //玩家昵称
	GameScore       *int64  `protobuf:"varint,4,req,name=GameScore" json:"GameScore,omitempty"`              //游戏总输赢
	InitMoney       *int64  `protobuf:"varint,5,req,name=InitMoney" json:"InitMoney,omitempty"`              //初始金币
	CurrentMoney    *int64  `protobuf:"varint,6,req,name=CurrentMoney" json:"CurrentMoney,omitempty"`        //当前金币
	BetInfoZhuang   *int64  `protobuf:"varint,7,req,name=BetInfoZhuang" json:"BetInfoZhuang,omitempty"`      //下注庄的金币数量
	BetInfoZhuangDZ *int64  `protobuf:"varint,8,req,name=BetInfoZhuangDZ" json:"BetInfoZhuangDZ,omitempty"`  //下注庄对子的金币数量
	BetInfoZhuangTW *int64  `protobuf:"varint,9,req,name=BetInfoZhuangTW" json:"BetInfoZhuangTW,omitempty"`  //下注庄天王的金币数量
	BetInfoZhuangSZ *int64  `protobuf:"varint,10,req,name=BetInfoZhuangSZ" json:"BetInfoZhuangSZ,omitempty"` //下注庄顺子的金币数量
	BetInfoXian     *int64  `protobuf:"varint,11,req,name=BetInfoXian" json:"BetInfoXian,omitempty"`         //下注闲的金币数量
	BetInfoXianDZ   *int64  `protobuf:"varint,12,req,name=BetInfoXianDZ" json:"BetInfoXianDZ,omitempty"`     //下注闲对子的金币数量
	BetInfoXianTW   *int64  `protobuf:"varint,13,req,name=BetInfoXianTW" json:"BetInfoXianTW,omitempty"`     //下注闲天王的金币数量
	BetInfoXianSZ   *int64  `protobuf:"varint,14,req,name=BetInfoXianSZ" json:"BetInfoXianSZ,omitempty"`     //下注闲顺子的金币数量
	BetInfoHe       *int64  `protobuf:"varint,15,req,name=BetInfoHe" json:"BetInfoHe,omitempty"`             //下注和的金币数量
	BetInfoHeTD     *int64  `protobuf:"varint,16,req,name=BetInfoHeTD" json:"BetInfoHeTD,omitempty"`         //下注同点和的金币数量
}

func (x *Http_Shaibao_UserInfo) Reset() {
	*x = Http_Shaibao_UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_HttpMsg_Shaibao_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Http_Shaibao_UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Http_Shaibao_UserInfo) ProtoMessage() {}

func (x *Http_Shaibao_UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_HttpMsg_Shaibao_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Http_Shaibao_UserInfo.ProtoReflect.Descriptor instead.
func (*Http_Shaibao_UserInfo) Descriptor() ([]byte, []int) {
	return file_HttpMsg_Shaibao_proto_rawDescGZIP(), []int{3}
}

func (x *Http_Shaibao_UserInfo) GetUserID() int32 {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetUserType() int32 {
	if x != nil && x.UserType != nil {
		return *x.UserType
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetNickName() string {
	if x != nil && x.NickName != nil {
		return *x.NickName
	}
	return ""
}

func (x *Http_Shaibao_UserInfo) GetGameScore() int64 {
	if x != nil && x.GameScore != nil {
		return *x.GameScore
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetInitMoney() int64 {
	if x != nil && x.InitMoney != nil {
		return *x.InitMoney
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetCurrentMoney() int64 {
	if x != nil && x.CurrentMoney != nil {
		return *x.CurrentMoney
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetBetInfoZhuang() int64 {
	if x != nil && x.BetInfoZhuang != nil {
		return *x.BetInfoZhuang
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetBetInfoZhuangDZ() int64 {
	if x != nil && x.BetInfoZhuangDZ != nil {
		return *x.BetInfoZhuangDZ
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetBetInfoZhuangTW() int64 {
	if x != nil && x.BetInfoZhuangTW != nil {
		return *x.BetInfoZhuangTW
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetBetInfoZhuangSZ() int64 {
	if x != nil && x.BetInfoZhuangSZ != nil {
		return *x.BetInfoZhuangSZ
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetBetInfoXian() int64 {
	if x != nil && x.BetInfoXian != nil {
		return *x.BetInfoXian
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetBetInfoXianDZ() int64 {
	if x != nil && x.BetInfoXianDZ != nil {
		return *x.BetInfoXianDZ
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetBetInfoXianTW() int64 {
	if x != nil && x.BetInfoXianTW != nil {
		return *x.BetInfoXianTW
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetBetInfoXianSZ() int64 {
	if x != nil && x.BetInfoXianSZ != nil {
		return *x.BetInfoXianSZ
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetBetInfoHe() int64 {
	if x != nil && x.BetInfoHe != nil {
		return *x.BetInfoHe
	}
	return 0
}

func (x *Http_Shaibao_UserInfo) GetBetInfoHeTD() int64 {
	if x != nil && x.BetInfoHeTD != nil {
		return *x.BetInfoHeTD
	}
	return 0
}

var File_HttpMsg_Shaibao_proto protoreflect.FileDescriptor

var file_HttpMsg_Shaibao_proto_rawDesc = []byte{
	0x0a, 0x15, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x73, 0x67, 0x5f, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x57, 0x0a, 0x15, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61,
	0x6f, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3e, 0x0a, 0x09, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e,
	0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x53, 0x68,
	0x61, 0x69, 0x62, 0x61, 0x6f, 0x5f, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x8c, 0x02, 0x0a, 0x16, 0x48,
	0x74, 0x74, 0x70, 0x5f, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f, 0x5f, 0x54, 0x61, 0x62, 0x6c,
	0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52,
	0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x47, 0x61, 0x6d, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0d,
	0x47, 0x61, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x02, 0x28,
	0x05, 0x52, 0x0a, 0x52, 0x6f, 0x62, 0x6f, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a,
	0x0d, 0x53, 0x75, 0x6d, 0x45, 0x61, 0x72, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x0d, 0x53, 0x75, 0x6d, 0x45, 0x61, 0x72, 0x6e, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x61, 0x72, 0x6e, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0e, 0x4c, 0x61, 0x73,
	0x74, 0x45, 0x61, 0x72, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x50,
	0x6f, 0x6f, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x07, 0x20, 0x02, 0x28, 0x03, 0x52, 0x11, 0x50, 0x6f, 0x6f, 0x6c, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd5, 0x06, 0x0a, 0x18, 0x48, 0x74,
	0x74, 0x70, 0x5f, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f, 0x5f, 0x54, 0x61, 0x62, 0x6c, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e,
	0x6f, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x07, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05,
	0x52, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x75, 0x6d, 0x45,
	0x61, 0x72, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x03, 0x52,
	0x0d, 0x53, 0x75, 0x6d, 0x45, 0x61, 0x72, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26,
	0x0a, 0x0e, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x61, 0x72, 0x6e, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x04, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0e, 0x4c, 0x61, 0x73, 0x74, 0x45, 0x61, 0x72, 0x6e,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x50, 0x6f, 0x6f, 0x6c, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x02, 0x28,
	0x03, 0x52, 0x11, 0x50, 0x6f, 0x6f, 0x6c, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0d, 0x47, 0x61, 0x6d,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x52, 0x6f,
	0x62, 0x6f, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0a,
	0x52, 0x6f, 0x62, 0x6f, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x6f,
	0x70, 0x57, 0x61, 0x74, 0x65, 0x72, 0x52, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x02, 0x28, 0x05,
	0x52, 0x0c, 0x50, 0x6f, 0x70, 0x57, 0x61, 0x74, 0x65, 0x72, 0x52, 0x61, 0x74, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x0a, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x26,
	0x0a, 0x0e, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x0a, 0x20, 0x02, 0x28, 0x09, 0x52, 0x0e, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x4e, 0x69,
	0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0c, 0x5a, 0x68,
	0x75, 0x61, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x12, 0x5a, 0x68,
	0x75, 0x61, 0x6e, 0x67, 0x53, 0x75, 0x72, 0x70, 0x6c, 0x75, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x0c, 0x20, 0x02, 0x28, 0x05, 0x52, 0x12, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x53, 0x75,
	0x72, 0x70, 0x6c, 0x75, 0x73, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x3a, 0x0a, 0x18, 0x5a, 0x68,
	0x75, 0x61, 0x6e, 0x67, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x0d, 0x20, 0x02, 0x28, 0x03, 0x52, 0x18, 0x5a, 0x68,
	0x75, 0x61, 0x6e, 0x67, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x34, 0x0a, 0x15, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67,
	0x49, 0x6e, 0x69, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0e, 0x20, 0x02, 0x28, 0x03, 0x52, 0x15, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x49, 0x6e, 0x69,
	0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x10,
	0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67,
	0x18, 0x0f, 0x20, 0x02, 0x28, 0x03, 0x52, 0x10, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x12, 0x26, 0x0a, 0x0e, 0x4f, 0x70, 0x65, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x58, 0x69, 0x61, 0x6e, 0x18, 0x10, 0x20, 0x02, 0x28, 0x03,
	0x52, 0x0e, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x58, 0x69, 0x61, 0x6e,
	0x12, 0x22, 0x0a, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x48, 0x65,
	0x18, 0x11, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0c, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x48, 0x65, 0x12, 0x2a, 0x0a, 0x10, 0x43, 0x75, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x70,
	0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x12, 0x20, 0x02, 0x28, 0x09, 0x52, 0x10,
	0x43, 0x75, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x32, 0x0a, 0x14, 0x43, 0x75, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x45, 0x61, 0x72, 0x6e, 0x18, 0x13, 0x20, 0x02, 0x28, 0x03, 0x52, 0x14,
	0x43, 0x75, 0x72, 0x65, 0x6e, 0x74, 0x4f, 0x70, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x45, 0x61, 0x72, 0x6e, 0x12, 0x3b, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x18, 0x14, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f, 0x5f, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x22, 0xbf, 0x04, 0x0a, 0x15, 0x48, 0x74, 0x74, 0x70, 0x5f, 0x53, 0x68, 0x61, 0x69, 0x62,
	0x61, 0x6f, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28,
	0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x18, 0x04, 0x20, 0x02, 0x28, 0x03, 0x52, 0x09,
	0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x49, 0x6e, 0x69,
	0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x05, 0x20, 0x02, 0x28, 0x03, 0x52, 0x09, 0x49, 0x6e,
	0x69, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x43, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x06, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0c, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x42,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x18, 0x07, 0x20, 0x02,
	0x28, 0x03, 0x52, 0x0d, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x5a, 0x68, 0x75, 0x61, 0x6e,
	0x67, 0x12, 0x28, 0x0a, 0x0f, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x5a, 0x68, 0x75, 0x61,
	0x6e, 0x67, 0x44, 0x5a, 0x18, 0x08, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0f, 0x42, 0x65, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x44, 0x5a, 0x12, 0x28, 0x0a, 0x0f, 0x42,
	0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x54, 0x57, 0x18, 0x09,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x0f, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x5a, 0x68, 0x75,
	0x61, 0x6e, 0x67, 0x54, 0x57, 0x12, 0x28, 0x0a, 0x0f, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x53, 0x5a, 0x18, 0x0a, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0f,
	0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x53, 0x5a, 0x12,
	0x20, 0x0a, 0x0b, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x58, 0x69, 0x61, 0x6e, 0x18, 0x0b,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x0b, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x58, 0x69, 0x61,
	0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x58, 0x69, 0x61, 0x6e,
	0x44, 0x5a, 0x18, 0x0c, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0d, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66,
	0x6f, 0x58, 0x69, 0x61, 0x6e, 0x44, 0x5a, 0x12, 0x24, 0x0a, 0x0d, 0x42, 0x65, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x58, 0x69, 0x61, 0x6e, 0x54, 0x57, 0x18, 0x0d, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0d,
	0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x58, 0x69, 0x61, 0x6e, 0x54, 0x57, 0x12, 0x24, 0x0a,
	0x0d, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x58, 0x69, 0x61, 0x6e, 0x53, 0x5a, 0x18, 0x0e,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x0d, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x58, 0x69, 0x61,
	0x6e, 0x53, 0x5a, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x65,
	0x18, 0x0f, 0x20, 0x02, 0x28, 0x03, 0x52, 0x09, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48, 0x65, 0x54, 0x44,
	0x18, 0x10, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0b, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x65, 0x54, 0x44, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f,
}

var (
	file_HttpMsg_Shaibao_proto_rawDescOnce sync.Once
	file_HttpMsg_Shaibao_proto_rawDescData = file_HttpMsg_Shaibao_proto_rawDesc
)

func file_HttpMsg_Shaibao_proto_rawDescGZIP() []byte {
	file_HttpMsg_Shaibao_proto_rawDescOnce.Do(func() {
		file_HttpMsg_Shaibao_proto_rawDescData = protoimpl.X.CompressGZIP(file_HttpMsg_Shaibao_proto_rawDescData)
	})
	return file_HttpMsg_Shaibao_proto_rawDescData
}

var file_HttpMsg_Shaibao_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_HttpMsg_Shaibao_proto_goTypes = []interface{}{
	(*Http_Shaibao_GameInfo)(nil),    // 0: netproto.Http_Shaibao_GameInfo
	(*Http_Shaibao_TableList)(nil),   // 1: netproto.Http_Shaibao_TableList
	(*Http_Shaibao_TableDetail)(nil), // 2: netproto.Http_Shaibao_TableDetail
	(*Http_Shaibao_UserInfo)(nil),    // 3: netproto.Http_Shaibao_UserInfo
}
var file_HttpMsg_Shaibao_proto_depIdxs = []int32{
	1, // 0: netproto.Http_Shaibao_GameInfo.TableList:type_name -> netproto.Http_Shaibao_TableList
	3, // 1: netproto.Http_Shaibao_TableDetail.UserInfo:type_name -> netproto.Http_Shaibao_UserInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_HttpMsg_Shaibao_proto_init() }
func file_HttpMsg_Shaibao_proto_init() {
	if File_HttpMsg_Shaibao_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_HttpMsg_Shaibao_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_Shaibao_GameInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_HttpMsg_Shaibao_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_Shaibao_TableList); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_HttpMsg_Shaibao_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_Shaibao_TableDetail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_HttpMsg_Shaibao_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Http_Shaibao_UserInfo); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_HttpMsg_Shaibao_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_HttpMsg_Shaibao_proto_goTypes,
		DependencyIndexes: file_HttpMsg_Shaibao_proto_depIdxs,
		MessageInfos:      file_HttpMsg_Shaibao_proto_msgTypes,
	}.Build()
	File_HttpMsg_Shaibao_proto = out.File
	file_HttpMsg_Shaibao_proto_rawDesc = nil
	file_HttpMsg_Shaibao_proto_goTypes = nil
	file_HttpMsg_Shaibao_proto_depIdxs = nil
}
