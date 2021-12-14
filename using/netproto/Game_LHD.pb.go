// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: Game_LHD.proto

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

type LHD_GameMessageClassID int32

const (
	LHD_GameMessageClassID_LHDSettingID             LHD_GameMessageClassID = 10001 //设置相关 S2C
	LHD_GameMessageClassID_LHDStateID               LHD_GameMessageClassID = 10002 //阶段ID S2C
	LHD_GameMessageClassID_LHDBetID                 LHD_GameMessageClassID = 10003 //下注ID C2S
	LHD_GameMessageClassID_LHDBroadCastBetID        LHD_GameMessageClassID = 10004 //广播下注 S2C
	LHD_GameMessageClassID_LHDGameResultPokerID     LHD_GameMessageClassID = 10005 //游戏结果扑克 S2C
	LHD_GameMessageClassID_LHDGameSceneID           LHD_GameMessageClassID = 10006 //游戏场景 断线重连
	LHD_GameMessageClassID_LHDGameResultWinID       LHD_GameMessageClassID = 10007 //游戏结果赢钱 S2C
	LHD_GameMessageClassID_LHDBetInfoID             LHD_GameMessageClassID = 10008 //用户断线重连已下注的消息
	LHD_GameMessageClassID_LHDBetRetID              LHD_GameMessageClassID = 10009 //下注结果 S2C
	LHD_GameMessageClassID_LHDShangZhuangID         LHD_GameMessageClassID = 10010 //上庄请求 C2S
	LHD_GameMessageClassID_LHDShangZhuangRetID      LHD_GameMessageClassID = 10011 //上庄请求结果 S2C
	LHD_GameMessageClassID_LHDGetShangZhuangListID  LHD_GameMessageClassID = 10012 //获取上庄列表 C2S
	LHD_GameMessageClassID_LHDShangZhuangListInfoID LHD_GameMessageClassID = 10013 //下发上庄列表 S2C
	LHD_GameMessageClassID_LHDZhuangInfoID          LHD_GameMessageClassID = 10014 //庄家信息 S2C
	LHD_GameMessageClassID_LHDBetClearID            LHD_GameMessageClassID = 10015 //玩家清除已下的注 C2S
	LHD_GameMessageClassID_LHDResultCardListID      LHD_GameMessageClassID = 10016 //前20次开牌结果
	LHD_GameMessageClassID_LHDBetClearRetID         LHD_GameMessageClassID = 10017 //清理下注结果
	LHD_GameMessageClassID_LHDRichestListID         LHD_GameMessageClassID = 10018 //大富豪列表的消息
	LHD_GameMessageClassID_LHDLuckyStarListID       LHD_GameMessageClassID = 10019 //明燈玩家的消息
	LHD_GameMessageClassID_LHDRoundID               LHD_GameMessageClassID = 10020 //龍虎鬥本局局數
	LHD_GameMessageClassID_LHDGameNoInfo            LHD_GameMessageClassID = 10021 //牌局編號
)

// Enum value maps for LHD_GameMessageClassID.
var (
	LHD_GameMessageClassID_name = map[int32]string{
		10001: "LHDSettingID",
		10002: "LHDStateID",
		10003: "LHDBetID",
		10004: "LHDBroadCastBetID",
		10005: "LHDGameResultPokerID",
		10006: "LHDGameSceneID",
		10007: "LHDGameResultWinID",
		10008: "LHDBetInfoID",
		10009: "LHDBetRetID",
		10010: "LHDShangZhuangID",
		10011: "LHDShangZhuangRetID",
		10012: "LHDGetShangZhuangListID",
		10013: "LHDShangZhuangListInfoID",
		10014: "LHDZhuangInfoID",
		10015: "LHDBetClearID",
		10016: "LHDResultCardListID",
		10017: "LHDBetClearRetID",
		10018: "LHDRichestListID",
		10019: "LHDLuckyStarListID",
		10020: "LHDRoundID",
		10021: "LHDGameNoInfo",
	}
	LHD_GameMessageClassID_value = map[string]int32{
		"LHDSettingID":             10001,
		"LHDStateID":               10002,
		"LHDBetID":                 10003,
		"LHDBroadCastBetID":        10004,
		"LHDGameResultPokerID":     10005,
		"LHDGameSceneID":           10006,
		"LHDGameResultWinID":       10007,
		"LHDBetInfoID":             10008,
		"LHDBetRetID":              10009,
		"LHDShangZhuangID":         10010,
		"LHDShangZhuangRetID":      10011,
		"LHDGetShangZhuangListID":  10012,
		"LHDShangZhuangListInfoID": 10013,
		"LHDZhuangInfoID":          10014,
		"LHDBetClearID":            10015,
		"LHDResultCardListID":      10016,
		"LHDBetClearRetID":         10017,
		"LHDRichestListID":         10018,
		"LHDLuckyStarListID":       10019,
		"LHDRoundID":               10020,
		"LHDGameNoInfo":            10021,
	}
)

