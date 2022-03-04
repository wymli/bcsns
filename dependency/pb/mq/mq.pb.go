// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: mq.proto

package mq

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

type UserChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentType string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`   // RFC 7231
	SendMsgId   string `protobuf:"bytes,3,opt,name=send_msg_id,json=sendMsgId,proto3" json:"send_msg_id,omitempty"`       // 用于服务端去重,客户端去重
	SendTime    int64  `protobuf:"varint,4,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`           // 用于接收端排序,但客户端时间不可信,除非先时间同步
	ServerTime  int64  `protobuf:"varint,5,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`     // 服务端接收时间
	ServerMsgId string `protobuf:"bytes,7,opt,name=server_msg_id,json=serverMsgId,proto3" json:"server_msg_id,omitempty"` // 消息唯一id
	SendUid     uint64 `protobuf:"varint,10,opt,name=send_uid,json=sendUid,proto3" json:"send_uid,omitempty"`             // user id
	RecvUid     uint64 `protobuf:"varint,20,opt,name=recv_uid,json=recvUid,proto3" json:"recv_uid,omitempty"`             // user id
	Data        []byte `protobuf:"bytes,30,opt,name=data,proto3" json:"data,omitempty"`                                   // payload
	IsTest      *bool  `protobuf:"varint,40,opt,name=is_test,json=isTest,proto3,oneof" json:"is_test,omitempty"`
}

func (x *UserChatMessage) Reset() {
	*x = UserChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserChatMessage) ProtoMessage() {}

func (x *UserChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserChatMessage.ProtoReflect.Descriptor instead.
func (*UserChatMessage) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{0}
}

func (x *UserChatMessage) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *UserChatMessage) GetSendMsgId() string {
	if x != nil {
		return x.SendMsgId
	}
	return ""
}

func (x *UserChatMessage) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *UserChatMessage) GetServerTime() int64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *UserChatMessage) GetServerMsgId() string {
	if x != nil {
		return x.ServerMsgId
	}
	return ""
}

func (x *UserChatMessage) GetSendUid() uint64 {
	if x != nil {
		return x.SendUid
	}
	return 0
}

func (x *UserChatMessage) GetRecvUid() uint64 {
	if x != nil {
		return x.RecvUid
	}
	return 0
}

func (x *UserChatMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *UserChatMessage) GetIsTest() bool {
	if x != nil && x.IsTest != nil {
		return *x.IsTest
	}
	return false
}

type RoomChatMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentType string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`   // RFC 7231
	SendMsgId   string `protobuf:"bytes,3,opt,name=send_msg_id,json=sendMsgId,proto3" json:"send_msg_id,omitempty"`       // 用于服务端去重,客户端去重
	SendTime    int64  `protobuf:"varint,4,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`           // 用于接收端排序,但客户端时间不可信,除非先时间同步
	ServerTime  int64  `protobuf:"varint,5,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`     // 服务端接收时间
	ServerMsgId string `protobuf:"bytes,7,opt,name=server_msg_id,json=serverMsgId,proto3" json:"server_msg_id,omitempty"` // 消息唯一id
	SendUid     uint64 `protobuf:"varint,10,opt,name=send_uid,json=sendUid,proto3" json:"send_uid,omitempty"`             // user id
	RoomId      uint64 `protobuf:"varint,19,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`                // room id
	Data        []byte `protobuf:"bytes,30,opt,name=data,proto3" json:"data,omitempty"`                                   // payload
	IsTest      *bool  `protobuf:"varint,40,opt,name=is_test,json=isTest,proto3,oneof" json:"is_test,omitempty"`
}

func (x *RoomChatMessage) Reset() {
	*x = RoomChatMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RoomChatMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoomChatMessage) ProtoMessage() {}

func (x *RoomChatMessage) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoomChatMessage.ProtoReflect.Descriptor instead.
func (*RoomChatMessage) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{1}
}

func (x *RoomChatMessage) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *RoomChatMessage) GetSendMsgId() string {
	if x != nil {
		return x.SendMsgId
	}
	return ""
}

func (x *RoomChatMessage) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *RoomChatMessage) GetServerTime() int64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *RoomChatMessage) GetServerMsgId() string {
	if x != nil {
		return x.ServerMsgId
	}
	return ""
}

func (x *RoomChatMessage) GetSendUid() uint64 {
	if x != nil {
		return x.SendUid
	}
	return 0
}

func (x *RoomChatMessage) GetRoomId() uint64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *RoomChatMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *RoomChatMessage) GetIsTest() bool {
	if x != nil && x.IsTest != nil {
		return *x.IsTest
	}
	return false
}

