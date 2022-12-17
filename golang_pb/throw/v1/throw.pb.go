// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: throw.proto

package v1

import (
	_struct "github.com/golang/protobuf/ptypes/struct"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IRequestProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action     string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	Method     string `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Callback   string `protobuf:"bytes,3,opt,name=callback,proto3" json:"callback,omitempty"`
	IsCompress bool   `protobuf:"varint,4,opt,name=isCompress,proto3" json:"isCompress,omitempty"` // true 表示 data 为 proto 编码的字节数组
	ChannelId  int32  `protobuf:"varint,5,opt,name=channelId,proto3" json:"channelId,omitempty"`
	Data       []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"` // data 是 proto 编码的字节
}

func (x *IRequestProtocol) Reset() {
	*x = IRequestProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_throw_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IRequestProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IRequestProtocol) ProtoMessage() {}

func (x *IRequestProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_throw_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IRequestProtocol.ProtoReflect.Descriptor instead.
func (*IRequestProtocol) Descriptor() ([]byte, []int) {
	return file_throw_proto_rawDescGZIP(), []int{0}
}

func (x *IRequestProtocol) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *IRequestProtocol) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *IRequestProtocol) GetCallback() string {
	if x != nil {
		return x.Callback
	}
	return ""
}

func (x *IRequestProtocol) GetIsCompress() bool {
	if x != nil {
		return x.IsCompress
	}
	return false
}

func (x *IRequestProtocol) GetChannelId() int32 {
	if x != nil {
		return x.ChannelId
	}
	return 0
}

func (x *IRequestProtocol) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type IResponseProtocol struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code       int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	IsCompress bool   `protobuf:"varint,2,opt,name=isCompress,proto3" json:"isCompress,omitempty"` // true 表示 data 为 proto 编码的字节数组
	Data       []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`              // data 是 proto 编码的字节
	Callback   string `protobuf:"bytes,4,opt,name=callback,proto3" json:"callback,omitempty"`
}

func (x *IResponseProtocol) Reset() {
	*x = IResponseProtocol{}
	if protoimpl.UnsafeEnabled {
		mi := &file_throw_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IResponseProtocol) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IResponseProtocol) ProtoMessage() {}

func (x *IResponseProtocol) ProtoReflect() protoreflect.Message {
	mi := &file_throw_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IResponseProtocol.ProtoReflect.Descriptor instead.
func (*IResponseProtocol) Descriptor() ([]byte, []int) {
	return file_throw_proto_rawDescGZIP(), []int{1}
}

func (x *IResponseProtocol) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *IResponseProtocol) GetIsCompress() bool {
	if x != nil {
		return x.IsCompress
	}
	return false
}

func (x *IResponseProtocol) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *IResponseProtocol) GetCallback() string {
	if x != nil {
		return x.Callback
	}
	return ""
}

type ReqData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route string          `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	Data  *_struct.Struct `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *ReqData) Reset() {
	*x = ReqData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_throw_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReqData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReqData) ProtoMessage() {}

func (x *ReqData) ProtoReflect() protoreflect.Message {
	mi := &file_throw_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReqData.ProtoReflect.Descriptor instead.
func (*ReqData) Descriptor() ([]byte, []int) {
	return file_throw_proto_rawDescGZIP(), []int{2}
}

func (x *ReqData) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *ReqData) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

type RespData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32           `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string          `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Data *_struct.Struct `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *RespData) Reset() {
	*x = RespData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_throw_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespData) ProtoMessage() {}

func (x *RespData) ProtoReflect() protoreflect.Message {
	mi := &file_throw_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespData.ProtoReflect.Descriptor instead.
func (*RespData) Descriptor() ([]byte, []int) {
	return file_throw_proto_rawDescGZIP(), []int{3}
}

func (x *RespData) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *RespData) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *RespData) GetData() *_struct.Struct {
	if x != nil {
		return x.Data
	}
	return nil
}

// for test
type DataInfoResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DataInfoResp) Reset() {
	*x = DataInfoResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_throw_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataInfoResp) ProtoMessage() {}

func (x *DataInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_throw_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataInfoResp.ProtoReflect.Descriptor instead.
func (*DataInfoResp) Descriptor() ([]byte, []int) {
	return file_throw_proto_rawDescGZIP(), []int{4}
}

func (x *DataInfoResp) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *DataInfoResp) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_throw_proto protoreflect.FileDescriptor

var file_throw_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x74, 0x68, 0x72, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74,
	0x68, 0x72, 0x6f, 0x77, 0x2e, 0x76, 0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb0, 0x01, 0x0a, 0x10, 0x49, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x12, 0x16,
	0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x1a,
	0x0a, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x63, 0x61, 0x6c, 0x6c, 0x62, 0x61, 0x63, 0x6b, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73,
	0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x68,
	0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63,
	0x68, 0x61, 0x6e, 0x6e, 0x65, 0x6c, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x77, 0x0a, 0x11,
	0x49, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f,
	0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x43, 0x6f, 0x6d,
	0x70, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x61, 0x6c,
	0x6c, 0x62, 0x61, 0x63, 0x6b, 0x22, 0x4c, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x22, 0x5d, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x2b, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x52, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0x34, 0x0a, 0x0c, 0x44, 0x61, 0x74, 0x61, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x75, 0x68, 0x61, 0x6e, 0x79, 0x75, 0x6a, 0x69,
	0x65, 0x2f, 0x74, 0x68, 0x72, 0x6f, 0x77, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x66, 0x61, 0x63,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x68, 0x72, 0x6f, 0x77, 0x2f, 0x76, 0x31,
	0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_throw_proto_rawDescOnce sync.Once
	file_throw_proto_rawDescData = file_throw_proto_rawDesc
)

func file_throw_proto_rawDescGZIP() []byte {
	file_throw_proto_rawDescOnce.Do(func() {
		file_throw_proto_rawDescData = protoimpl.X.CompressGZIP(file_throw_proto_rawDescData)
	})
	return file_throw_proto_rawDescData
}

var file_throw_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_throw_proto_goTypes = []interface{}{
	(*IRequestProtocol)(nil),  // 0: throw.v1.IRequestProtocol
	(*IResponseProtocol)(nil), // 1: throw.v1.IResponseProtocol
	(*ReqData)(nil),           // 2: throw.v1.ReqData
	(*RespData)(nil),          // 3: throw.v1.RespData
	(*DataInfoResp)(nil),      // 4: throw.v1.DataInfoResp
	(*_struct.Struct)(nil),    // 5: google.protobuf.Struct
}
var file_throw_proto_depIdxs = []int32{
	5, // 0: throw.v1.ReqData.data:type_name -> google.protobuf.Struct
	5, // 1: throw.v1.RespData.data:type_name -> google.protobuf.Struct
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_throw_proto_init() }
func file_throw_proto_init() {
	if File_throw_proto != nil {
		return
	}
	file_game_model_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_throw_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IRequestProtocol); i {
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
		file_throw_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IResponseProtocol); i {
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
		file_throw_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReqData); i {
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
		file_throw_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespData); i {
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
		file_throw_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataInfoResp); i {
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
			RawDescriptor: file_throw_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_throw_proto_goTypes,
		DependencyIndexes: file_throw_proto_depIdxs,
		MessageInfos:      file_throw_proto_msgTypes,
	}.Build()
	File_throw_proto = out.File
	file_throw_proto_rawDesc = nil
	file_throw_proto_goTypes = nil
	file_throw_proto_depIdxs = nil
}
