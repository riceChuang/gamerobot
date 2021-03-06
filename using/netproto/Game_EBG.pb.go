// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: Game_EBG.proto

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

type EBG_GameMessageClassID int32

const (
	EBG_GameMessageClassID_EBGStateID         EBG_GameMessageClassID = 10001 //阶段变化ID        S2C
	EBG_GameMessageClassID_EBGSettingID       EBG_GameMessageClassID = 10002 //设置相关          S2C
	EBG_GameMessageClassID_EBGGameInfoID      EBG_GameMessageClassID = 10003 //该局游戏信息      S2C
	EBG_GameMessageClassID_EBGBetConfigID     EBG_GameMessageClassID = 10004 //抢庄和下注倍数配置  S2C
	EBG_GameMessageClassID_EBGBCVieBankerID   EBG_GameMessageClassID = 10005 //广播用户抢庄       S2C
	EBG_GameMessageClassID_EBGVieBankerID     EBG_GameMessageClassID = 10006 //用户抢庄          C2S
	EBG_GameMessageClassID_EBGConfirmBankerID EBG_GameMessageClassID = 10007 //确定庄家          S2C
	EBG_GameMessageClassID_EBGBetInfoID       EBG_GameMessageClassID = 10008 //广播用户下注       S2C
	EBG_GameMessageClassID_EBGBetID           EBG_GameMessageClassID = 10009 //用户下注           C2S
	EBG_GameMessageClassID_EBGDiceResultID    EBG_GameMessageClassID = 10010 //掷骰子点数        S2C
	EBG_GameMessageClassID_EBGBCHandCardsID   EBG_GameMessageClassID = 10011 //广播用户牌        S2C
	EBG_GameMessageClassID_EBGGameSceneID     EBG_GameMessageClassID = 10012 //游戏场景          S2C
	EBG_GameMessageClassID_EBGGameResultID    EBG_GameMessageClassID = 10013 //游戏结果          S2C
	EBG_GameMessageClassID_EBGContinueID      EBG_GameMessageClassID = 10014 //是否接着下一局     S2C
)

// Enum value maps for EBG_GameMessageClassID.
var (
	EBG_GameMessageClassID_name = map[int32]string{
		10001: "EBGStateID",
		10002: "EBGSettingID",
		10003: "EBGGameInfoID",
		10004: "EBGBetConfigID",
		10005: "EBGBCVieBankerID",
		10006: "EBGVieBankerID",
		10007: "EBGConfirmBankerID",
		10008: "EBGBetInfoID",
		10009: "EBGBetID",
		10010: "EBGDiceResultID",
		10011: "EBGBCHandCardsID",
		10012: "EBGGameSceneID",
		10013: "EBGGameResultID",
		10014: "EBGContinueID",
	}
	EBG_GameMessageClassID_value = map[string]int32{
		"EBGStateID":         10001,
		"EBGSettingID":       10002,
		"EBGGameInfoID":      10003,
		"EBGBetConfigID":     10004,
		"EBGBCVieBankerID":   10005,
		"EBGVieBankerID":     10006,
		"EBGConfirmBankerID": 10007,
		"EBGBetInfoID":       10008,
		"EBGBetID":           10009,
		"EBGDiceResultID":    10010,
		"EBGBCHandCardsID":   10011,
		"EBGGameSceneID":     10012,
		"EBGGameResultID":    10013,
		"EBGContinueID":      10014,
	}
)

func (x EBG_GameMessageClassID) Enum() *EBG_GameMessageClassID {
	p := new(EBG_GameMessageClassID)
	*p = x
	return p
}

func (x EBG_GameMessageClassID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EBG_GameMessageClassID) Descriptor() protoreflect.EnumDescriptor {
	return file_Game_EBG_proto_enumTypes[0].Descriptor()
}

func (EBG_GameMessageClassID) Type() protoreflect.EnumType {
	return &file_Game_EBG_proto_enumTypes[0]
}

func (x EBG_GameMessageClassID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EBG_GameMessageClassID) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EBG_GameMessageClassID(num)
	return nil
}

// Deprecated: Use EBG_GameMessageClassID.Descriptor instead.
func (EBG_GameMessageClassID) EnumDescriptor() ([]byte, []int) {
	return file_Game_EBG_proto_rawDescGZIP(), []int{0}
}

