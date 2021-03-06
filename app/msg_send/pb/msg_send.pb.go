// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: msg_send.proto

package pb

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SendUserMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"` // 内容类型
	SendMsgId   int64  `protobuf:"varint,2,opt,name=send_msg_id,json=sendMsgId,proto3" json:"send_msg_id,omitempty"`    // 用于服务端去重,客户端去重
	SendUserId  uint64 `protobuf:"varint,4,opt,name=send_user_id,json=sendUserId,proto3" json:"send_user_id,omitempty"` // sender user id
	RecvUserId  uint64 `protobuf:"varint,5,opt,name=recv_user_id,json=recvUserId,proto3" json:"recv_user_id,omitempty"` // chat room id
	Data        []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`                                  // payload body
	IsTest      *bool  `protobuf:"varint,7,opt,name=is_test,json=isTest,proto3,oneof" json:"is_test,omitempty"`         // 是否测试
}

func (x *SendUserMsgReq) Reset() {
	*x = SendUserMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_send_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendUserMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendUserMsgReq) ProtoMessage() {}

func (x *SendUserMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_send_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendUserMsgReq.ProtoReflect.Descriptor instead.
func (*SendUserMsgReq) Descriptor() ([]byte, []int) {
	return file_msg_send_proto_rawDescGZIP(), []int{0}
}

func (x *SendUserMsgReq) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *SendUserMsgReq) GetSendMsgId() int64 {
	if x != nil {
		return x.SendMsgId
	}
	return 0
}

func (x *SendUserMsgReq) GetSendUserId() uint64 {
	if x != nil {
		return x.SendUserId
	}
	return 0
}

func (x *SendUserMsgReq) GetRecvUserId() uint64 {
	if x != nil {
		return x.RecvUserId
	}
	return 0
}

func (x *SendUserMsgReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SendUserMsgReq) GetIsTest() bool {
	if x != nil && x.IsTest != nil {
		return *x.IsTest
	}
	return false
}

type SendUserMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId int64 `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *SendUserMsgResp) Reset() {
	*x = SendUserMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_send_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendUserMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendUserMsgResp) ProtoMessage() {}

func (x *SendUserMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_msg_send_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendUserMsgResp.ProtoReflect.Descriptor instead.
func (*SendUserMsgResp) Descriptor() ([]byte, []int) {
	return file_msg_send_proto_rawDescGZIP(), []int{1}
}

func (x *SendUserMsgResp) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

type SendRoomMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentType string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"` // 内容类型
	SendMsgId   int64  `protobuf:"varint,2,opt,name=send_msg_id,json=sendMsgId,proto3" json:"send_msg_id,omitempty"`    // 用于服务端去重,客户端去重
	SendUserId  uint64 `protobuf:"varint,4,opt,name=send_user_id,json=sendUserId,proto3" json:"send_user_id,omitempty"` // sender user id
	RoomId      uint64 `protobuf:"varint,5,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`               // chat room id
	Data        []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`                                  // payload body
	IsTest      *bool  `protobuf:"varint,7,opt,name=is_test,json=isTest,proto3,oneof" json:"is_test,omitempty"`         // 是否测试
}

func (x *SendRoomMsgReq) Reset() {
	*x = SendRoomMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_send_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRoomMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRoomMsgReq) ProtoMessage() {}

func (x *SendRoomMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_send_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRoomMsgReq.ProtoReflect.Descriptor instead.
func (*SendRoomMsgReq) Descriptor() ([]byte, []int) {
	return file_msg_send_proto_rawDescGZIP(), []int{2}
}

func (x *SendRoomMsgReq) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *SendRoomMsgReq) GetSendMsgId() int64 {
	if x != nil {
		return x.SendMsgId
	}
	return 0
}

func (x *SendRoomMsgReq) GetSendUserId() uint64 {
	if x != nil {
		return x.SendUserId
	}
	return 0
}

func (x *SendRoomMsgReq) GetRoomId() uint64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *SendRoomMsgReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SendRoomMsgReq) GetIsTest() bool {
	if x != nil && x.IsTest != nil {
		return *x.IsTest
	}
	return false
}

type SendRoomMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId int64 `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *SendRoomMsgResp) Reset() {
	*x = SendRoomMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_send_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendRoomMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendRoomMsgResp) ProtoMessage() {}

func (x *SendRoomMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_msg_send_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendRoomMsgResp.ProtoReflect.Descriptor instead.
func (*SendRoomMsgResp) Descriptor() ([]byte, []int) {
	return file_msg_send_proto_rawDescGZIP(), []int{3}
}

func (x *SendRoomMsgResp) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

type PostMomentsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentType   string `protobuf:"bytes,1,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`          // 内容类型
	SendMomentsId int64  `protobuf:"varint,2,opt,name=send_moments_id,json=sendMomentsId,proto3" json:"send_moments_id,omitempty"` // 用于服务端去重,客户端去重
	UserId        uint64 `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                        // sender user id
	Data          []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`                                           // payload body
	IsTest        *bool  `protobuf:"varint,6,opt,name=is_test,json=isTest,proto3,oneof" json:"is_test,omitempty"`                  // 是否测试
}

func (x *PostMomentsReq) Reset() {
	*x = PostMomentsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_send_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostMomentsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostMomentsReq) ProtoMessage() {}

func (x *PostMomentsReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_send_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostMomentsReq.ProtoReflect.Descriptor instead.
func (*PostMomentsReq) Descriptor() ([]byte, []int) {
	return file_msg_send_proto_rawDescGZIP(), []int{4}
}

func (x *PostMomentsReq) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *PostMomentsReq) GetSendMomentsId() int64 {
	if x != nil {
		return x.SendMomentsId
	}
	return 0
}

