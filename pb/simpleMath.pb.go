// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.9.0
// source: simpleMath.proto

package pb

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

type AddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *AddRequest) Reset() {
	*x = AddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleMath_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRequest) ProtoMessage() {}

func (x *AddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simpleMath_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRequest.ProtoReflect.Descriptor instead.
func (*AddRequest) Descriptor() ([]byte, []int) {
	return file_simpleMath_proto_rawDescGZIP(), []int{0}
}

func (x *AddRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *AddRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type AddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res int32 `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *AddResponse) Reset() {
	*x = AddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleMath_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddResponse) ProtoMessage() {}

func (x *AddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_simpleMath_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddResponse.ProtoReflect.Descriptor instead.
func (*AddResponse) Descriptor() ([]byte, []int) {
	return file_simpleMath_proto_rawDescGZIP(), []int{1}
}

func (x *AddResponse) GetRes() int32 {
	if x != nil {
		return x.Res
	}
	return 0
}

type SubtractRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *SubtractRequest) Reset() {
	*x = SubtractRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleMath_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubtractRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubtractRequest) ProtoMessage() {}

func (x *SubtractRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simpleMath_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubtractRequest.ProtoReflect.Descriptor instead.
func (*SubtractRequest) Descriptor() ([]byte, []int) {
	return file_simpleMath_proto_rawDescGZIP(), []int{2}
}

func (x *SubtractRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *SubtractRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type SubtractResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res int32 `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *SubtractResponse) Reset() {
	*x = SubtractResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleMath_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubtractResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubtractResponse) ProtoMessage() {}

func (x *SubtractResponse) ProtoReflect() protoreflect.Message {
	mi := &file_simpleMath_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubtractResponse.ProtoReflect.Descriptor instead.
func (*SubtractResponse) Descriptor() ([]byte, []int) {
	return file_simpleMath_proto_rawDescGZIP(), []int{3}
}

func (x *SubtractResponse) GetRes() int32 {
	if x != nil {
		return x.Res
	}
	return 0
}

type MultiplyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *MultiplyRequest) Reset() {
	*x = MultiplyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleMath_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplyRequest) ProtoMessage() {}

func (x *MultiplyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simpleMath_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplyRequest.ProtoReflect.Descriptor instead.
func (*MultiplyRequest) Descriptor() ([]byte, []int) {
	return file_simpleMath_proto_rawDescGZIP(), []int{4}
}

func (x *MultiplyRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *MultiplyRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type MultiplyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res int32 `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *MultiplyResponse) Reset() {
	*x = MultiplyResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleMath_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MultiplyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MultiplyResponse) ProtoMessage() {}

func (x *MultiplyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_simpleMath_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MultiplyResponse.ProtoReflect.Descriptor instead.
func (*MultiplyResponse) Descriptor() ([]byte, []int) {
	return file_simpleMath_proto_rawDescGZIP(), []int{5}
}

func (x *MultiplyResponse) GetRes() int32 {
	if x != nil {
		return x.Res
	}
	return 0
}

type DivideRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int32 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *DivideRequest) Reset() {
	*x = DivideRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleMath_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivideRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivideRequest) ProtoMessage() {}

func (x *DivideRequest) ProtoReflect() protoreflect.Message {
	mi := &file_simpleMath_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivideRequest.ProtoReflect.Descriptor instead.
func (*DivideRequest) Descriptor() ([]byte, []int) {
	return file_simpleMath_proto_rawDescGZIP(), []int{6}
}

func (x *DivideRequest) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *DivideRequest) GetB() int32 {
	if x != nil {
		return x.B
	}
	return 0
}

type DivideResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Res int32 `protobuf:"varint,1,opt,name=res,proto3" json:"res,omitempty"`
}

func (x *DivideResponse) Reset() {
	*x = DivideResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_simpleMath_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DivideResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DivideResponse) ProtoMessage() {}

func (x *DivideResponse) ProtoReflect() protoreflect.Message {
	mi := &file_simpleMath_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DivideResponse.ProtoReflect.Descriptor instead.
func (*DivideResponse) Descriptor() ([]byte, []int) {
	return file_simpleMath_proto_rawDescGZIP(), []int{7}
}

func (x *DivideResponse) GetRes() int32 {
	if x != nil {
		return x.Res
	}
	return 0
}

var File_simpleMath_proto protoreflect.FileDescriptor

var file_simpleMath_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x68, 0x22, 0x28,
	0x0a, 0x0a, 0x61, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x22, 0x1f, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x73, 0x22, 0x2d, 0x0a, 0x0f, 0x73, 0x75, 0x62,
	0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x22, 0x24, 0x0a, 0x10, 0x73, 0x75, 0x62, 0x74,
	0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x72, 0x65, 0x73, 0x22, 0x2d,
	0x0a, 0x0f, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x61, 0x12,
	0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62, 0x22, 0x24, 0x0a,
	0x10, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03,
	0x72, 0x65, 0x73, 0x22, 0x2b, 0x0a, 0x0d, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x01, 0x62,
	0x22, 0x22, 0x0a, 0x0e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x72, 0x65, 0x73, 0x32, 0xa2, 0x02, 0x0a, 0x11, 0x53, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d,
	0x61, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x03, 0x61, 0x64,
	0x64, 0x12, 0x16, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x68, 0x2e, 0x61,
	0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x4d, 0x61, 0x74, 0x68, 0x2e, 0x61, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x08, 0x73, 0x75, 0x62, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x12, 0x1b, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x68, 0x2e, 0x73, 0x75,
	0x62, 0x74, 0x72, 0x61, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x68, 0x2e, 0x73, 0x75, 0x62, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a,
	0x08, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x12, 0x1b, 0x2e, 0x73, 0x69, 0x6d, 0x70,
	0x6c, 0x65, 0x4d, 0x61, 0x74, 0x68, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d,
	0x61, 0x74, 0x68, 0x2e, 0x6d, 0x75, 0x6c, 0x74, 0x69, 0x70, 0x6c, 0x79, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65,
	0x12, 0x19, 0x2e, 0x73, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x68, 0x2e, 0x64, 0x69,
	0x76, 0x69, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x69,
	0x6d, 0x70, 0x6c, 0x65, 0x4d, 0x61, 0x74, 0x68, 0x2e, 0x64, 0x69, 0x76, 0x69, 0x64, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_simpleMath_proto_rawDescOnce sync.Once
	file_simpleMath_proto_rawDescData = file_simpleMath_proto_rawDesc
)

func file_simpleMath_proto_rawDescGZIP() []byte {
	file_simpleMath_proto_rawDescOnce.Do(func() {
		file_simpleMath_proto_rawDescData = protoimpl.X.CompressGZIP(file_simpleMath_proto_rawDescData)
	})
	return file_simpleMath_proto_rawDescData
}

var file_simpleMath_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_simpleMath_proto_goTypes = []interface{}{
	(*AddRequest)(nil),       // 0: simpleMath.addRequest
	(*AddResponse)(nil),      // 1: simpleMath.addResponse
	(*SubtractRequest)(nil),  // 2: simpleMath.subtractRequest
	(*SubtractResponse)(nil), // 3: simpleMath.subtractResponse
	(*MultiplyRequest)(nil),  // 4: simpleMath.multiplyRequest
	(*MultiplyResponse)(nil), // 5: simpleMath.multiplyResponse
	(*DivideRequest)(nil),    // 6: simpleMath.divideRequest
	(*DivideResponse)(nil),   // 7: simpleMath.divideResponse
}
var file_simpleMath_proto_depIdxs = []int32{
	0, // 0: simpleMath.SimpleMathService.add:input_type -> simpleMath.addRequest
	2, // 1: simpleMath.SimpleMathService.subtract:input_type -> simpleMath.subtractRequest
	4, // 2: simpleMath.SimpleMathService.multiply:input_type -> simpleMath.multiplyRequest
	6, // 3: simpleMath.SimpleMathService.divide:input_type -> simpleMath.divideRequest
	1, // 4: simpleMath.SimpleMathService.add:output_type -> simpleMath.addResponse
	3, // 5: simpleMath.SimpleMathService.subtract:output_type -> simpleMath.subtractResponse
	5, // 6: simpleMath.SimpleMathService.multiply:output_type -> simpleMath.multiplyResponse
	7, // 7: simpleMath.SimpleMathService.divide:output_type -> simpleMath.divideResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_simpleMath_proto_init() }
func file_simpleMath_proto_init() {
	if File_simpleMath_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_simpleMath_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRequest); i {
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
		file_simpleMath_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddResponse); i {
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
		file_simpleMath_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubtractRequest); i {
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
		file_simpleMath_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubtractResponse); i {
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
		file_simpleMath_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplyRequest); i {
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
		file_simpleMath_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MultiplyResponse); i {
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
		file_simpleMath_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivideRequest); i {
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
		file_simpleMath_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DivideResponse); i {
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
			RawDescriptor: file_simpleMath_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_simpleMath_proto_goTypes,
		DependencyIndexes: file_simpleMath_proto_depIdxs,
		MessageInfos:      file_simpleMath_proto_msgTypes,
	}.Build()
	File_simpleMath_proto = out.File
	file_simpleMath_proto_rawDesc = nil
	file_simpleMath_proto_goTypes = nil
	file_simpleMath_proto_depIdxs = nil
}
