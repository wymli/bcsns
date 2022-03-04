// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: msg_gateway.proto

package pb

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

type PushMsgReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ContentType  string   `protobuf:"bytes,2,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	SendMsgId    string   `protobuf:"bytes,3,opt,name=send_msg_id,json=sendMsgId,proto3" json:"send_msg_id,omitempty"`       // 用于服务端去重,客户端去重
	SendTime     int64    `protobuf:"varint,4,opt,name=send_time,json=sendTime,proto3" json:"send_time,omitempty"`           // 用于接收端排序,但客户端时间不可信,除非先时间同步
	ServerTime   int64    `protobuf:"varint,5,opt,name=server_time,json=serverTime,proto3" json:"server_time,omitempty"`     // 服务端接收时间
	RecvSeq      int64    `protobuf:"varint,6,opt,name=recv_seq,json=recvSeq,proto3" json:"recv_seq,omitempty"`              // 消息序号,给接收端timeline消息编号,用于增量pull
	ServerMsgId  string   `protobuf:"bytes,7,opt,name=server_msg_id,json=serverMsgId,proto3" json:"server_msg_id,omitempty"` // 消息唯一id
	SendUid      string   `protobuf:"bytes,10,opt,name=send_uid,json=sendUid,proto3" json:"send_uid,omitempty"`              // user id
	SendNickname string   `protobuf:"bytes,11,opt,name=send_nickname,json=sendNickname,proto3" json:"send_nickname,omitempty"`
	SendAvatar   string   `protobuf:"bytes,12,opt,name=send_avatar,json=sendAvatar,proto3" json:"send_avatar,omitempty"`
	RoomId       *string  `protobuf:"bytes,19,opt,name=room_id,json=roomId,proto3,oneof" json:"room_id,omitempty"` // room id
	RecvUid      []string `protobuf:"bytes,20,rep,name=recv_uid,json=recvUid,proto3" json:"recv_uid,omitempty"`    // user id,群聊中,可能多个用户连在一台机器上,所以可以批量请求
	Data         []byte   `protobuf:"bytes,30,opt,name=data,proto3" json:"data,omitempty"`
	IsTest       *bool    `protobuf:"varint,40,opt,name=is_test,json=isTest,proto3,oneof" json:"is_test,omitempty"`
}

func (x *PushMsgReq) Reset() {
	*x = PushMsgReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgReq) ProtoMessage() {}

func (x *PushMsgReq) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgReq.ProtoReflect.Descriptor instead.
func (*PushMsgReq) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{0}
}

func (x *PushMsgReq) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *PushMsgReq) GetSendMsgId() string {
	if x != nil {
		return x.SendMsgId
	}
	return ""
}

func (x *PushMsgReq) GetSendTime() int64 {
	if x != nil {
		return x.SendTime
	}
	return 0
}

func (x *PushMsgReq) GetServerTime() int64 {
	if x != nil {
		return x.ServerTime
	}
	return 0
}

func (x *PushMsgReq) GetRecvSeq() int64 {
	if x != nil {
		return x.RecvSeq
	}
	return 0
}

func (x *PushMsgReq) GetServerMsgId() string {
	if x != nil {
		return x.ServerMsgId
	}
	return ""
}

func (x *PushMsgReq) GetSendUid() string {
	if x != nil {
		return x.SendUid
	}
	return ""
}

func (x *PushMsgReq) GetSendNickname() string {
	if x != nil {
		return x.SendNickname
	}
	return ""
}

func (x *PushMsgReq) GetSendAvatar() string {
	if x != nil {
		return x.SendAvatar
	}
	return ""
}

func (x *PushMsgReq) GetRoomId() string {
	if x != nil && x.RoomId != nil {
		return *x.RoomId
	}
	return ""
}

func (x *PushMsgReq) GetRecvUid() []string {
	if x != nil {
		return x.RecvUid
	}
	return nil
}

func (x *PushMsgReq) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PushMsgReq) GetIsTest() bool {
	if x != nil && x.IsTest != nil {
		return *x.IsTest
	}
	return false
}

type PushMsgResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PushMsgResp) Reset() {
	*x = PushMsgResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_msg_gateway_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushMsgResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushMsgResp) ProtoMessage() {}

func (x *PushMsgResp) ProtoReflect() protoreflect.Message {
	mi := &file_msg_gateway_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushMsgResp.ProtoReflect.Descriptor instead.
func (*PushMsgResp) Descriptor() ([]byte, []int) {
	return file_msg_gateway_proto_rawDescGZIP(), []int{1}
}

var File_msg_gateway_proto protoreflect.FileDescriptor

var file_msg_gateway_proto_rawDesc = []byte{
	0x0a, 0x11, 0x6d, 0x73, 0x67, 0x5f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x22, 0xb0, 0x03, 0x0a,
	0x0a, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1e,
	0x0a, 0x0b, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x73, 0x65, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0a, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08,
	0x72, 0x65, 0x63, 0x76, 0x5f, 0x73, 0x65, 0x71, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x72, 0x65, 0x63, 0x76, 0x53, 0x65, 0x71, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x5f, 0x6d, 0x73, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x73, 0x67, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73,
	0x65, 0x6e, 0x64, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73,
	0x65, 0x6e, 0x64, 0x55, 0x69, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x65, 0x6e, 0x64, 0x5f, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x65, 0x6e, 0x64, 0x4e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x6e, 0x64, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x6e, 0x64, 0x41, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12, 0x1c, 0x0a, 0x07,
	0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x65,
	0x63, 0x76, 0x5f, 0x75, 0x69, 0x64, 0x18, 0x14, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x63, 0x76, 0x55, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x1e, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a, 0x07, 0x69, 0x73, 0x5f,
	0x74, 0x65, 0x73, 0x74, 0x18, 0x28, 0x20, 0x01, 0x28, 0x08, 0x48, 0x01, 0x52, 0x06, 0x69, 0x73,
	0x54, 0x65, 0x73, 0x74, 0x88, 0x01, 0x01, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x69, 0x64, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x69, 0x73, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x22,
	0x0d, 0x0a, 0x0b, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x32, 0x41,
	0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x36, 0x0a, 0x07, 0x50, 0x75, 0x73,
	0x68, 0x4d, 0x73, 0x67, 0x12, 0x13, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x50,
	0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x4d, 0x73, 0x67, 0x52, 0x65, 0x73, 0x70, 0x22,
	0x00, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_msg_gateway_proto_rawDescOnce sync.Once
	file_msg_gateway_proto_rawDescData = file_msg_gateway_proto_rawDesc
)

func file_msg_gateway_proto_rawDescGZIP() []byte {
	file_msg_gateway_proto_rawDescOnce.Do(func() {
		file_msg_gateway_proto_rawDescData = protoimpl.X.CompressGZIP(file_msg_gateway_proto_rawDescData)
	})
	return file_msg_gateway_proto_rawDescData
}

var file_msg_gateway_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_msg_gateway_proto_goTypes = []interface{}{
	(*PushMsgReq)(nil),  // 0: gateway.PushMsgReq
	(*PushMsgResp)(nil), // 1: gateway.PushMsgResp
}
var file_msg_gateway_proto_depIdxs = []int32{
	0, // 0: gateway.Gateway.PushMsg:input_type -> gateway.PushMsgReq
	1, // 1: gateway.Gateway.PushMsg:output_type -> gateway.PushMsgResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_msg_gateway_proto_init() }
func file_msg_gateway_proto_init() {
	if File_msg_gateway_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_msg_gateway_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgReq); i {
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
		file_msg_gateway_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushMsgResp); i {
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
	file_msg_gateway_proto_msgTypes[0].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_msg_gateway_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_msg_gateway_proto_goTypes,
		DependencyIndexes: file_msg_gateway_proto_depIdxs,
		MessageInfos:      file_msg_gateway_proto_msgTypes,
	}.Build()
	File_msg_gateway_proto = out.File
	file_msg_gateway_proto_rawDesc = nil
	file_msg_gateway_proto_goTypes = nil
	file_msg_gateway_proto_depIdxs = nil
}