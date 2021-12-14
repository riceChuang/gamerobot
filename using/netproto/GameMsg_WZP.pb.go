// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: GameMsg_WZP.proto

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

type WZP_GameConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseAmount   *int32  `protobuf:"varint,1,req,name=baseAmount" json:"baseAmount,omitempty"`
	BetTimeout   *int32  `protobuf:"varint,2,req,name=betTimeout" json:"betTimeout,omitempty"`
	BetAmount    []int32 `protobuf:"varint,3,rep,name=betAmount" json:"betAmount,omitempty"`
	ReadyTimeout *int32  `protobuf:"varint,4,req,name=ReadyTimeout" json:"ReadyTimeout,omitempty"`
}

func (x *WZP_GameConfig) Reset() {
	*x = WZP_GameConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameMsg_WZP_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WZP_GameConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WZP_GameConfig) ProtoMessage() {}

func (x *WZP_GameConfig) ProtoReflect() protoreflect.Message {
	mi := &file_GameMsg_WZP_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WZP_GameConfig.ProtoReflect.Descriptor instead.
func (*WZP_GameConfig) Descriptor() ([]byte, []int) {
	return file_GameMsg_WZP_proto_rawDescGZIP(), []int{0}
}

func (x *WZP_GameConfig) GetBaseAmount() int32 {
	if x != nil && x.BaseAmount != nil {
		return *x.BaseAmount
	}
	return 0
}

func (x *WZP_GameConfig) GetBetTimeout() int32 {
	if x != nil && x.BetTimeout != nil {
		return *x.BetTimeout
	}
	return 0
}

func (x *WZP_GameConfig) GetBetAmount() []int32 {
	if x != nil {
		return x.BetAmount
	}
	return nil
}

func (x *WZP_GameConfig) GetReadyTimeout() int32 {
	if x != nil && x.ReadyTimeout != nil {
		return *x.ReadyTimeout
	}
	return 0
}

type WZP_NotifyUserGameTrun struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID        *int32 `protobuf:"varint,1,req,name=UserID" json:"UserID,omitempty"`
	TrunNo        *int32 `protobuf:"varint,2,req,name=TrunNo" json:"TrunNo,omitempty"`
	Timeout       *int32 `protobuf:"varint,3,req,name=Timeout" json:"Timeout,omitempty"`
	IsFirst       *int32 `protobuf:"varint,4,req,name=IsFirst" json:"IsFirst,omitempty"`
	IsLast        *int32 `protobuf:"varint,5,req,name=IsLast" json:"IsLast,omitempty"`
	IsSuoha       *int32 `protobuf:"varint,6,req,name=IsSuoha" json:"IsSuoha,omitempty"`
	TurnBetAmount *int64 `protobuf:"varint,7,req,name=TurnBetAmount" json:"TurnBetAmount,omitempty"`
}

func (x *WZP_NotifyUserGameTrun) Reset() {
	*x = WZP_NotifyUserGameTrun{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameMsg_WZP_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WZP_NotifyUserGameTrun) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WZP_NotifyUserGameTrun) ProtoMessage() {}

func (x *WZP_NotifyUserGameTrun) ProtoReflect() protoreflect.Message {
	mi := &file_GameMsg_WZP_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WZP_NotifyUserGameTrun.ProtoReflect.Descriptor instead.
func (*WZP_NotifyUserGameTrun) Descriptor() ([]byte, []int) {
	return file_GameMsg_WZP_proto_rawDescGZIP(), []int{1}
}

func (x *WZP_NotifyUserGameTrun) GetUserID() int32 {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return 0
}

func (x *WZP_NotifyUserGameTrun) GetTrunNo() int32 {
	if x != nil && x.TrunNo != nil {
		return *x.TrunNo
	}
	return 0
}

func (x *WZP_NotifyUserGameTrun) GetTimeout() int32 {
	if x != nil && x.Timeout != nil {
		return *x.Timeout
	}
	return 0
}

func (x *WZP_NotifyUserGameTrun) GetIsFirst() int32 {
	if x != nil && x.IsFirst != nil {
		return *x.IsFirst
	}
	return 0
}

func (x *WZP_NotifyUserGameTrun) GetIsLast() int32 {
	if x != nil && x.IsLast != nil {
		return *x.IsLast
	}
	return 0
}

func (x *WZP_NotifyUserGameTrun) GetIsSuoha() int32 {
	if x != nil && x.IsSuoha != nil {
		return *x.IsSuoha
	}
	return 0
}

func (x *WZP_NotifyUserGameTrun) GetTurnBetAmount() int64 {
	if x != nil && x.TurnBetAmount != nil {
		return *x.TurnBetAmount
	}
	return 0
}

type WZP_UserGameAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ActionID  *int32 `protobuf:"varint,1,req,name=ActionID" json:"ActionID,omitempty"`
	BetAmount *int64 `protobuf:"varint,2,req,name=BetAmount" json:"BetAmount,omitempty"`
}