//牌型定义
type EBG_CardType int32

const (
	EBG_CardType_King    EBG_CardType = 1  //天王
	EBG_CardType_PairNin EBG_CardType = 2  //九宝
	EBG_CardType_PairEig EBG_CardType = 3  //八宝
	EBG_CardType_PairSev EBG_CardType = 4  //七宝
	EBG_CardType_PairSix EBG_CardType = 5  //六宝
	EBG_CardType_PairFiv EBG_CardType = 6  //五宝
	EBG_CardType_PairFou EBG_CardType = 7  //四宝
	EBG_CardType_PairThr EBG_CardType = 8  //三宝
	EBG_CardType_PairTwo EBG_CardType = 9  //二宝
	EBG_CardType_PairOne EBG_CardType = 10 //一宝
	EBG_CardType_HalfNin EBG_CardType = 11 //九点半
	EBG_CardType_Nine    EBG_CardType = 12 //九点
	EBG_CardType_HalfEig EBG_CardType = 13 //八点半
	EBG_CardType_Eight   EBG_CardType = 14 //八点
	EBG_CardType_HalfSev EBG_CardType = 15 //七点半
	EBG_CardType_Seven   EBG_CardType = 16 //七点
	EBG_CardType_HalfSix EBG_CardType = 17 //六点半
	EBG_CardType_Six     EBG_CardType = 18 //六点
	EBG_CardType_HalfFiv EBG_CardType = 19 //五点半
	EBG_CardType_Five    EBG_CardType = 20 //五点
	EBG_CardType_HalfFou EBG_CardType = 21 //四点半
	EBG_CardType_Four    EBG_CardType = 22 //四点
	EBG_CardType_HalfThr EBG_CardType = 23 //三点半
	EBG_CardType_Three   EBG_CardType = 24 //三点
	EBG_CardType_HalfTwo EBG_CardType = 25 //二点半
	EBG_CardType_Two     EBG_CardType = 26 //二点
	EBG_CardType_HalfOne EBG_CardType = 27 //一点半
	EBG_CardType_One     EBG_CardType = 28 //一点
	EBG_CardType_Ten     EBG_CardType = 29 //憋十
)

// Enum value maps for EBG_CardType.
var (
	EBG_CardType_name = map[int32]string{
		1:  "King",
		2:  "PairNin",
		3:  "PairEig",
		4:  "PairSev",
		5:  "PairSix",
		6:  "PairFiv",
		7:  "PairFou",
		8:  "PairThr",
		9:  "PairTwo",
		10: "PairOne",
		11: "HalfNin",
		12: "Nine",
		13: "HalfEig",
		14: "Eight",
		15: "HalfSev",
		16: "Seven",
		17: "HalfSix",
		18: "Six",
		19: "HalfFiv",
		20: "Five",
		21: "HalfFou",
		22: "Four",
		23: "HalfThr",
		24: "Three",
		25: "HalfTwo",
		26: "Two",
		27: "HalfOne",
		28: "One",
		29: "Ten",
	}
	EBG_CardType_value = map[string]int32{
		"King":    1,
		"PairNin": 2,
		"PairEig": 3,
		"PairSev": 4,
		"PairSix": 5,
		"PairFiv": 6,
		"PairFou": 7,
		"PairThr": 8,
		"PairTwo": 9,
		"PairOne": 10,
		"HalfNin": 11,
		"Nine":    12,
		"HalfEig": 13,
		"Eight":   14,
		"HalfSev": 15,
		"Seven":   16,
		"HalfSix": 17,
		"Six":     18,
		"HalfFiv": 19,
		"Five":    20,
		"HalfFou": 21,
		"Four":    22,
		"HalfThr": 23,
		"Three":   24,
		"HalfTwo": 25,
		"Two":     26,
		"HalfOne": 27,
		"One":     28,
		"Ten":     29,
	}
)

func (x EBG_CardType) Enum() *EBG_CardType {
	p := new(EBG_CardType)
	*p = x
	return p
}

func (x EBG_CardType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EBG_CardType) Descriptor() protoreflect.EnumDescriptor {
	return file_Game_EBG_proto_enumTypes[1].Descriptor()
}

