// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: Game_Shaibao.proto

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

type Shaibao_GameMessageClassID int32

const (
	Shaibao_GameMessageClassID_ShaibaoSettingID             Shaibao_GameMessageClassID = 10001 //设置相关 S2C
	Shaibao_GameMessageClassID_ShaibaoStateID               Shaibao_GameMessageClassID = 10002 //阶段ID S2C
	Shaibao_GameMessageClassID_ShaibaoBetID                 Shaibao_GameMessageClassID = 10003 //下注ID C2S
	Shaibao_GameMessageClassID_ShaibaoBroadCastBetID        Shaibao_GameMessageClassID = 10004 //广播下注 S2C
	Shaibao_GameMessageClassID_ShaibaoGameResultPokerID     Shaibao_GameMessageClassID = 10005 //游戏结果扑克 S2C
	Shaibao_GameMessageClassID_ShaibaoGameSceneID           Shaibao_GameMessageClassID = 10006 //游戏场景 断线重连
	Shaibao_GameMessageClassID_ShaibaoGameResultWinID       Shaibao_GameMessageClassID = 10007 //游戏结果赢钱 S2C
	Shaibao_GameMessageClassID_ShaibaoBetInfoID             Shaibao_GameMessageClassID = 10008 //用户断线重连已下注的消息
	Shaibao_GameMessageClassID_ShaibaoBetRetID              Shaibao_GameMessageClassID = 10009 //下注结果 S2C
	Shaibao_GameMessageClassID_ShaibaoShangZhuangID         Shaibao_GameMessageClassID = 10010 //上庄请求 C2S
	Shaibao_GameMessageClassID_ShaibaoShangZhuangRetID      Shaibao_GameMessageClassID = 10011 //上庄请求结果 S2C
	Shaibao_GameMessageClassID_ShaibaoGetShangZhuangListID  Shaibao_GameMessageClassID = 10012 //获取上庄列表 C2S
	Shaibao_GameMessageClassID_ShaibaoShangZhuangListInfoID Shaibao_GameMessageClassID = 10013 //下发上庄列表 S2C
	Shaibao_GameMessageClassID_ShaibaoZhuangInfoID          Shaibao_GameMessageClassID = 10014 //庄家信息 S2C
	Shaibao_GameMessageClassID_ShaibaoBetClearID            Shaibao_GameMessageClassID = 10015 //玩家清除已下的注 C2S
	Shaibao_GameMessageClassID_ShaibaoResultCardListID      Shaibao_GameMessageClassID = 10016 //前20次开牌结果
	Shaibao_GameMessageClassID_ShaibaoBetClearRetID         Shaibao_GameMessageClassID = 10017 //清理下注结果
)

// Enum value maps for Shaibao_GameMessageClassID.
var (
	Shaibao_GameMessageClassID_name = map[int32]string{
		10001: "ShaibaoSettingID",
		10002: "ShaibaoStateID",
		10003: "ShaibaoBetID",
		10004: "ShaibaoBroadCastBetID",
		10005: "ShaibaoGameResultPokerID",
		10006: "ShaibaoGameSceneID",
		10007: "ShaibaoGameResultWinID",
		10008: "ShaibaoBetInfoID",
		10009: "ShaibaoBetRetID",
		10010: "ShaibaoShangZhuangID",
		10011: "ShaibaoShangZhuangRetID",
		10012: "ShaibaoGetShangZhuangListID",
		10013: "ShaibaoShangZhuangListInfoID",
		10014: "ShaibaoZhuangInfoID",
		10015: "ShaibaoBetClearID",
		10016: "ShaibaoResultCardListID",
		10017: "ShaibaoBetClearRetID",
	}
	Shaibao_GameMessageClassID_value = map[string]int32{
		"ShaibaoSettingID":             10001,
		"ShaibaoStateID":               10002,
		"ShaibaoBetID":                 10003,
		"ShaibaoBroadCastBetID":        10004,
		"ShaibaoGameResultPokerID":     10005,
		"ShaibaoGameSceneID":           10006,
		"ShaibaoGameResultWinID":       10007,
		"ShaibaoBetInfoID":             10008,
		"ShaibaoBetRetID":              10009,
		"ShaibaoShangZhuangID":         10010,
		"ShaibaoShangZhuangRetID":      10011,
		"ShaibaoGetShangZhuangListID":  10012,
		"ShaibaoShangZhuangListInfoID": 10013,
		"ShaibaoZhuangInfoID":          10014,
		"ShaibaoBetClearID":            10015,
		"ShaibaoResultCardListID":      10016,
		"ShaibaoBetClearRetID":         10017,
	}
)

func (x Shaibao_GameMessageClassID) Enum() *Shaibao_GameMessageClassID {
	p := new(Shaibao_GameMessageClassID)
	*p = x
	return p
}

func (x Shaibao_GameMessageClassID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Shaibao_GameMessageClassID) Descriptor() protoreflect.EnumDescriptor {
	return file_Game_Shaibao_proto_enumTypes[0].Descriptor()
}

func (Shaibao_GameMessageClassID) Type() protoreflect.EnumType {
	return &file_Game_Shaibao_proto_enumTypes[0]
}

func (x Shaibao_GameMessageClassID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *Shaibao_GameMessageClassID) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = Shaibao_GameMessageClassID(num)
	return nil
}

// Deprecated: Use Shaibao_GameMessageClassID.Descriptor instead.
func (Shaibao_GameMessageClassID) EnumDescriptor() ([]byte, []int) {
	return file_Game_Shaibao_proto_rawDescGZIP(), []int{0}
}