func (x *WZP_UserGameAction) Reset() {
	*x = WZP_UserGameAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameMsg_WZP_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WZP_UserGameAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WZP_UserGameAction) ProtoMessage() {}

func (x *WZP_UserGameAction) ProtoReflect() protoreflect.Message {
	mi := &file_GameMsg_WZP_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WZP_UserGameAction.ProtoReflect.Descriptor instead.
func (*WZP_UserGameAction) Descriptor() ([]byte, []int) {
	return file_GameMsg_WZP_proto_rawDescGZIP(), []int{2}
}

func (x *WZP_UserGameAction) GetActionID() int32 {
	if x != nil && x.ActionID != nil {
		return *x.ActionID
	}
	return 0
}

func (x *WZP_UserGameAction) GetBetAmount() int64 {
	if x != nil && x.BetAmount != nil {
		return *x.BetAmount
	}
	return 0
}

type WZP_NotifyUserGameAction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID         *int32    `protobuf:"varint,1,req,name=UserID" json:"UserID,omitempty"`
	ActionID       *int32    `protobuf:"varint,2,req,name=ActionID" json:"ActionID,omitempty"`
	BetAmount      *int64    `protobuf:"varint,3,req,name=BetAmount" json:"BetAmount,omitempty"`
	TotalBetAmount *int64    `protobuf:"varint,4,req,name=TotalBetAmount" json:"TotalBetAmount,omitempty"`
	IsSuoha        *int32    `protobuf:"varint,5,req,name=IsSuoha" json:"IsSuoha,omitempty"`
	GameUser       *GameUser `protobuf:"bytes,6,req,name=GameUser" json:"GameUser,omitempty"`
	UserBetAmount  *int64    `protobuf:"varint,7,req,name=UserBetAmount" json:"UserBetAmount,omitempty"`
}

func (x *WZP_NotifyUserGameAction) Reset() {
	*x = WZP_NotifyUserGameAction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameMsg_WZP_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WZP_NotifyUserGameAction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WZP_NotifyUserGameAction) ProtoMessage() {}

func (x *WZP_NotifyUserGameAction) ProtoReflect() protoreflect.Message {
	mi := &file_GameMsg_WZP_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WZP_NotifyUserGameAction.ProtoReflect.Descriptor instead.
func (*WZP_NotifyUserGameAction) Descriptor() ([]byte, []int) {
	return file_GameMsg_WZP_proto_rawDescGZIP(), []int{3}
}

func (x *WZP_NotifyUserGameAction) GetUserID() int32 {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return 0
}

func (x *WZP_NotifyUserGameAction) GetActionID() int32 {
	if x != nil && x.ActionID != nil {
		return *x.ActionID
	}
	return 0
}

func (x *WZP_NotifyUserGameAction) GetBetAmount() int64 {
	if x != nil && x.BetAmount != nil {
		return *x.BetAmount
	}
	return 0
}

func (x *WZP_NotifyUserGameAction) GetTotalBetAmount() int64 {
	if x != nil && x.TotalBetAmount != nil {
		return *x.TotalBetAmount
	}
	return 0
}

func (x *WZP_NotifyUserGameAction) GetIsSuoha() int32 {
	if x != nil && x.IsSuoha != nil {
		return *x.IsSuoha
	}
	return 0
}

func (x *WZP_NotifyUserGameAction) GetGameUser() *GameUser {
	if x != nil {
		return x.GameUser
	}
	return nil
}

func (x *WZP_NotifyUserGameAction) GetUserBetAmount() int64 {
	if x != nil && x.UserBetAmount != nil {
		return *x.UserBetAmount
	}
	return 0
}

type WZP_UserBetLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID *int32 `protobuf:"varint,1,req,name=UserID" json:"UserID,omitempty"`
	Amount *int64 `protobuf:"varint,2,req,name=Amount" json:"Amount,omitempty"`
	TrunNo *int32 `protobuf:"varint,3,req,name=TrunNo" json:"TrunNo,omitempty"`
}