func (EBG_CardType) Type() protoreflect.EnumType {
	return &file_Game_EBG_proto_enumTypes[1]
}

func (x EBG_CardType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *EBG_CardType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = EBG_CardType(num)
	return nil
}

// Deprecated: Use EBG_CardType.Descriptor instead.
func (EBG_CardType) EnumDescriptor() ([]byte, []int) {
	return file_Game_EBG_proto_rawDescGZIP(), []int{1}
}

var File_Game_EBG_proto protoreflect.FileDescriptor

var file_Game_EBG_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x5f, 0x45, 0x42, 0x47, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xb8, 0x02, 0x0a, 0x16, 0x45,
	0x42, 0x47, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x0f, 0x0a, 0x0a, 0x45, 0x42, 0x47, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x49, 0x44, 0x10, 0x91, 0x4e, 0x12, 0x11, 0x0a, 0x0c, 0x45, 0x42, 0x47, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x10, 0x92, 0x4e, 0x12, 0x12, 0x0a, 0x0d, 0x45, 0x42, 0x47,
	0x47, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x10, 0x93, 0x4e, 0x12, 0x13, 0x0a,
	0x0e, 0x45, 0x42, 0x47, 0x42, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x44, 0x10,
	0x94, 0x4e, 0x12, 0x15, 0x0a, 0x10, 0x45, 0x42, 0x47, 0x42, 0x43, 0x56, 0x69, 0x65, 0x42, 0x61,
	0x6e, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x10, 0x95, 0x4e, 0x12, 0x13, 0x0a, 0x0e, 0x45, 0x42, 0x47,
	0x56, 0x69, 0x65, 0x42, 0x61, 0x6e, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x10, 0x96, 0x4e, 0x12, 0x17,
	0x0a, 0x12, 0x45, 0x42, 0x47, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x72, 0x6d, 0x42, 0x61, 0x6e, 0x6b,
	0x65, 0x72, 0x49, 0x44, 0x10, 0x97, 0x4e, 0x12, 0x11, 0x0a, 0x0c, 0x45, 0x42, 0x47, 0x42, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x10, 0x98, 0x4e, 0x12, 0x0d, 0x0a, 0x08, 0x45, 0x42,
	0x47, 0x42, 0x65, 0x74, 0x49, 0x44, 0x10, 0x99, 0x4e, 0x12, 0x14, 0x0a, 0x0f, 0x45, 0x42, 0x47,
	0x44, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x44, 0x10, 0x9a, 0x4e, 0x12,
	0x15, 0x0a, 0x10, 0x45, 0x42, 0x47, 0x42, 0x43, 0x48, 0x61, 0x6e, 0x64, 0x43, 0x61, 0x72, 0x64,
	0x73, 0x49, 0x44, 0x10, 0x9b, 0x4e, 0x12, 0x13, 0x0a, 0x0e, 0x45, 0x42, 0x47, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x44, 0x10, 0x9c, 0x4e, 0x12, 0x14, 0x0a, 0x0f, 0x45,
	0x42, 0x47, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x44, 0x10, 0x9d,
	0x4e, 0x12, 0x12, 0x0a, 0x0d, 0x45, 0x42, 0x47, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65,
	0x49, 0x44, 0x10, 0x9e, 0x4e, 0x2a, 0xe5, 0x02, 0x0a, 0x0c, 0x45, 0x42, 0x47, 0x5f, 0x43, 0x61,
	0x72, 0x64, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x67, 0x10, 0x01,
	0x12, 0x0b, 0x0a, 0x07, 0x50, 0x61, 0x69, 0x72, 0x4e, 0x69, 0x6e, 0x10, 0x02, 0x12, 0x0b, 0x0a,
	0x07, 0x50, 0x61, 0x69, 0x72, 0x45, 0x69, 0x67, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x61,
	0x69, 0x72, 0x53, 0x65, 0x76, 0x10, 0x04, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x61, 0x69, 0x72, 0x53,
	0x69, 0x78, 0x10, 0x05, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x61, 0x69, 0x72, 0x46, 0x69, 0x76, 0x10,
	0x06, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x61, 0x69, 0x72, 0x46, 0x6f, 0x75, 0x10, 0x07, 0x12, 0x0b,
	0x0a, 0x07, 0x50, 0x61, 0x69, 0x72, 0x54, 0x68, 0x72, 0x10, 0x08, 0x12, 0x0b, 0x0a, 0x07, 0x50,
	0x61, 0x69, 0x72, 0x54, 0x77, 0x6f, 0x10, 0x09, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x61, 0x69, 0x72,
	0x4f, 0x6e, 0x65, 0x10, 0x0a, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x61, 0x6c, 0x66, 0x4e, 0x69, 0x6e,
	0x10, 0x0b, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x69, 0x6e, 0x65, 0x10, 0x0c, 0x12, 0x0b, 0x0a, 0x07,
	0x48, 0x61, 0x6c, 0x66, 0x45, 0x69, 0x67, 0x10, 0x0d, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x69, 0x67,
	0x68, 0x74, 0x10, 0x0e, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x61, 0x6c, 0x66, 0x53, 0x65, 0x76, 0x10,
	0x0f, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x65, 0x76, 0x65, 0x6e, 0x10, 0x10, 0x12, 0x0b, 0x0a, 0x07,
	0x48, 0x61, 0x6c, 0x66, 0x53, 0x69, 0x78, 0x10, 0x11, 0x12, 0x07, 0x0a, 0x03, 0x53, 0x69, 0x78,
	0x10, 0x12, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x61, 0x6c, 0x66, 0x46, 0x69, 0x76, 0x10, 0x13, 0x12,
	0x08, 0x0a, 0x04, 0x46, 0x69, 0x76, 0x65, 0x10, 0x14, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x61, 0x6c,
	0x66, 0x46, 0x6f, 0x75, 0x10, 0x15, 0x12, 0x08, 0x0a, 0x04, 0x46, 0x6f, 0x75, 0x72, 0x10, 0x16,
	0x12, 0x0b, 0x0a, 0x07, 0x48, 0x61, 0x6c, 0x66, 0x54, 0x68, 0x72, 0x10, 0x17, 0x12, 0x09, 0x0a,
	0x05, 0x54, 0x68, 0x72, 0x65, 0x65, 0x10, 0x18, 0x12, 0x0b, 0x0a, 0x07, 0x48, 0x61, 0x6c, 0x66,
	0x54, 0x77, 0x6f, 0x10, 0x19, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x77, 0x6f, 0x10, 0x1a, 0x12, 0x0b,
	0x0a, 0x07, 0x48, 0x61, 0x6c, 0x66, 0x4f, 0x6e, 0x65, 0x10, 0x1b, 0x12, 0x07, 0x0a, 0x03, 0x4f,
	0x6e, 0x65, 0x10, 0x1c, 0x12, 0x07, 0x0a, 0x03, 0x54, 0x65, 0x6e, 0x10, 0x1d, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_Game_EBG_proto_rawDescOnce sync.Once
	file_Game_EBG_proto_rawDescData = file_Game_EBG_proto_rawDesc
)

func file_Game_EBG_proto_rawDescGZIP() []byte {
	file_Game_EBG_proto_rawDescOnce.Do(func() {
		file_Game_EBG_proto_rawDescData = protoimpl.X.CompressGZIP(file_Game_EBG_proto_rawDescData)
	})
	return file_Game_EBG_proto_rawDescData
}

var file_Game_EBG_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_Game_EBG_proto_goTypes = []interface{}{
	(EBG_GameMessageClassID)(0), // 0: netproto.EBG_GameMessageClassID
	(EBG_CardType)(0),           // 1: netproto.EBG_CardType
}
var file_Game_EBG_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Game_EBG_proto_init() }
func file_Game_EBG_proto_init() {
	if File_Game_EBG_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Game_EBG_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Game_EBG_proto_goTypes,
		DependencyIndexes: file_Game_EBG_proto_depIdxs,
		EnumInfos:         file_Game_EBG_proto_enumTypes,
	}.Build()
	File_Game_EBG_proto = out.File
	file_Game_EBG_proto_rawDesc = nil
	file_Game_EBG_proto_goTypes = nil
	file_Game_EBG_proto_depIdxs = nil
}
