// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: Game_BAR28.proto

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

type BAR28_GameMessageClassID int32

const (
	BAR28_GameMessageClassID_BAR28PushQiangZhuangBets    BAR28_GameMessageClassID = 10001  //推送抢庄选择倍数
	BAR28_GameMessageClassID_BAR28QiangZhuangID          BAR28_GameMessageClassID = 10002  //抢庄下注
	BAR28_GameMessageClassID_BAR28NotifyQiangZhuangBetID BAR28_GameMessageClassID = 10003  //广播抢庄倍数
	BAR28_GameMessageClassID_BAR28NotifyZhuang           BAR28_GameMessageClassID = 100012 //广播本局庄家
	BAR28_GameMessageClassID_BAR28PushXianBets           BAR28_GameMessageClassID = 10004  //推送庄家选择倍数
	BAR28_GameMessageClassID_BAR28XianjiaBetTimesID      BAR28_GameMessageClassID = 10005  //闲家下注
	BAR28_GameMessageClassID_BAR28NotifyXianjiaBetID     BAR28_GameMessageClassID = 10006  //广播闲家倍数
	BAR28_GameMessageClassID_BAR28PushUserCards          BAR28_GameMessageClassID = 100013 //广播用户牌
	BAR28_GameMessageClassID_BAR28GameRetID              BAR28_GameMessageClassID = 10007  //游戏结算
	BAR28_GameMessageClassID_BAR28GameStepID             BAR28_GameMessageClassID = 10008  //游戏步骤
	BAR28_GameMessageClassID_BAR28GameSceneID            BAR28_GameMessageClassID = 100010 //重连
	BAR28_GameMessageClassID_BAR28GameNo                 BAR28_GameMessageClassID = 10011  //游戏编号
	BAR28_GameMessageClassID_BAR28PokerRecord            BAR28_GameMessageClassID = 10014  //游戏编号
	BAR28_GameMessageClassID_BAR28EarlyClosure           BAR28_GameMessageClassID = 10015  //提前结束
)

// Enum value maps for BAR28_GameMessageClassID.
var (
	BAR28_GameMessageClassID_name = map[int32]string{
		10001:  "BAR28PushQiangZhuangBets",
		10002:  "BAR28QiangZhuangID",
		10003:  "BAR28NotifyQiangZhuangBetID",
		100012: "BAR28NotifyZhuang",
		10004:  "BAR28PushXianBets",
		10005:  "BAR28XianjiaBetTimesID",
		10006:  "BAR28NotifyXianjiaBetID",
		100013: "BAR28PushUserCards",
		10007:  "BAR28GameRetID",
		10008:  "BAR28GameStepID",
		100010: "BAR28GameSceneID",
		10011:  "BAR28GameNo",
		10014:  "BAR28PokerRecord",
		10015:  "BAR28EarlyClosure",
	}
	BAR28_GameMessageClassID_value = map[string]int32{
		"BAR28PushQiangZhuangBets":    10001,
		"BAR28QiangZhuangID":          10002,
		"BAR28NotifyQiangZhuangBetID": 10003,
		"BAR28NotifyZhuang":           100012,
		"BAR28PushXianBets":           10004,
		"BAR28XianjiaBetTimesID":      10005,
		"BAR28NotifyXianjiaBetID":     10006,
		"BAR28PushUserCards":          100013,
		"BAR28GameRetID":              10007,
		"BAR28GameStepID":             10008,
		"BAR28GameSceneID":            100010,
		"BAR28GameNo":                 10011,
		"BAR28PokerRecord":            10014,
		"BAR28EarlyClosure":           10015,
	}
)

func (x BAR28_GameMessageClassID) Enum() *BAR28_GameMessageClassID {
	p := new(BAR28_GameMessageClassID)
	*p = x
	return p
}

func (x BAR28_GameMessageClassID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BAR28_GameMessageClassID) Descriptor() protoreflect.EnumDescriptor {
	return file_Game_BAR28_proto_enumTypes[0].Descriptor()
}

func (BAR28_GameMessageClassID) Type() protoreflect.EnumType {
	return &file_Game_BAR28_proto_enumTypes[0]
}

func (x BAR28_GameMessageClassID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *BAR28_GameMessageClassID) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = BAR28_GameMessageClassID(num)
	return nil
}

// Deprecated: Use BAR28_GameMessageClassID.Descriptor instead.
func (BAR28_GameMessageClassID) EnumDescriptor() ([]byte, []int) {
	return file_Game_BAR28_proto_rawDescGZIP(), []int{0}
}

var File_Game_BAR28_proto protoreflect.FileDescriptor