func (x *WZP_UserBetLog) Reset() {
	*x = WZP_UserBetLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameMsg_WZP_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WZP_UserBetLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WZP_UserBetLog) ProtoMessage() {}

func (x *WZP_UserBetLog) ProtoReflect() protoreflect.Message {
	mi := &file_GameMsg_WZP_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WZP_UserBetLog.ProtoReflect.Descriptor instead.
func (*WZP_UserBetLog) Descriptor() ([]byte, []int) {
	return file_GameMsg_WZP_proto_rawDescGZIP(), []int{4}
}

func (x *WZP_UserBetLog) GetUserID() int32 {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return 0
}

func (x *WZP_UserBetLog) GetAmount() int64 {
	if x != nil && x.Amount != nil {
		return *x.Amount
	}
	return 0
}

func (x *WZP_UserBetLog) GetTrunNo() int32 {
	if x != nil && x.TrunNo != nil {
		return *x.TrunNo
	}
	return 0
}

type WZP_GameBetLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BetLogs        []*WZP_UserBetLog `protobuf:"bytes,1,rep,name=BetLogs" json:"BetLogs,omitempty"`
	TotalBetAmount *int32            `protobuf:"varint,2,req,name=TotalBetAmount" json:"TotalBetAmount,omitempty"`
}

func (x *WZP_GameBetLog) Reset() {
	*x = WZP_GameBetLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameMsg_WZP_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WZP_GameBetLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WZP_GameBetLog) ProtoMessage() {}

func (x *WZP_GameBetLog) ProtoReflect() protoreflect.Message {
	mi := &file_GameMsg_WZP_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WZP_GameBetLog.ProtoReflect.Descriptor instead.
func (*WZP_GameBetLog) Descriptor() ([]byte, []int) {
	return file_GameMsg_WZP_proto_rawDescGZIP(), []int{5}
}

func (x *WZP_GameBetLog) GetBetLogs() []*WZP_UserBetLog {
	if x != nil {
		return x.BetLogs
	}
	return nil
}

func (x *WZP_GameBetLog) GetTotalBetAmount() int32 {
	if x != nil && x.TotalBetAmount != nil {
		return *x.TotalBetAmount
	}
	return 0
}

type WZP_UserRet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserID    *int32     `protobuf:"varint,1,req,name=UserID" json:"UserID,omitempty"`
	SeatNo    *int32     `protobuf:"varint,2,req,name=SeatNo" json:"SeatNo,omitempty"`
	CardType  *int32     `protobuf:"varint,3,opt,name=CardType" json:"CardType,omitempty"`
	WinLose   *int64     `protobuf:"varint,4,req,name=WinLose" json:"WinLose,omitempty"`
	BetAmount *int64     `protobuf:"varint,5,req,name=BetAmount" json:"BetAmount,omitempty"`
	UserPoker *UserPoker `protobuf:"bytes,6,opt,name=UserPoker" json:"UserPoker,omitempty"`
}

func (x *WZP_UserRet) Reset() {
	*x = WZP_UserRet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameMsg_WZP_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WZP_UserRet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WZP_UserRet) ProtoMessage() {}

func (x *WZP_UserRet) ProtoReflect() protoreflect.Message {
	mi := &file_GameMsg_WZP_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WZP_UserRet.ProtoReflect.Descriptor instead.
func (*WZP_UserRet) Descriptor() ([]byte, []int) {
	return file_GameMsg_WZP_proto_rawDescGZIP(), []int{6}
}

func (x *WZP_UserRet) GetUserID() int32 {
	if x != nil && x.UserID != nil {
		return *x.UserID
	}
	return 0
}

func (x *WZP_UserRet) GetSeatNo() int32 {
	if x != nil && x.SeatNo != nil {
		return *x.SeatNo
	}
	return 0
}

func (x *WZP_UserRet) GetCardType() int32 {
	if x != nil && x.CardType != nil {
		return *x.CardType
	}
	return 0
}

func (x *WZP_UserRet) GetWinLose() int64 {
	if x != nil && x.WinLose != nil {
		return *x.WinLose
	}
	return 0
}

func (x *WZP_UserRet) GetBetAmount() int64 {
	if x != nil && x.BetAmount != nil {
		return *x.BetAmount
	}
	return 0
}

func (x *WZP_UserRet) GetUserPoker() *UserPoker {
	if x != nil {
		return x.UserPoker
	}
	return nil
}

type WZP_GameRet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GameBetAmount *int64         `protobuf:"varint,1,req,name=GameBetAmount" json:"GameBetAmount,omitempty"`
	UserRets      []*WZP_UserRet `protobuf:"bytes,2,rep,name=UserRets" json:"UserRets,omitempty"`
	GiveUp        *int32         `protobuf:"varint,3,req,name=GiveUp" json:"GiveUp,omitempty"`
	WinUserID     *int32         `protobuf:"varint,4,req,name=WinUserID" json:"WinUserID,omitempty"`
}