func (x LHD_GameMessageClassID) Enum() *LHD_GameMessageClassID {
	p := new(LHD_GameMessageClassID)
	*p = x
	return p
}

func (x LHD_GameMessageClassID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LHD_GameMessageClassID) Descriptor() protoreflect.EnumDescriptor {
	return file_Game_LHD_proto_enumTypes[0].Descriptor()
}

func (LHD_GameMessageClassID) Type() protoreflect.EnumType {
	return &file_Game_LHD_proto_enumTypes[0]
}

func (x LHD_GameMessageClassID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LHD_GameMessageClassID) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LHD_GameMessageClassID(num)
	return nil
}

// Deprecated: Use LHD_GameMessageClassID.Descriptor instead.
func (LHD_GameMessageClassID) EnumDescriptor() ([]byte, []int) {
	return file_Game_LHD_proto_rawDescGZIP(), []int{0}
}

var File_Game_LHD_proto protoreflect.FileDescriptor

var file_Game_LHD_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x5f, 0x4c, 0x48, 0x44, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xef, 0x03, 0x0a, 0x16, 0x4c,
	0x48, 0x44, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x11, 0x0a, 0x0c, 0x4c, 0x48, 0x44, 0x53, 0x65, 0x74, 0x74,
	0x69, 0x6e, 0x67, 0x49, 0x44, 0x10, 0x91, 0x4e, 0x12, 0x0f, 0x0a, 0x0a, 0x4c, 0x48, 0x44, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x49, 0x44, 0x10, 0x92, 0x4e, 0x12, 0x0d, 0x0a, 0x08, 0x4c, 0x48, 0x44,
	0x42, 0x65, 0x74, 0x49, 0x44, 0x10, 0x93, 0x4e, 0x12, 0x16, 0x0a, 0x11, 0x4c, 0x48, 0x44, 0x42,
	0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x42, 0x65, 0x74, 0x49, 0x44, 0x10, 0x94, 0x4e,
	0x12, 0x19, 0x0a, 0x14, 0x4c, 0x48, 0x44, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x50, 0x6f, 0x6b, 0x65, 0x72, 0x49, 0x44, 0x10, 0x95, 0x4e, 0x12, 0x13, 0x0a, 0x0e, 0x4c,
	0x48, 0x44, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x44, 0x10, 0x96, 0x4e,
	0x12, 0x17, 0x0a, 0x12, 0x4c, 0x48, 0x44, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x57, 0x69, 0x6e, 0x49, 0x44, 0x10, 0x97, 0x4e, 0x12, 0x11, 0x0a, 0x0c, 0x4c, 0x48, 0x44,
	0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x10, 0x98, 0x4e, 0x12, 0x10, 0x0a, 0x0b,
	0x4c, 0x48, 0x44, 0x42, 0x65, 0x74, 0x52, 0x65, 0x74, 0x49, 0x44, 0x10, 0x99, 0x4e, 0x12, 0x15,
	0x0a, 0x10, 0x4c, 0x48, 0x44, 0x53, 0x68, 0x61, 0x6e, 0x67, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67,
	0x49, 0x44, 0x10, 0x9a, 0x4e, 0x12, 0x18, 0x0a, 0x13, 0x4c, 0x48, 0x44, 0x53, 0x68, 0x61, 0x6e,
	0x67, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x74, 0x49, 0x44, 0x10, 0x9b, 0x4e, 0x12,
	0x1c, 0x0a, 0x17, 0x4c, 0x48, 0x44, 0x47, 0x65, 0x74, 0x53, 0x68, 0x61, 0x6e, 0x67, 0x5a, 0x68,
	0x75, 0x61, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x10, 0x9c, 0x4e, 0x12, 0x1d, 0x0a,
	0x18, 0x4c, 0x48, 0x44, 0x53, 0x68, 0x61, 0x6e, 0x67, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x10, 0x9d, 0x4e, 0x12, 0x14, 0x0a, 0x0f,
	0x4c, 0x48, 0x44, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x10,
	0x9e, 0x4e, 0x12, 0x12, 0x0a, 0x0d, 0x4c, 0x48, 0x44, 0x42, 0x65, 0x74, 0x43, 0x6c, 0x65, 0x61,
	0x72, 0x49, 0x44, 0x10, 0x9f, 0x4e, 0x12, 0x18, 0x0a, 0x13, 0x4c, 0x48, 0x44, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x43, 0x61, 0x72, 0x64, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x10, 0xa0, 0x4e,
	0x12, 0x15, 0x0a, 0x10, 0x4c, 0x48, 0x44, 0x42, 0x65, 0x74, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52,
	0x65, 0x74, 0x49, 0x44, 0x10, 0xa1, 0x4e, 0x12, 0x15, 0x0a, 0x10, 0x4c, 0x48, 0x44, 0x52, 0x69,
	0x63, 0x68, 0x65, 0x73, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x49, 0x44, 0x10, 0xa2, 0x4e, 0x12, 0x17,
	0x0a, 0x12, 0x4c, 0x48, 0x44, 0x4c, 0x75, 0x63, 0x6b, 0x79, 0x53, 0x74, 0x61, 0x72, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x44, 0x10, 0xa3, 0x4e, 0x12, 0x0f, 0x0a, 0x0a, 0x4c, 0x48, 0x44, 0x52, 0x6f,
	0x75, 0x6e, 0x64, 0x49, 0x44, 0x10, 0xa4, 0x4e, 0x12, 0x12, 0x0a, 0x0d, 0x4c, 0x48, 0x44, 0x47,
	0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x10, 0xa5, 0x4e, 0x42, 0x0c, 0x5a, 0x0a,
	0x2e, 0x2f, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_Game_LHD_proto_rawDescOnce sync.Once
	file_Game_LHD_proto_rawDescData = file_Game_LHD_proto_rawDesc
)

func file_Game_LHD_proto_rawDescGZIP() []byte {
	file_Game_LHD_proto_rawDescOnce.Do(func() {
		file_Game_LHD_proto_rawDescData = protoimpl.X.CompressGZIP(file_Game_LHD_proto_rawDescData)
	})
	return file_Game_LHD_proto_rawDescData
}

var file_Game_LHD_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Game_LHD_proto_goTypes = []interface{}{
	(LHD_GameMessageClassID)(0), // 0: netproto.LHD_GameMessageClassID
}
var file_Game_LHD_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Game_LHD_proto_init() }
func file_Game_LHD_proto_init() {
	if File_Game_LHD_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Game_LHD_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Game_LHD_proto_goTypes,
		DependencyIndexes: file_Game_LHD_proto_depIdxs,
		EnumInfos:         file_Game_LHD_proto_enumTypes,
	}.Build()
	File_Game_LHD_proto = out.File
	file_Game_LHD_proto_rawDesc = nil
	file_Game_LHD_proto_goTypes = nil
	file_Game_LHD_proto_depIdxs = nil
}
