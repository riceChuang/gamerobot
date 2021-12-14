// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: Game_FQZS.proto

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

type FQZS_GameMessageClassID int32

const (
	FQZS_GameMessageClassID_FQZSSettingID             FQZS_GameMessageClassID = 10001 //S2C  设置相关
	FQZS_GameMessageClassID_FQZSStateID               FQZS_GameMessageClassID = 10002 //S2C  阶段ID
	FQZS_GameMessageClassID_FQZSBetID                 FQZS_GameMessageClassID = 10003 //C2S  下注ID
	FQZS_GameMessageClassID_FQZSBroadCastBetID        FQZS_GameMessageClassID = 10004 //S2C  广播下注
	FQZS_GameMessageClassID_FQZSGameResultID          FQZS_GameMessageClassID = 10005 //S2C  游戏结果扑克
	FQZS_GameMessageClassID_FQZSGameResultUserMoneyID FQZS_GameMessageClassID = 10007 //S2C  游戏结果赢钱
	FQZS_GameMessageClassID_FQZSBetAckID              FQZS_GameMessageClassID = 10009 //S2C  下注结果
	FQZS_GameMessageClassID_FQZSResultListID          FQZS_GameMessageClassID = 10016 //前20次开牌结果
	FQZS_GameMessageClassID_FQZSGameNoInfoID          FQZS_GameMessageClassID = 10017 //賽局編號
	FQZS_GameMessageClassID_FQZSSceneID               FQZS_GameMessageClassID = 10018 //遊戲當前資訊與設定
)

// Enum value maps for FQZS_GameMessageClassID.
var (
	FQZS_GameMessageClassID_name = map[int32]string{
		10001: "FQZSSettingID",
		10002: "FQZSStateID",
		10003: "FQZSBetID",
		10004: "FQZSBroadCastBetID",
		10005: "FQZSGameResultID",
		10007: "FQZSGameResultUserMoneyID",
		10009: "FQZSBetAckID",
		10016: "FQZSResultListID",
		10017: "FQZSGameNoInfoID",
		10018: "FQZSSceneID",
	}
	FQZS_GameMessageClassID_value = map[string]int32{
		"FQZSSettingID":             10001,
		"FQZSStateID":               10002,
		"FQZSBetID":                 10003,
		"FQZSBroadCastBetID":        10004,
		"FQZSGameResultID":          10005,
		"FQZSGameResultUserMoneyID": 10007,
		"FQZSBetAckID":              10009,
		"FQZSResultListID":          10016,
		"FQZSGameNoInfoID":          10017,
		"FQZSSceneID":               10018,
	}
)

func (x FQZS_GameMessageClassID) Enum() *FQZS_GameMessageClassID {
	p := new(FQZS_GameMessageClassID)
	*p = x
	return p
}

func (x FQZS_GameMessageClassID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FQZS_GameMessageClassID) Descriptor() protoreflect.EnumDescriptor {
	return file_Game_FQZS_proto_enumTypes[0].Descriptor()
}

func (FQZS_GameMessageClassID) Type() protoreflect.EnumType {
	return &file_Game_FQZS_proto_enumTypes[0]
}

func (x FQZS_GameMessageClassID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *FQZS_GameMessageClassID) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = FQZS_GameMessageClassID(num)
	return nil
}

// Deprecated: Use FQZS_GameMessageClassID.Descriptor instead.
func (FQZS_GameMessageClassID) EnumDescriptor() ([]byte, []int) {
	return file_Game_FQZS_proto_rawDescGZIP(), []int{0}
}

var File_Game_FQZS_proto protoreflect.FileDescriptor

var file_Game_FQZS_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x47, 0x61, 0x6d, 0x65, 0x5f, 0x46, 0x51, 0x5a, 0x53, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0xf2, 0x01, 0x0a, 0x17,
	0x46, 0x51, 0x5a, 0x53, 0x5f, 0x47, 0x61, 0x6d, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x0d, 0x46, 0x51, 0x5a, 0x53, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x49, 0x44, 0x10, 0x91, 0x4e, 0x12, 0x10, 0x0a, 0x0b, 0x46,
	0x51, 0x5a, 0x53, 0x53, 0x74, 0x61, 0x74, 0x65, 0x49, 0x44, 0x10, 0x92, 0x4e, 0x12, 0x0e, 0x0a,
	0x09, 0x46, 0x51, 0x5a, 0x53, 0x42, 0x65, 0x74, 0x49, 0x44, 0x10, 0x93, 0x4e, 0x12, 0x17, 0x0a,
	0x12, 0x46, 0x51, 0x5a, 0x53, 0x42, 0x72, 0x6f, 0x61, 0x64, 0x43, 0x61, 0x73, 0x74, 0x42, 0x65,
	0x74, 0x49, 0x44, 0x10, 0x94, 0x4e, 0x12, 0x15, 0x0a, 0x10, 0x46, 0x51, 0x5a, 0x53, 0x47, 0x61,
	0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x49, 0x44, 0x10, 0x95, 0x4e, 0x12, 0x1e, 0x0a,
	0x19, 0x46, 0x51, 0x5a, 0x53, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x49, 0x44, 0x10, 0x97, 0x4e, 0x12, 0x11, 0x0a,
	0x0c, 0x46, 0x51, 0x5a, 0x53, 0x42, 0x65, 0x74, 0x41, 0x63, 0x6b, 0x49, 0x44, 0x10, 0x99, 0x4e,
	0x12, 0x15, 0x0a, 0x10, 0x46, 0x51, 0x5a, 0x53, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x49, 0x44, 0x10, 0xa0, 0x4e, 0x12, 0x15, 0x0a, 0x10, 0x46, 0x51, 0x5a, 0x53, 0x47,
	0x61, 0x6d, 0x65, 0x4e, 0x6f, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x10, 0xa1, 0x4e, 0x12, 0x10,
	0x0a, 0x0b, 0x46, 0x51, 0x5a, 0x53, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x44, 0x10, 0xa2, 0x4e,
	0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_Game_FQZS_proto_rawDescOnce sync.Once
	file_Game_FQZS_proto_rawDescData = file_Game_FQZS_proto_rawDesc
)

func file_Game_FQZS_proto_rawDescGZIP() []byte {
	file_Game_FQZS_proto_rawDescOnce.Do(func() {
		file_Game_FQZS_proto_rawDescData = protoimpl.X.CompressGZIP(file_Game_FQZS_proto_rawDescData)
	})
	return file_Game_FQZS_proto_rawDescData
}

var file_Game_FQZS_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Game_FQZS_proto_goTypes = []interface{}{
	(FQZS_GameMessageClassID)(0), // 0: netproto.FQZS_GameMessageClassID
}
var file_Game_FQZS_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Game_FQZS_proto_init() }
func file_Game_FQZS_proto_init() {
	if File_Game_FQZS_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Game_FQZS_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Game_FQZS_proto_goTypes,
		DependencyIndexes: file_Game_FQZS_proto_depIdxs,
		EnumInfos:         file_Game_FQZS_proto_enumTypes,
	}.Build()
	File_Game_FQZS_proto = out.File
	file_Game_FQZS_proto_rawDesc = nil
	file_Game_FQZS_proto_goTypes = nil
	file_Game_FQZS_proto_depIdxs = nil
}