func (x *PostMomentsReq) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *PostMomentsReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PostMomentsReq) GetIsTest() bool {
	if x != nil && x.IsTest != nil {
		return *x.IsTest
	}
	return false
}

type PostMomentsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerId int64 `protobuf:"varint,1,opt,name=server_id,json=serverId,proto3" json:"server_id,omitempty"`
}

func (x *PostMomentsResp) Reset() {
	*x = PostMomentsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_send_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostMomentsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostMomentsResp) ProtoMessage() {}

func (x *PostMomentsResp) ProtoReflect() protoreflect.Message {
	mi := &file_msg_send_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostMomentsResp.ProtoReflect.Descriptor instead.
func (*PostMomentsResp) Descriptor() ([]byte, []int) {
	return file_msg_send_proto_rawDescGZIP(), []int{5}
}

func (x *PostMomentsResp) GetServerId() int64 {
	if x != nil {
		return x.ServerId
	}
	return 0
}

var File_msg_send_proto protoreflect.FileDescriptor

var file_msg_send_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x73, 0x67, 0x5f, 0x73, 0x65, 0x6e, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x02, 0x70, 0x62, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xd5, 0x01, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4d,
	0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64,
	0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73,
	0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64,
	0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a,
	0x73, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0c, 0x72, 0x65,
	0x63, 0x76, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x0a, 0x72, 0x65, 0x63, 0x76, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x1c, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x48, 0x00, 0x52, 0x06, 0x69, 0x73, 0x54, 0x65, 0x73, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a,
	0x0a, 0x08, 0x5f, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x0f, 0x53, 0x65,
	0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0xcc, 0x01, 0x0a, 0x0e, 0x53,
	0x65, 0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1e, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x64,
	0x12, 0x20, 0x0a, 0x0c, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x1c, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x48, 0x00, 0x52, 0x06, 0x69, 0x73, 0x54, 0x65, 0x73, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a,
	0x08, 0x5f, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x22, 0x2e, 0x0a, 0x0f, 0x53, 0x65, 0x6e,
	0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0xb2, 0x01, 0x0a, 0x0e, 0x50, 0x6f,
	0x73, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x26, 0x0a, 0x0f, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x69, 0x73, 0x54, 0x65, 0x73, 0x74, 0x88,
	0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x22, 0x2e,
	0x0a, 0x0f, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x32, 0x8d,
	0x02, 0x0a, 0x07, 0x4d, 0x73, 0x67, 0x53, 0x65, 0x6e, 0x64, 0x12, 0x53, 0x0a, 0x0b, 0x53, 0x65,
	0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x22, 0x10, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x73, 0x67, 0x3a, 0x01, 0x2a, 0x12,
	0x54, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x73, 0x67, 0x12, 0x12,
	0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x4d, 0x73, 0x67, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d,
	0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x22,
	0x11, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x6d,
	0x73, 0x67, 0x3a, 0x01, 0x2a, 0x12, 0x57, 0x0a, 0x0b, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x6f, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x12, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x4d, 0x6f,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x6f,
	0x73, 0x74, 0x4d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x1f, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x19, 0x22, 0x14, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x3a, 0x01, 0x2a, 0x42, 0x06,
	0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_msg_send_proto_rawDescOnce sync.Once
	file_msg_send_proto_rawDescData = file_msg_send_proto_rawDesc
)

func file_msg_send_proto_rawDescGZIP() []byte {
	file_msg_send_proto_rawDescOnce.Do(func() {
		file_msg_send_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_send_proto_rawDescData)
	})
	return file_msg_send_proto_rawDescData
}

var file_msg_send_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_msg_send_proto_goTypes = []interface{}{
	(*SendUserMsgReq)(nil),  // 0: pb.SendUserMsgReq
	(*SendUserMsgResp)(nil), // 1: pb.SendUserMsgResp
	(*SendRoomMsgReq)(nil),  // 2: pb.SendRoomMsgReq
	(*SendRoomMsgResp)(nil), // 3: pb.SendRoomMsgResp
	(*PostMomentsReq)(nil),  // 4: pb.PostMomentsReq
	(*PostMomentsResp)(nil), // 5: pb.PostMomentsResp
}
var file_msg_send_proto_depIdxs = []int32{
	0, // 0: pb.MsgSend.SendUserMsg:input_type -> pb.SendUserMsgReq
	2, // 1: pb.MsgSend.SendRoomMsg:input_type -> pb.SendRoomMsgReq
	4, // 2: pb.MsgSend.PostMoments:input_type -> pb.PostMomentsReq
	1, // 3: pb.MsgSend.SendUserMsg:output_type -> pb.SendUserMsgResp
	3, // 4: pb.MsgSend.SendRoomMsg:output_type -> pb.SendRoomMsgResp
	5, // 5: pb.MsgSend.PostMoments:output_type -> pb.PostMomentsResp
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msg_send_proto_init() }
func file_msg_send_proto_init() {
	if File_msg_send_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_send_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendUserMsgReq); i {
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
		file_msg_send_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendUserMsgResp); i {
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
		file_msg_send_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRoomMsgReq); i {
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
		file_msg_send_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendRoomMsgResp); i {
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
		file_msg_send_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostMomentsReq); i {
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
		file_msg_send_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PostMomentsResp); i {
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
	file_msg_send_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_msg_send_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_msg_send_proto_msgTypes[4].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_send_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msg_send_proto_goTypes,
		DependencyIndexes: file_msg_send_proto_depIdxs,
		MessageInfos:      file_msg_send_proto_msgTypes,
	}.Build()
	File_msg_send_proto = out.File
	file_msg_send_proto_rawDesc = nil
	file_msg_send_proto_goTypes = nil
	file_msg_send_proto_depIdxs = nil
}