func (x *WZP_GameRet) Reset() {
	*x = WZP_GameRet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameMsg_WZP_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WZP_GameRet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WZP_GameRet) ProtoMessage() {}

func (x *WZP_GameRet) ProtoReflect() protoreflect.Message {
	mi := &file_GameMsg_WZP_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WZP_GameRet.ProtoReflect.Descriptor instead.
func (*WZP_GameRet) Descriptor() ([]byte, []int) {
	return file_GameMsg_WZP_proto_rawDescGZIP(), []int{7}
}

func (x *WZP_GameRet) GetGameBetAmount() int64 {
	if x != nil && x.GameBetAmount != nil {
		return *x.GameBetAmount
	}
	return 0
}

func (x *WZP_GameRet) GetUserRets() []*WZP_UserRet {
	if x != nil {
		return x.UserRets
	}
	return nil
}

func (x *WZP_GameRet) GetGiveUp() int32 {
	if x != nil && x.GiveUp != nil {
		return *x.GiveUp
	}
	return 0
}

func (x *WZP_GameRet) GetWinUserID() int32 {
	if x != nil && x.WinUserID != nil {
		return *x.WinUserID
	}
	return 0
}

type WZP_GameScene struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TurnNo         *int32                  `protobuf:"varint,1,req,name=TurnNo" json:"TurnNo,omitempty"`
	GiveUpUsers    []int32                 `protobuf:"varint,2,rep,name=GiveUpUsers" json:"GiveUpUsers,omitempty"`
	GameRet        *WZP_GameRet            `protobuf:"bytes,3,opt,name=GameRet" json:"GameRet,omitempty"`
	UserPoker      []*UserPoker            `protobuf:"bytes,4,rep,name=UserPoker" json:"UserPoker,omitempty"`
	TotalBetAmount *int64                  `protobuf:"varint,5,req,name=TotalBetAmount" json:"TotalBetAmount,omitempty"`
	GameTurnUser   *WZP_NotifyUserGameTrun `protobuf:"bytes,6,opt,name=GameTurnUser" json:"GameTurnUser,omitempty"`
	BetLogs        []int64                 `protobuf:"varint,7,rep,name=BetLogs" json:"BetLogs,omitempty"`
}

func (x *WZP_GameScene) Reset() {
	*x = WZP_GameScene{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameMsg_WZP_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WZP_GameScene) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WZP_GameScene) ProtoMessage() {}

func (x *WZP_GameScene) ProtoReflect() protoreflect.Message {
	mi := &file_GameMsg_WZP_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WZP_GameScene.ProtoReflect.Descriptor instead.
func (*WZP_GameScene) Descriptor() ([]byte, []int) {
	return file_GameMsg_WZP_proto_rawDescGZIP(), []int{8}
}

func (x *WZP_GameScene) GetTurnNo() int32 {
	if x != nil && x.TurnNo != nil {
		return *x.TurnNo
	}
	return 0
}

func (x *WZP_GameScene) GetGiveUpUsers() []int32 {
	if x != nil {
		return x.GiveUpUsers
	}
	return nil
}

func (x *WZP_GameScene) GetGameRet() *WZP_GameRet {
	if x != nil {
		return x.GameRet
	}
	return nil
}

func (x *WZP_GameScene) GetUserPoker() []*UserPoker {
	if x != nil {
		return x.UserPoker
	}
	return nil
}

func (x *WZP_GameScene) GetTotalBetAmount() int64 {
	if x != nil && x.TotalBetAmount != nil {
		return *x.TotalBetAmount
	}
	return 0
}

func (x *WZP_GameScene) GetGameTurnUser() *WZP_NotifyUserGameTrun {
	if x != nil {
		return x.GameTurnUser
	}
	return nil
}

func (x *WZP_GameScene) GetBetLogs() []int64 {
	if x != nil {
		return x.BetLogs
	}
	return nil
}

type WZP_CardType struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TypeID *int32 `protobuf:"varint,1,req,name=TypeID" json:"TypeID,omitempty"`
}

