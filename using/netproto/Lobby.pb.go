// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.13.0
// source: Lobby.proto

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

type LobbyMsgClassID int32

const (
	LobbyMsgClassID_GetHallIds              LobbyMsgClassID = 1 //取得大廳資訊
	LobbyMsgClassID_GetHallRetIds           LobbyMsgClassID = 2 //返回大廳資訊
	LobbyMsgClassID_AllGetHallIds           LobbyMsgClassID = 3 //取得大廳資訊
	LobbyMsgClassID_AllGetHallRetIds        LobbyMsgClassID = 4 //所有玩家返回大廳資訊
	LobbyMsgClassID_GetLobbySetConfig       LobbyMsgClassID = 5 //取得大廳設定配置
	LobbyMsgClassID_GetLobbySetConfigRet    LobbyMsgClassID = 6 //返回大廳設定配置
	LobbyMsgClassID_GuestLobbyLoginID       LobbyMsgClassID = 7
	LobbyMsgClassID_LobbyLoginRetID         LobbyMsgClassID = 8
	LobbyMsgClassID_UserLobbyLogoutID       LobbyMsgClassID = 9  //lobby 登出
	LobbyMsgClassID_GetAllLobbySetConfig    LobbyMsgClassID = 10 //大廳設定配置
	LobbyMsgClassID_GetAllLobbySetConfigRet LobbyMsgClassID = 11 //返回大廳設定配置
)

// Enum value maps for LobbyMsgClassID.
var (
	LobbyMsgClassID_name = map[int32]string{
		1:  "GetHallIds",
		2:  "GetHallRetIds",
		3:  "AllGetHallIds",
		4:  "AllGetHallRetIds",
		5:  "GetLobbySetConfig",
		6:  "GetLobbySetConfigRet",
		7:  "GuestLobbyLoginID",
		8:  "LobbyLoginRetID",
		9:  "UserLobbyLogoutID",
		10: "GetAllLobbySetConfig",
		11: "GetAllLobbySetConfigRet",
	}
	LobbyMsgClassID_value = map[string]int32{
		"GetHallIds":              1,
		"GetHallRetIds":           2,
		"AllGetHallIds":           3,
		"AllGetHallRetIds":        4,
		"GetLobbySetConfig":       5,
		"GetLobbySetConfigRet":    6,
		"GuestLobbyLoginID":       7,
		"LobbyLoginRetID":         8,
		"UserLobbyLogoutID":       9,
		"GetAllLobbySetConfig":    10,
		"GetAllLobbySetConfigRet": 11,
	}
)

func (x LobbyMsgClassID) Enum() *LobbyMsgClassID {
	p := new(LobbyMsgClassID)
	*p = x
	return p
}

func (x LobbyMsgClassID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LobbyMsgClassID) Descriptor() protoreflect.EnumDescriptor {
	return file_Lobby_proto_enumTypes[0].Descriptor()
}

func (LobbyMsgClassID) Type() protoreflect.EnumType {
	return &file_Lobby_proto_enumTypes[0]
}

func (x LobbyMsgClassID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LobbyMsgClassID) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LobbyMsgClassID(num)
	return nil
}

// Deprecated: Use LobbyMsgClassID.Descriptor instead.
func (LobbyMsgClassID) EnumDescriptor() ([]byte, []int) {
	return file_Lobby_proto_rawDescGZIP(), []int{0}
}

var File_Lobby_proto protoreflect.FileDescriptor

var file_Lobby_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x6e,
	0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x88, 0x02, 0x0a, 0x0f, 0x4c, 0x6f, 0x62, 0x62,
	0x79, 0x4d, 0x73, 0x67, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x48, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x73, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x47,
	0x65, 0x74, 0x48, 0x61, 0x6c, 0x6c, 0x52, 0x65, 0x74, 0x49, 0x64, 0x73, 0x10, 0x02, 0x12, 0x11,
	0x0a, 0x0d, 0x41, 0x6c, 0x6c, 0x47, 0x65, 0x74, 0x48, 0x61, 0x6c, 0x6c, 0x49, 0x64, 0x73, 0x10,
	0x03, 0x12, 0x14, 0x0a, 0x10, 0x41, 0x6c, 0x6c, 0x47, 0x65, 0x74, 0x48, 0x61, 0x6c, 0x6c, 0x52,
	0x65, 0x74, 0x49, 0x64, 0x73, 0x10, 0x04, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4c, 0x6f,
	0x62, 0x62, 0x79, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x10, 0x05, 0x12, 0x18,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x74, 0x10, 0x06, 0x12, 0x15, 0x0a, 0x11, 0x47, 0x75, 0x65, 0x73,
	0x74, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x44, 0x10, 0x07, 0x12,
	0x13, 0x0a, 0x0f, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x74,
	0x49, 0x44, 0x10, 0x08, 0x12, 0x15, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x62, 0x62,
	0x79, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x49, 0x44, 0x10, 0x09, 0x12, 0x18, 0x0a, 0x14, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c, 0x6f, 0x62, 0x62, 0x79, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x10, 0x0a, 0x12, 0x1b, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x4c,
	0x6f, 0x62, 0x62, 0x79, 0x53, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x74,
	0x10, 0x0b, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
}

var (
	file_Lobby_proto_rawDescOnce sync.Once
	file_Lobby_proto_rawDescData = file_Lobby_proto_rawDesc
)

func file_Lobby_proto_rawDescGZIP() []byte {
	file_Lobby_proto_rawDescOnce.Do(func() {
		file_Lobby_proto_rawDescData = protoimpl.X.CompressGZIP(file_Lobby_proto_rawDescData)
	})
	return file_Lobby_proto_rawDescData
}

var file_Lobby_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Lobby_proto_goTypes = []interface{}{
	(LobbyMsgClassID)(0), // 0: netproto.LobbyMsgClassID
}
var file_Lobby_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Lobby_proto_init() }
func file_Lobby_proto_init() {
	if File_Lobby_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Lobby_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Lobby_proto_goTypes,
		DependencyIndexes: file_Lobby_proto_depIdxs,
		EnumInfos:         file_Lobby_proto_enumTypes,
	}.Build()
	File_Lobby_proto = out.File
	file_Lobby_proto_rawDesc = nil
	file_Lobby_proto_goTypes = nil
	file_Lobby_proto_depIdxs = nil
}