type Moments struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentType string `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`   // RFC 7231
	MomentsId   string `protobuf:"bytes,3,opt,name=moments_id,json=momentsId,proto3" json:"moments_id,omitempty"`         // 用于服务端去重,客户端去重
	ServerTime  int64  `protobuf:"varint,5,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`     // 服务端接收时间
	ServerMsgId string `protobuf:"bytes,7,opt,name=server_msg_id,json=serverMsgId,proto3" json:"server_msg_id,omitempty"` // 消息唯一id, 红点服务
	UserId      uint64 `protobuf:"varint,10,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`                // user id
	Data        []byte `protobuf:"bytes,30,opt,name=data,proto3" json:"data,omitempty"`                                   // payload
	IsTest      *bool  `protobuf:"varint,40,opt,name=is_test,json=isTest,proto3,oneof" json:"is_test,omitempty"`
}

func (x *Moments) Reset() {
	*x = Moments{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mq_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Moments) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Moments) ProtoMessage() {}

func (x *Moments) ProtoReflect() protoreflect.Message {
	mi := &file_mq_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Moments.ProtoReflect.Descriptor instead.
func (*Moments) Descriptor() ([]byte, []int) {
	return file_mq_proto_rawDescGZIP(), []int{2}
}

func (x *Moments) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *Moments) GetMomentsId() string {
	if x != nil {
		return x.MomentsId
	}
	return ""
}

func (x *Moments) GetServerTime() int64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *Moments) GetServerMsgId() string {
	if x != nil {
		return x.ServerMsgId
	}
	return ""
}

func (x *Moments) GetUserId() uint64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Moments) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Moments) GetIsTest() bool {
	if x != nil && x.IsTest != nil {
		return *x.IsTest
	}
	return false
}

var File_mq_proto protoreflect.FileDescriptor

var file_mq_proto_rawDesc = []byte{
	0x0a, 0x08, 0x6d, 0x71, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x6d, 0x71, 0x22, 0xaa,
	0x02, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x73,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64,
	0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6d, 0x73,
	0x67, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x5f,
	0x75, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x55,
	0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65, 0x63, 0x76, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x14,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x72, 0x65, 0x63, 0x76, 0x55, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x12, 0x1c, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x28, 0x20, 0x01,
	0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x69, 0x73, 0x54, 0x65, 0x73, 0x74, 0x88, 0x01, 0x01, 0x42,
	0x0a, 0x0a, 0x08, 0x5f, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x22, 0xa8, 0x02, 0x0a, 0x0f,
	0x52, 0x6f, 0x6f, 0x6d, 0x43, 0x68, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1e, 0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x22, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69,
	0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d,
	0x73, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x75, 0x69, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x73, 0x65, 0x6e, 0x64, 0x55, 0x69, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x04,
	0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x1e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x07,
	0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52,
	0x06, 0x69, 0x73, 0x54, 0x65, 0x73, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x69,
	0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x22, 0xe7, 0x01, 0x0a, 0x07, 0x4d, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79,
	0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x6d, 0x6f, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x6f, 0x6d, 0x65, 0x6e,
	0x74, 0x73, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74,
	0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f,
	0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x1e, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x06, 0x69, 0x73, 0x54, 0x65, 0x73,
	0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74,
	0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77,
	0x79, 0x6d, 0x6c, 0x69, 0x2f, 0x62, 0x63, 0x73, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x79, 0x2f, 0x70, 0x62, 0x2f, 0x6d, 0x71, 0x3b, 0x6d, 0x71, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mq_proto_rawDescOnce sync.Once
	file_mq_proto_rawDescData = file_mq_proto_rawDesc
)

func file_mq_proto_rawDescGZIP() []byte {
	file_mq_proto_rawDescOnce.Do(func() {
		file_mq_proto_rawDescData = protoimpl.X.CompressGZIP(file_mq_proto_rawDescData)
	})
	return file_mq_proto_rawDescData
}

var file_mq_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_mq_proto_goTypes = []interface{}{
	(*UserChatMessage)(nil), // 0: mq.UserChatMessage
	(*RoomChatMessage)(nil), // 1: mq.RoomChatMessage
	(*Moments)(nil),         // 2: mq.Moments
}
var file_mq_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mq_proto_init() }
func file_mq_proto_init() {
	if File_mq_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mq_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserChatMessage); i {
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
		file_mq_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RoomChatMessage); i {
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
		file_mq_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Moments); i {
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
	file_mq_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_mq_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_mq_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_mq_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mq_proto_goTypes,
		DependencyIndexes: file_mq_proto_depIdxs,
		MessageInfos:      file_mq_proto_msgTypes,
	}.Build()
	File_mq_proto = out.File
	file_mq_proto_rawDesc = nil
	file_mq_proto_goTypes = nil
	file_mq_proto_depIdxs = nil
}