// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: game_model.proto

package v1

import (
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

type GameStart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"` // 游戏即将开始 倒计时 3 s
}

func (x *GameStart) Reset() {
	*x = GameStart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GameStart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GameStart) ProtoMessage() {}

func (x *GameStart) ProtoReflect() protoreflect.Message {
	mi := &file_game_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GameStart.ProtoReflect.Descriptor instead.
func (*GameStart) Descriptor() ([]byte, []int) {
	return file_game_model_proto_rawDescGZIP(), []int{0}
}

func (x *GameStart) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type AttackOnceInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid int64 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"` // 发起攻击的用户 id
}

func (x *AttackOnceInfo) Reset() {
	*x = AttackOnceInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AttackOnceInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttackOnceInfo) ProtoMessage() {}

func (x *AttackOnceInfo) ProtoReflect() protoreflect.Message {
	mi := &file_game_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttackOnceInfo.ProtoReflect.Descriptor instead.
func (*AttackOnceInfo) Descriptor() ([]byte, []int) {
	return file_game_model_proto_rawDescGZIP(), []int{1}
}

func (x *AttackOnceInfo) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

type ThrowResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Players []*PlayerStatus `protobuf:"bytes,1,rep,name=players,proto3" json:"players,omitempty"`
}

func (x *ThrowResult) Reset() {
	*x = ThrowResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ThrowResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ThrowResult) ProtoMessage() {}

func (x *ThrowResult) ProtoReflect() protoreflect.Message {
	mi := &file_game_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ThrowResult.ProtoReflect.Descriptor instead.
func (*ThrowResult) Descriptor() ([]byte, []int) {
	return file_game_model_proto_rawDescGZIP(), []int{2}
}

func (x *ThrowResult) GetPlayers() []*PlayerStatus {
	if x != nil {
		return x.Players
	}
	return nil
}

type PlayerStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid  int32 `protobuf:"varint,1,opt,name=uid,proto3" json:"uid,omitempty"`
	Hp   int32 `protobuf:"varint,2,opt,name=hp,proto3" json:"hp,omitempty"`
	Name int32 `protobuf:"varint,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *PlayerStatus) Reset() {
	*x = PlayerStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerStatus) ProtoMessage() {}

func (x *PlayerStatus) ProtoReflect() protoreflect.Message {
	mi := &file_game_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerStatus.ProtoReflect.Descriptor instead.
func (*PlayerStatus) Descriptor() ([]byte, []int) {
	return file_game_model_proto_rawDescGZIP(), []int{3}
}

func (x *PlayerStatus) GetUid() int32 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *PlayerStatus) GetHp() int32 {
	if x != nil {
		return x.Hp
	}
	return 0
}

func (x *PlayerStatus) GetName() int32 {
	if x != nil {
		return x.Name
	}
	return 0
}

// 退出游戏房间 todo
type QuitRoom struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *QuitRoom) Reset() {
	*x = QuitRoom{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QuitRoom) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QuitRoom) ProtoMessage() {}

func (x *QuitRoom) ProtoReflect() protoreflect.Message {
	mi := &file_game_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QuitRoom.ProtoReflect.Descriptor instead.
func (*QuitRoom) Descriptor() ([]byte, []int) {
	return file_game_model_proto_rawDescGZIP(), []int{4}
}

var File_game_model_proto protoreflect.FileDescriptor

var file_game_model_proto_rawDesc = []byte{
	0x0a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x74, 0x68, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x22, 0x1d, 0x0a, 0x09,
	0x47, 0x61, 0x6d, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x22, 0x0a, 0x0e, 0x41,
	0x74, 0x74, 0x61, 0x63, 0x6b, 0x4f, 0x6e, 0x63, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x10, 0x0a,
	0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22,
	0x3f, 0x0a, 0x0b, 0x54, 0x68, 0x72, 0x6f, 0x77, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x30,
	0x0a, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x16, 0x2e, 0x74, 0x68, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x73,
	0x22, 0x44, 0x0a, 0x0c, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x75,
	0x69, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x68, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02,
	0x68, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x0a, 0x0a, 0x08, 0x51, 0x75, 0x69, 0x74, 0x52, 0x6f,
	0x6f, 0x6d, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x73, 0x75, 0x68, 0x61, 0x6e, 0x79, 0x75, 0x6a, 0x69, 0x65, 0x2f, 0x74, 0x68, 0x72, 0x6f,
	0x77, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x74, 0x68, 0x72, 0x6f, 0x77, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_model_proto_rawDescOnce sync.Once
	file_game_model_proto_rawDescData = file_game_model_proto_rawDesc
)

func file_game_model_proto_rawDescGZIP() []byte {
	file_game_model_proto_rawDescOnce.Do(func() {
		file_game_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_model_proto_rawDescData)
	})
	return file_game_model_proto_rawDescData
}

var file_game_model_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_game_model_proto_goTypes = []interface{}{
	(*GameStart)(nil),      // 0: throw.v1.GameStart
	(*AttackOnceInfo)(nil), // 1: throw.v1.AttackOnceInfo
	(*ThrowResult)(nil),    // 2: throw.v1.ThrowResult
	(*PlayerStatus)(nil),   // 3: throw.v1.PlayerStatus
	(*QuitRoom)(nil),       // 4: throw.v1.QuitRoom
}
var file_game_model_proto_depIdxs = []int32{
	3, // 0: throw.v1.ThrowResult.players:type_name -> throw.v1.PlayerStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_game_model_proto_init() }
func file_game_model_proto_init() {
	if File_game_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GameStart); i {
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
		file_game_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AttackOnceInfo); i {
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
		file_game_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ThrowResult); i {
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
		file_game_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerStatus); i {
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
		file_game_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QuitRoom); i {
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
			RawDescriptor: file_game_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_game_model_proto_goTypes,
		DependencyIndexes: file_game_model_proto_depIdxs,
		MessageInfos:      file_game_model_proto_msgTypes,
	}.Build()
	File_game_model_proto = out.File
	file_game_model_proto_rawDesc = nil
	file_game_model_proto_goTypes = nil
	file_game_model_proto_depIdxs = nil
}
