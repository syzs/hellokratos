// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: hellokratos.proto

package hellokratos

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

type CreateHellokratosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateHellokratosRequest) Reset() {
	*x = CreateHellokratosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellokratos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHellokratosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHellokratosRequest) ProtoMessage() {}

func (x *CreateHellokratosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hellokratos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHellokratosRequest.ProtoReflect.Descriptor instead.
func (*CreateHellokratosRequest) Descriptor() ([]byte, []int) {
	return file_hellokratos_proto_rawDescGZIP(), []int{0}
}

type CreateHellokratosReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateHellokratosReply) Reset() {
	*x = CreateHellokratosReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellokratos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateHellokratosReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateHellokratosReply) ProtoMessage() {}

func (x *CreateHellokratosReply) ProtoReflect() protoreflect.Message {
	mi := &file_hellokratos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateHellokratosReply.ProtoReflect.Descriptor instead.
func (*CreateHellokratosReply) Descriptor() ([]byte, []int) {
	return file_hellokratos_proto_rawDescGZIP(), []int{1}
}

type UpdateHellokratosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateHellokratosRequest) Reset() {
	*x = UpdateHellokratosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellokratos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHellokratosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHellokratosRequest) ProtoMessage() {}

func (x *UpdateHellokratosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hellokratos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHellokratosRequest.ProtoReflect.Descriptor instead.
func (*UpdateHellokratosRequest) Descriptor() ([]byte, []int) {
	return file_hellokratos_proto_rawDescGZIP(), []int{2}
}

type UpdateHellokratosReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateHellokratosReply) Reset() {
	*x = UpdateHellokratosReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellokratos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateHellokratosReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateHellokratosReply) ProtoMessage() {}

func (x *UpdateHellokratosReply) ProtoReflect() protoreflect.Message {
	mi := &file_hellokratos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateHellokratosReply.ProtoReflect.Descriptor instead.
func (*UpdateHellokratosReply) Descriptor() ([]byte, []int) {
	return file_hellokratos_proto_rawDescGZIP(), []int{3}
}

type DeleteHellokratosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteHellokratosRequest) Reset() {
	*x = DeleteHellokratosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellokratos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHellokratosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHellokratosRequest) ProtoMessage() {}

func (x *DeleteHellokratosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hellokratos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHellokratosRequest.ProtoReflect.Descriptor instead.
func (*DeleteHellokratosRequest) Descriptor() ([]byte, []int) {
	return file_hellokratos_proto_rawDescGZIP(), []int{4}
}

type DeleteHellokratosReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteHellokratosReply) Reset() {
	*x = DeleteHellokratosReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellokratos_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteHellokratosReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteHellokratosReply) ProtoMessage() {}

func (x *DeleteHellokratosReply) ProtoReflect() protoreflect.Message {
	mi := &file_hellokratos_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteHellokratosReply.ProtoReflect.Descriptor instead.
func (*DeleteHellokratosReply) Descriptor() ([]byte, []int) {
	return file_hellokratos_proto_rawDescGZIP(), []int{5}
}

type GetHellokratosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHellokratosRequest) Reset() {
	*x = GetHellokratosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellokratos_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHellokratosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHellokratosRequest) ProtoMessage() {}

func (x *GetHellokratosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hellokratos_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHellokratosRequest.ProtoReflect.Descriptor instead.
func (*GetHellokratosRequest) Descriptor() ([]byte, []int) {
	return file_hellokratos_proto_rawDescGZIP(), []int{6}
}

type GetHellokratosReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetHellokratosReply) Reset() {
	*x = GetHellokratosReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellokratos_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetHellokratosReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetHellokratosReply) ProtoMessage() {}

func (x *GetHellokratosReply) ProtoReflect() protoreflect.Message {
	mi := &file_hellokratos_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetHellokratosReply.ProtoReflect.Descriptor instead.
func (*GetHellokratosReply) Descriptor() ([]byte, []int) {
	return file_hellokratos_proto_rawDescGZIP(), []int{7}
}

type ListHellokratosRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListHellokratosRequest) Reset() {
	*x = ListHellokratosRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellokratos_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHellokratosRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHellokratosRequest) ProtoMessage() {}

func (x *ListHellokratosRequest) ProtoReflect() protoreflect.Message {
	mi := &file_hellokratos_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHellokratosRequest.ProtoReflect.Descriptor instead.
func (*ListHellokratosRequest) Descriptor() ([]byte, []int) {
	return file_hellokratos_proto_rawDescGZIP(), []int{8}
}

type ListHellokratosReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListHellokratosReply) Reset() {
	*x = ListHellokratosReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hellokratos_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListHellokratosReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListHellokratosReply) ProtoMessage() {}