var File_Game_Shaibao_proto protoreflect.FileDescriptor

var file_Game_Shaibao_proto_rawDesc = []byte{
	0x0a, 0x12, 0x47, 0x61, 0x6d, 0x65, 0x5f, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xe2,
	0x03, 0x0a, 0x1a, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x15, 0x0a,
	0x10, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49,
	0x44, 0x10, 0x91, 0x4e, 0x12, 0x13, 0x0a, 0x0e, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x49, 0x44, 0x10, 0x92, 0x4e, 0x12, 0x11, 0x0a, 0x0c, 0x53, 0x68, 0x61,
	0x69, 0x62, 0x61, 0x6f, 0x42, 0x65, 0x74, 0x49, 0x44, 0x10, 0x93, 0x4e, 0x12, 0x1a, 0x0a, 0x15,
	0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74,
	0x42, 0x65, 0x74, 0x49, 0x44, 0x10, 0x94, 0x4e, 0x12, 0x1d, 0x0a, 0x18, 0x53, 0x68, 0x61, 0x69,
	0x62, 0x61, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x6f, 0x6b,
	0x65, 0x72, 0x49, 0x44, 0x10, 0x95, 0x4e, 0x12, 0x17, 0x0a, 0x12, 0x53, 0x68, 0x61, 0x69, 0x62,
	0x61, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x44, 0x10, 0x96, 0x4e,
	0x12, 0x1b, 0x0a, 0x16, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f, 0x47, 0x61, 0x6d, 0x65, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x57, 0x69, 0x6e, 0x49, 0x44, 0x10, 0x97, 0x4e, 0x12, 0x15, 0x0a,
	0x10, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f, 0x42, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x49,
	0x44, 0x10, 0x98, 0x4e, 0x12, 0x14, 0x0a, 0x0f, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f, 0x42,
	0x65, 0x74, 0x52, 0x65, 0x74, 0x49, 0x44, 0x10, 0x99, 0x4e, 0x12, 0x19, 0x0a, 0x14, 0x53, 0x68,
	0x61, 0x69, 0x62, 0x61, 0x6f, 0x53, 0x68, 0x61, 0x6e, 0x67, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67,
	0x49, 0x44, 0x10, 0x9a, 0x4e, 0x12, 0x1c, 0x0a, 0x17, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f,
	0x53, 0x68, 0x61, 0x6e, 0x67, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x52, 0x65, 0x74, 0x49, 0x44,
	0x10, 0x9b, 0x4e, 0x12, 0x20, 0x0a, 0x1b, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f, 0x47, 0x65,
	0x74, 0x53, 0x68, 0x61, 0x6e, 0x67, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x44, 0x10, 0x9c, 0x4e, 0x12, 0x21, 0x0a, 0x1c, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f,
	0x53, 0x68, 0x61, 0x6e, 0x67, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x49, 0x44, 0x10, 0x9d, 0x4e, 0x12, 0x18, 0x0a, 0x13, 0x53, 0x68, 0x61, 0x69,
	0x62, 0x61, 0x6f, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x10,
	0x9e, 0x4e, 0x12, 0x16, 0x0a, 0x11, 0x53, 0x68, 0x61, 0x69, 0x62, 0x61, 0x6f, 0x42, 0x65, 0x74,
	0x43, 0x6c, 0x65, 0x61, 0x72, 0x49, 0x44, 0x10, 0x9f, 0x4e, 0x12, 0x1c, 0x0a, 0x17, 0x53, 0x68,
	0x61, 0x69, 0x62, 0x61, 0x6f, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x43, 0x61, 0x72, 0x64, 0x4c,
	0x69, 0x73, 0x74, 0x49, 0x44, 0x10, 0xa0, 0x4e, 0x12, 0x19, 0x0a, 0x14, 0x53, 0x68, 0x61, 0x69,
	0x62, 0x61, 0x6f, 0x42, 0x65, 0x74, 0x43, 0x6c, 0x65, 0x61, 0x72, 0x52, 0x65, 0x74, 0x49, 0x44,
	0x10, 0xa1, 0x4e, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74,
	0x6f,
}

var (
	file_Game_Shaibao_proto_rawDescOnce sync.Once
	file_Game_Shaibao_proto_rawDescData = file_Game_Shaibao_proto_rawDesc
)

func file_Game_Shaibao_proto_rawDescGZIP() []byte {
	file_Game_Shaibao_proto_rawDescOnce.Do(func() {
		file_Game_Shaibao_proto_rawDescData = protoimpl.X.CompressGZIP(file_Game_Shaibao_proto_rawDescData)
	})
	return file_Game_Shaibao_proto_rawDescData
}

var file_Game_Shaibao_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Game_Shaibao_proto_goTypes = []interface{}{
	(Shaibao_GameMessageClassID)(0), // 0: netproto.Shaibao_GameMessageClassID
}
var file_Game_Shaibao_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Game_Shaibao_proto_init() }
func file_Game_Shaibao_proto_init() {
	if File_Game_Shaibao_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Game_Shaibao_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Game_Shaibao_proto_goTypes,
		DependencyIndexes: file_Game_Shaibao_proto_depIdxs,
		EnumInfos:         file_Game_Shaibao_proto_enumTypes,
	}.Build()
	File_Game_Shaibao_proto = out.File
	file_Game_Shaibao_proto_rawDesc = nil
	file_Game_Shaibao_proto_goTypes = nil
	file_Game_Shaibao_proto_depIdxs = nil
}