func (x *WZP_CardType) Reset() {
	*x = WZP_CardType{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameMsg_WZP_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WZP_CardType) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WZP_CardType) ProtoMessage() {}

func (x *WZP_CardType) ProtoReflect() protoreflect.Message {
	mi := &file_GameMsg_WZP_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WZP_CardType.ProtoReflect.Descriptor instead.
func (*WZP_CardType) Descriptor() ([]byte, []int) {
	return file_GameMsg_WZP_proto_rawDescGZIP(), []int{9}
}

func (x *WZP_CardType) GetTypeID() int32 {
	if x != nil && x.TypeID != nil {
		return *x.TypeID
	}
	return 0
}

var File_GameMsg_WZP_proto protoreflect.FileDescriptor

var file_GameMsg_WZP_proto_rawDesc = []byte{
	0x0a, 0x11, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x73, 0x67, 0x5f, 0x57, 0x5a, 0x50, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x47,
	0x61, 0x6d, 0x65, 0x4d, 0x73, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x92, 0x01, 0x0a,
	0x0e, 0x57, 0x5a, 0x50, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x62, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x02, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12,
	0x1c, 0x0a, 0x09, 0x62, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x05, 0x52, 0x09, 0x62, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a,
	0x0c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20,
	0x02, 0x28, 0x05, 0x52, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x79, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75,
	0x74, 0x22, 0xd4, 0x01, 0x0a, 0x16, 0x57, 0x5a, 0x50, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x72, 0x75, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73,
	0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x72, 0x75, 0x6e, 0x4e, 0x6f, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x54, 0x72, 0x75, 0x6e, 0x4e, 0x6f, 0x12, 0x18, 0x0a, 0x07,
	0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x07, 0x54,
	0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x46, 0x69, 0x72, 0x73,
	0x74, 0x18, 0x04, 0x20, 0x02, 0x28, 0x05, 0x52, 0x07, 0x49, 0x73, 0x46, 0x69, 0x72, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x49, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x18, 0x05, 0x20, 0x02, 0x28, 0x05,
	0x52, 0x06, 0x49, 0x73, 0x4c, 0x61, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x49, 0x73, 0x53, 0x75,
	0x6f, 0x68, 0x61, 0x18, 0x06, 0x20, 0x02, 0x28, 0x05, 0x52, 0x07, 0x49, 0x73, 0x53, 0x75, 0x6f,
	0x68, 0x61, 0x12, 0x24, 0x0a, 0x0d, 0x54, 0x75, 0x72, 0x6e, 0x42, 0x65, 0x74, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0d, 0x54, 0x75, 0x72, 0x6e, 0x42,
	0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x4e, 0x0a, 0x12, 0x57, 0x5a, 0x50, 0x5f,
	0x55, 0x73, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05,
	0x52, 0x08, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x65,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52, 0x09, 0x42,
	0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x84, 0x02, 0x0a, 0x18, 0x57, 0x5a, 0x50,
	0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79, 0x55, 0x73, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x1a, 0x0a,
	0x08, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52,
	0x08, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x42, 0x65, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x02, 0x28, 0x03, 0x52, 0x09, 0x42, 0x65,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c,
	0x42, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x02, 0x28, 0x03, 0x52,
	0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x49, 0x73, 0x53, 0x75, 0x6f, 0x68, 0x61, 0x18, 0x05, 0x20, 0x02, 0x28, 0x05,
	0x52, 0x07, 0x49, 0x73, 0x53, 0x75, 0x6f, 0x68, 0x61, 0x12, 0x2e, 0x0a, 0x08, 0x47, 0x61, 0x6d,
	0x65, 0x55, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6e, 0x65,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x08, 0x47, 0x61, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0d, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x02, 0x28, 0x03,
	0x52, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x42, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22,
	0x58, 0x0a, 0x0e, 0x57, 0x5a, 0x50, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x42, 0x65, 0x74, 0x4c, 0x6f,
	0x67, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20, 0x02, 0x28,
	0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x03, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x72, 0x75, 0x6e, 0x4e, 0x6f, 0x18, 0x03, 0x20, 0x02, 0x28,
	0x05, 0x52, 0x06, 0x54, 0x72, 0x75, 0x6e, 0x4e, 0x6f, 0x22, 0x6c, 0x0a, 0x0e, 0x57, 0x5a, 0x50,
	0x5f, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x12, 0x32, 0x0a, 0x07, 0x42,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6e,
	0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x5a, 0x50, 0x5f, 0x55, 0x73, 0x65, 0x72,
	0x42, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x52, 0x07, 0x42, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x12,
	0x26, 0x0a, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x65,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xc4, 0x01, 0x0a, 0x0b, 0x57, 0x5a, 0x50, 0x5f,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x12,
	0x16, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x6f, 0x18, 0x02, 0x20, 0x02, 0x28, 0x05, 0x52,
	0x06, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x61, 0x72, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x43, 0x61, 0x72, 0x64, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x57, 0x69, 0x6e, 0x4c, 0x6f, 0x73, 0x65, 0x18, 0x04,
	0x20, 0x02, 0x28, 0x03, 0x52, 0x07, 0x57, 0x69, 0x6e, 0x4c, 0x6f, 0x73, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x42, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x02, 0x28, 0x03,
	0x52, 0x09, 0x42, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x55,
	0x73, 0x65, 0x72, 0x50, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f,
	0x6b, 0x65, 0x72, 0x52, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x6b, 0x65, 0x72, 0x22, 0x9c,
	0x01, 0x0a, 0x0b, 0x57, 0x5a, 0x50, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x74, 0x12, 0x24,
	0x0a, 0x0d, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0d, 0x47, 0x61, 0x6d, 0x65, 0x42, 0x65, 0x74, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x57, 0x5a, 0x50, 0x5f, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x74, 0x52, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x74, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x47, 0x69, 0x76, 0x65, 0x55,
	0x70, 0x18, 0x03, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x47, 0x69, 0x76, 0x65, 0x55, 0x70, 0x12,
	0x1c, 0x0a, 0x09, 0x57, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x18, 0x04, 0x20, 0x02,
	0x28, 0x05, 0x52, 0x09, 0x57, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x44, 0x22, 0xb5, 0x02,
	0x0a, 0x0d, 0x57, 0x5a, 0x50, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x54, 0x75, 0x72, 0x6e, 0x4e, 0x6f, 0x18, 0x01, 0x20, 0x02, 0x28, 0x05, 0x52,
	0x06, 0x54, 0x75, 0x72, 0x6e, 0x4e, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x47, 0x69, 0x76, 0x65, 0x55,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x47, 0x69,
	0x76, 0x65, 0x55, 0x70, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x2f, 0x0a, 0x07, 0x47, 0x61, 0x6d,
	0x65, 0x52, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6e, 0x65, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x5a, 0x50, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65,
	0x74, 0x52, 0x07, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x74, 0x12, 0x31, 0x0a, 0x09, 0x55, 0x73,
	0x65, 0x72, 0x50, 0x6f, 0x6b, 0x65, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x6b,
	0x65, 0x72, 0x52, 0x09, 0x55, 0x73, 0x65, 0x72, 0x50, 0x6f, 0x6b, 0x65, 0x72, 0x12, 0x26, 0x0a,
	0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x65, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x02, 0x28, 0x03, 0x52, 0x0e, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x65, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x44, 0x0a, 0x0c, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x75, 0x72,
	0x6e, 0x55, 0x73, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x6e, 0x65,
	0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x57, 0x5a, 0x50, 0x5f, 0x4e, 0x6f, 0x74, 0x69, 0x66,
	0x79, 0x55, 0x73, 0x65, 0x72, 0x47, 0x61, 0x6d, 0x65, 0x54, 0x72, 0x75, 0x6e, 0x52, 0x0c, 0x47,
	0x61, 0x6d, 0x65, 0x54, 0x75, 0x72, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x42,
	0x65, 0x74, 0x4c, 0x6f, 0x67, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07, 0x42, 0x65,
	0x74, 0x4c, 0x6f, 0x67, 0x73, 0x22, 0x26, 0x0a, 0x0c, 0x57, 0x5a, 0x50, 0x5f, 0x43, 0x61, 0x72,
	0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x05, 0x52, 0x06, 0x54, 0x79, 0x70, 0x65, 0x49, 0x44, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_GameMsg_WZP_proto_rawDescOnce sync.Once
	file_GameMsg_WZP_proto_rawDescData = file_GameMsg_WZP_proto_rawDesc
)

func file_GameMsg_WZP_proto_rawDescGZIP() []byte {
	file_GameMsg_WZP_proto_rawDescOnce.Do(func() {
		file_GameMsg_WZP_proto_rawDescData = protoimpl.X.CompressGZIP(file_GameMsg_WZP_proto_rawDescData)
	})
	return file_GameMsg_WZP_proto_rawDescData
}

var file_GameMsg_WZP_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_GameMsg_WZP_proto_goTypes = []interface{}{
	(*WZP_GameConfig)(nil),           // 0: netproto.WZP_GameConfig
	(*WZP_NotifyUserGameTrun)(nil),   // 1: netproto.WZP_NotifyUserGameTrun
	(*WZP_UserGameAction)(nil),       // 2: netproto.WZP_UserGameAction
	(*WZP_NotifyUserGameAction)(nil), // 3: netproto.WZP_NotifyUserGameAction
	(*WZP_UserBetLog)(nil),           // 4: netproto.WZP_UserBetLog
	(*WZP_GameBetLog)(nil),           // 5: netproto.WZP_GameBetLog
	(*WZP_UserRet)(nil),              // 6: netproto.WZP_UserRet
	(*WZP_GameRet)(nil),              // 7: netproto.WZP_GameRet
	(*WZP_GameScene)(nil),            // 8: netproto.WZP_GameScene
	(*WZP_CardType)(nil),             // 9: netproto.WZP_CardType
	(*GameUser)(nil),                 // 10: netproto.GameUser
	(*UserPoker)(nil),                // 11: netproto.UserPoker
}
var file_GameMsg_WZP_proto_depIdxs = []int32{
	10, // 0: netproto.WZP_NotifyUserGameAction.GameUser:type_name -> netproto.GameUser
	4,  // 1: netproto.WZP_GameBetLog.BetLogs:type_name -> netproto.WZP_UserBetLog
	11, // 2: netproto.WZP_UserRet.UserPoker:type_name -> netproto.UserPoker
	6,  // 3: netproto.WZP_GameRet.UserRets:type_name -> netproto.WZP_UserRet
	7,  // 4: netproto.WZP_GameScene.GameRet:type_name -> netproto.WZP_GameRet
	11, // 5: netproto.WZP_GameScene.UserPoker:type_name -> netproto.UserPoker
	1,  // 6: netproto.WZP_GameScene.GameTurnUser:type_name -> netproto.WZP_NotifyUserGameTrun
	7,  // [7:7] is the sub-list for method output_type
	7,  // [7:7] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_GameMsg_WZP_proto_init() }
func file_GameMsg_WZP_proto_init() {
	if File_GameMsg_WZP_proto != nil {
		return
	}
	file_GameMsg_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_GameMsg_WZP_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WZP_GameConfig); i {
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
		file_GameMsg_WZP_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WZP_NotifyUserGameTrun); i {
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
		file_GameMsg_WZP_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WZP_UserGameAction); i {
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
		file_GameMsg_WZP_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WZP_NotifyUserGameAction); i {
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
		file_GameMsg_WZP_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WZP_UserBetLog); i {
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
		file_GameMsg_WZP_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WZP_GameBetLog); i {
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
		file_GameMsg_WZP_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WZP_UserRet); i {
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
		file_GameMsg_WZP_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WZP_GameRet); i {
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
		file_GameMsg_WZP_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WZP_GameScene); i {
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
		file_GameMsg_WZP_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WZP_CardType); i {
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
			RawDescriptor: file_GameMsg_WZP_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GameMsg_WZP_proto_goTypes,
		DependencyIndexes: file_GameMsg_WZP_proto_depIdxs,
		MessageInfos:      file_GameMsg_WZP_proto_msgTypes,
	}.Build()
	File_GameMsg_WZP_proto = out.File
	file_GameMsg_WZP_proto_rawDesc = nil
	file_GameMsg_WZP_proto_goTypes = nil
	file_GameMsg_WZP_proto_depIdxs = nil
}