func (x *ListHellokratosReply) ProtoReflect() protoreflect.Message {
	mi := &file_hellokratos_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListHellokratosReply.ProtoReflect.Descriptor instead.
func (*ListHellokratosReply) Descriptor() ([]byte, []int) {
	return file_hellokratos_proto_rawDescGZIP(), []int{9}
}

var File_hellokratos_proto protoreflect.FileDescriptor

var file_hellokratos_proto_rawDesc = []byte{
	0x0a, 0x11, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x22, 0x1a, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x18, 0x0a, 0x16, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x0a, 0x18, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x15, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65,
	0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x16, 0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x8b, 0x04, 0x0a, 0x0b, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x67, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x29, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68,
	0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x67, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x67, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x12,
	0x29, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x5e, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b,
	0x72, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x26, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x61, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x12, 0x27, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f,
	0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x3c, 0x0a, 0x0f, 0x61, 0x70, 0x69, 0x2e, 0x68, 0x65,
	0x6c, 0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x50, 0x01, 0x5a, 0x27, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c,
	0x6c, 0x6f, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x3b, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hellokratos_proto_rawDescOnce sync.Once
	file_hellokratos_proto_rawDescData = file_hellokratos_proto_rawDesc
)

func file_hellokratos_proto_rawDescGZIP() []byte {
	file_hellokratos_proto_rawDescOnce.Do(func() {
		file_hellokratos_proto_rawDescData = protoimpl.X.CompressGZIP(file_hellokratos_proto_rawDescData)
	})
	return file_hellokratos_proto_rawDescData
}

var file_hellokratos_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_hellokratos_proto_goTypes = []interface{}{
	(*CreateHellokratosRequest)(nil), // 0: api.hellokratos.CreateHellokratosRequest
	(*CreateHellokratosReply)(nil),   // 1: api.hellokratos.CreateHellokratosReply
	(*UpdateHellokratosRequest)(nil), // 2: api.hellokratos.UpdateHellokratosRequest
	(*UpdateHellokratosReply)(nil),   // 3: api.hellokratos.UpdateHellokratosReply
	(*DeleteHellokratosRequest)(nil), // 4: api.hellokratos.DeleteHellokratosRequest
	(*DeleteHellokratosReply)(nil),   // 5: api.hellokratos.DeleteHellokratosReply
	(*GetHellokratosRequest)(nil),    // 6: api.hellokratos.GetHellokratosRequest
	(*GetHellokratosReply)(nil),      // 7: api.hellokratos.GetHellokratosReply
	(*ListHellokratosRequest)(nil),   // 8: api.hellokratos.ListHellokratosRequest
	(*ListHellokratosReply)(nil),     // 9: api.hellokratos.ListHellokratosReply
}
var file_hellokratos_proto_depIdxs = []int32{
	0, // 0: api.hellokratos.Hellokratos.CreateHellokratos:input_type -> api.hellokratos.CreateHellokratosRequest
	2, // 1: api.hellokratos.Hellokratos.UpdateHellokratos:input_type -> api.hellokratos.UpdateHellokratosRequest
	4, // 2: api.hellokratos.Hellokratos.DeleteHellokratos:input_type -> api.hellokratos.DeleteHellokratosRequest
	6, // 3: api.hellokratos.Hellokratos.GetHellokratos:input_type -> api.hellokratos.GetHellokratosRequest
	8, // 4: api.hellokratos.Hellokratos.ListHellokratos:input_type -> api.hellokratos.ListHellokratosRequest
	1, // 5: api.hellokratos.Hellokratos.CreateHellokratos:output_type -> api.hellokratos.CreateHellokratosReply
	3, // 6: api.hellokratos.Hellokratos.UpdateHellokratos:output_type -> api.hellokratos.UpdateHellokratosReply
	5, // 7: api.hellokratos.Hellokratos.DeleteHellokratos:output_type -> api.hellokratos.DeleteHellokratosReply
	7, // 8: api.hellokratos.Hellokratos.GetHellokratos:output_type -> api.hellokratos.GetHellokratosReply
	9, // 9: api.hellokratos.Hellokratos.ListHellokratos:output_type -> api.hellokratos.ListHellokratosReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_hellokratos_proto_init() }
func file_hellokratos_proto_init() {
	if File_hellokratos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hellokratos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHellokratosRequest); i {
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
		file_hellokratos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateHellokratosReply); i {
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
		file_hellokratos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHellokratosRequest); i {
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
		file_hellokratos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateHellokratosReply); i {
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
		file_hellokratos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHellokratosRequest); i {
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
		file_hellokratos_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteHellokratosReply); i {
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
		file_hellokratos_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHellokratosRequest); i {
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
		file_hellokratos_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetHellokratosReply); i {
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
		file_hellokratos_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHellokratosRequest); i {
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
		file_hellokratos_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListHellokratosReply); i {
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
			RawDescriptor: file_hellokratos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_hellokratos_proto_goTypes,
		DependencyIndexes: file_hellokratos_proto_depIdxs,
		MessageInfos:      file_hellokratos_proto_msgTypes,
	}.Build()
	File_hellokratos_proto = out.File
	file_hellokratos_proto_rawDesc = nil
	file_hellokratos_proto_goTypes = nil
	file_hellokratos_proto_depIdxs = nil
}
