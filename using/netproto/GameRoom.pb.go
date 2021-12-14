// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.13.0
// source: GameRoom.proto

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

type GameRoomClassID int32

const (
	GameRoomClassID_LoginRoomID      GameRoomClassID = 1
	GameRoomClassID_LoginRoomRetID   GameRoomClassID = 2
	GameRoomClassID_UserQueueID      GameRoomClassID = 3
	GameRoomClassID_GameReadyID      GameRoomClassID = 4
	GameRoomClassID_RequestGameVerID GameRoomClassID = 5
	GameRoomClassID_GameVerInfoID    GameRoomClassID = 6
	GameRoomClassID_KickRoomAllRobot GameRoomClassID = 7
)

// Enum value maps for GameRoomClassID.
var (
	GameRoomClassID_name = map[int32]string{
		1: "LoginRoomID",
		2: "LoginRoomRetID",
		3: "UserQueueID",
		4: "GameReadyID",
		5: "RequestGameVerID",
		6: "GameVerInfoID",
		7: "KickRoomAllRobot",
	}
	GameRoomClassID_value = map[string]int32{
		"LoginRoomID":      1,
		"LoginRoomRetID":   2,
		"UserQueueID":      3,
		"GameReadyID":      4,
		"RequestGameVerID": 5,
		"GameVerInfoID":    6,
		"KickRoomAllRobot": 7,
	}
)

func (x GameRoomClassID) Enum() *GameRoomClassID {
	p := new(GameRoomClassID)
	*p = x
	return p
}

func (x GameRoomClassID) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GameRoomClassID) Descriptor() protoreflect.EnumDescriptor {
	return file_GameRoom_proto_enumTypes[0].Descriptor()
}

func (GameRoomClassID) Type() protoreflect.EnumType {
	return &file_GameRoom_proto_enumTypes[0]
}

func (x GameRoomClassID) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *GameRoomClassID) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = GameRoomClassID(num)
	return nil
}

// Deprecated: Use GameRoomClassID.Descriptor instead.
func (GameRoomClassID) EnumDescriptor() ([]byte, []int) {
	return file_GameRoom_proto_rawDescGZIP(), []int{0}
}

var File_GameRoom_proto protoreflect.FileDescriptor

var file_GameRoom_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x97, 0x01, 0x0a, 0x0f, 0x47,
	0x61, 0x6d, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x12, 0x0f,
	0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x44, 0x10, 0x01, 0x12,
	0x12, 0x0a, 0x0e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x74, 0x49,
	0x44, 0x10, 0x02, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x49, 0x44, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x61, 0x64,
	0x79, 0x49, 0x44, 0x10, 0x04, 0x12, 0x14, 0x0a, 0x10, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x47, 0x61, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x49, 0x44, 0x10, 0x05, 0x12, 0x11, 0x0a, 0x0d, 0x47,
	0x61, 0x6d, 0x65, 0x56, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x49, 0x44, 0x10, 0x06, 0x12, 0x14,
	0x0a, 0x10, 0x4b, 0x69, 0x63, 0x6b, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x6c, 0x6c, 0x52, 0x6f, 0x62,
	0x6f, 0x74, 0x10, 0x07, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x2f, 0x6e, 0x65, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f,
}

var (
	file_GameRoom_proto_rawDescOnce sync.Once
	file_GameRoom_proto_rawDescData = file_GameRoom_proto_rawDesc
)

func file_GameRoom_proto_rawDescGZIP() []byte {
	file_GameRoom_proto_rawDescOnce.Do(func() {
		file_GameRoom_proto_rawDescData = protoimpl.X.CompressGZIP(file_GameRoom_proto_rawDescData)
	})
	return file_GameRoom_proto_rawDescData
}

var file_GameRoom_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_GameRoom_proto_goTypes = []interface{}{
	(GameRoomClassID)(0), // 0: netproto.GameRoomClassID
}
var file_GameRoom_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_GameRoom_proto_init() }
func file_GameRoom_proto_init() {
	if File_GameRoom_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_GameRoom_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_GameRoom_proto_goTypes,
		DependencyIndexes: file_GameRoom_proto_depIdxs,
		EnumInfos:         file_GameRoom_proto_enumTypes,
	}.Build()
	File_GameRoom_proto = out.File
	file_GameRoom_proto_rawDesc = nil
	file_GameRoom_proto_goTypes = nil
	file_GameRoom_proto_depIdxs = nil
}