var file_Game_BAR28_proto_rawDesc = []byte{
	0x0a, 0x10, 0x47, 0x61, 0x6d, 0x65, 0x5f, 0x42, 0x41, 0x52, 0x32, 0x38, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xfe, 0x02, 0x0a,
	0x18, 0x42, 0x41, 0x52, 0x32, 0x38, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x1d, 0x0a, 0x18, 0x42, 0x41, 0x52,
	0x32, 0x38, 0x50, 0x75, 0x73, 0x68, 0x51, 0x69, 0x61, 0x6e, 0x67, 0x5a, 0x68, 0x75, 0x61, 0x6e,
	0x67, 0x42, 0x65, 0x74, 0x73, 0x10, 0x91, 0x4e, 0x12, 0x17, 0x0a, 0x12, 0x42, 0x41, 0x52, 0x32,
	0x38, 0x51, 0x69, 0x61, 0x6e, 0x67, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x49, 0x44, 0x10, 0x92,
	0x4e, 0x12, 0x20, 0x0a, 0x1b, 0x42, 0x41, 0x52, 0x32, 0x38, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x51, 0x69, 0x61, 0x6e, 0x67, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x42, 0x65, 0x74, 0x49, 0x44,
	0x10, 0x93, 0x4e, 0x12, 0x17, 0x0a, 0x11, 0x42, 0x41, 0x52, 0x32, 0x38, 0x4e, 0x6f, 0x74, 0x69,
	0x66, 0x79, 0x5a, 0x68, 0x75, 0x61, 0x6e, 0x67, 0x10, 0xac, 0x8d, 0x06, 0x12, 0x16, 0x0a, 0x11,
	0x42, 0x41, 0x52, 0x32, 0x38, 0x50, 0x75, 0x73, 0x68, 0x58, 0x69, 0x61, 0x6e, 0x42, 0x65, 0x74,
	0x73, 0x10, 0x94, 0x4e, 0x12, 0x1b, 0x0a, 0x16, 0x42, 0x41, 0x52, 0x32, 0x38, 0x58, 0x69, 0x61,
	0x6e, 0x6a, 0x69, 0x61, 0x42, 0x65, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x49, 0x44, 0x10, 0x95,
	0x4e, 0x12, 0x1c, 0x0a, 0x17, 0x42, 0x41, 0x52, 0x32, 0x38, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x79,
	0x58, 0x69, 0x61, 0x6e, 0x6a, 0x69, 0x61, 0x42, 0x65, 0x74, 0x49, 0x44, 0x10, 0x96, 0x4e, 0x12,
	0x18, 0x0a, 0x12, 0x42, 0x41, 0x52, 0x32, 0x38, 0x50, 0x75, 0x73, 0x68, 0x55, 0x73, 0x65, 0x72,
	0x43, 0x61, 0x72, 0x64, 0x73, 0x10, 0xad, 0x8d, 0x06, 0x12, 0x13, 0x0a, 0x0e, 0x42, 0x41, 0x52,
	0x32, 0x38, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x74, 0x49, 0x44, 0x10, 0x97, 0x4e, 0x12, 0x14,
	0x0a, 0x0f, 0x42, 0x41, 0x52, 0x32, 0x38, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x65, 0x70, 0x49,
	0x44, 0x10, 0x98, 0x4e, 0x12, 0x16, 0x0a, 0x10, 0x42, 0x41, 0x52, 0x32, 0x38, 0x47, 0x61, 0x6d,
	0x65, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x44, 0x10, 0xaa, 0x8d, 0x06, 0x12, 0x10, 0x0a, 0x0b,
	0x42, 0x41, 0x52, 0x32, 0x38, 0x47, 0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x10, 0x9b, 0x4e, 0x12, 0x15,
	0x0a, 0x10, 0x42, 0x41, 0x52, 0x32, 0x38, 0x50, 0x6f, 0x6b, 0x65, 0x72, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x10, 0x9e, 0x4e, 0x12, 0x16, 0x0a, 0x11, 0x42, 0x41, 0x52, 0x32, 0x38, 0x45, 0x61,
	0x72, 0x6c, 0x79, 0x43, 0x6c, 0x6f, 0x73, 0x75, 0x72, 0x65, 0x10, 0x9f, 0x4e, 0x42, 0x0c, 0x5a,
	0x0a, 0x2e, 0x2f, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_Game_BAR28_proto_rawDescOnce sync.Once
	file_Game_BAR28_proto_rawDescData = file_Game_BAR28_proto_rawDesc
)

func file_Game_BAR28_proto_rawDescGZIP() []byte {
	file_Game_BAR28_proto_rawDescOnce.Do(func() {
		file_Game_BAR28_proto_rawDescData = protoimpl.X.CompressGZIP(file_Game_BAR28_proto_rawDescData)
	})
	return file_Game_BAR28_proto_rawDescData
}

var file_Game_BAR28_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Game_BAR28_proto_goTypes = []interface{}{
	(BAR28_GameMessageClassID)(0), // 0: netproto.BAR28_GameMessageClassID
}
var file_Game_BAR28_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Game_BAR28_proto_init() }
func file_Game_BAR28_proto_init() {
	if File_Game_BAR28_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Game_BAR28_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Game_BAR28_proto_goTypes,
		DependencyIndexes: file_Game_BAR28_proto_depIdxs,
		EnumInfos:         file_Game_BAR28_proto_enumTypes,
	}.Build()
	File_Game_BAR28_proto = out.File
	file_Game_BAR28_proto_rawDesc = nil
	file_Game_BAR28_proto_goTypes = nil
	file_Game_BAR28_proto_depIdxs = nil
}