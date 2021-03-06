// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: online.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// OnlineClient is the client API for Online service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OnlineClient interface {
	OnlineUser(ctx context.Context, in *OnlineUserReq, opts ...grpc.CallOption) (*OnlineUserResp, error)
	BatchOnlineUser(ctx context.Context, in *BatchOnlineUserReq, opts ...grpc.CallOption) (*BatchOnlineUserResp, error)
	OfflineUser(ctx context.Context, in *OfflineUserReq, opts ...grpc.CallOption) (*OfflineUserResp, error)
	BatchOfflineUser(ctx context.Context, in *BatchOfflineUserReq, opts ...grpc.CallOption) (*BatchOfflineUserResp, error)
	KeepAliveUser(ctx context.Context, in *KeepAliveUserReq, opts ...grpc.CallOption) (*KeepAliveUserResp, error)
	// is not online
	BatchKeepAliveUser(ctx context.Context, in *BatchKeepAliveUserReq, opts ...grpc.CallOption) (*BatchKeepAliveUserResp, error)
	GetUserGateway(ctx context.Context, in *GetUserGatewayReq, opts ...grpc.CallOption) (*GetUserGatewayResp, error)
	BatchGetUserGateway(ctx context.Context, in *BatchGetUserGatewayReq, opts ...grpc.CallOption) (*BatchGetUserGatewayResp, error)
	GetAllOnlineUser(ctx context.Context, in *GetAllOnlineUserReq, opts ...grpc.CallOption) (*GetAllOnlineUserResp, error)
}

type onlineClient struct {
	cc grpc.ClientConnInterface
}

func NewOnlineClient(cc grpc.ClientConnInterface) OnlineClient {
	return &onlineClient{cc}
}

func (c *onlineClient) OnlineUser(ctx context.Context, in *OnlineUserReq, opts ...grpc.CallOption) (*OnlineUserResp, error) {
	out := new(OnlineUserResp)
	err := c.cc.Invoke(ctx, "/pb.online/OnlineUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineClient) BatchOnlineUser(ctx context.Context, in *BatchOnlineUserReq, opts ...grpc.CallOption) (*BatchOnlineUserResp, error) {
	out := new(BatchOnlineUserResp)
	err := c.cc.Invoke(ctx, "/pb.online/BatchOnlineUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineClient) OfflineUser(ctx context.Context, in *OfflineUserReq, opts ...grpc.CallOption) (*OfflineUserResp, error) {
	out := new(OfflineUserResp)
	err := c.cc.Invoke(ctx, "/pb.online/OfflineUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineClient) BatchOfflineUser(ctx context.Context, in *BatchOfflineUserReq, opts ...grpc.CallOption) (*BatchOfflineUserResp, error) {
	out := new(BatchOfflineUserResp)
	err := c.cc.Invoke(ctx, "/pb.online/BatchOfflineUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineClient) KeepAliveUser(ctx context.Context, in *KeepAliveUserReq, opts ...grpc.CallOption) (*KeepAliveUserResp, error) {
	out := new(KeepAliveUserResp)
	err := c.cc.Invoke(ctx, "/pb.online/KeepAliveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineClient) BatchKeepAliveUser(ctx context.Context, in *BatchKeepAliveUserReq, opts ...grpc.CallOption) (*BatchKeepAliveUserResp, error) {
	out := new(BatchKeepAliveUserResp)
	err := c.cc.Invoke(ctx, "/pb.online/BatchKeepAliveUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineClient) GetUserGateway(ctx context.Context, in *GetUserGatewayReq, opts ...grpc.CallOption) (*GetUserGatewayResp, error) {
	out := new(GetUserGatewayResp)
	err := c.cc.Invoke(ctx, "/pb.online/GetUserGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineClient) BatchGetUserGateway(ctx context.Context, in *BatchGetUserGatewayReq, opts ...grpc.CallOption) (*BatchGetUserGatewayResp, error) {
	out := new(BatchGetUserGatewayResp)
	err := c.cc.Invoke(ctx, "/pb.online/BatchGetUserGateway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *onlineClient) GetAllOnlineUser(ctx context.Context, in *GetAllOnlineUserReq, opts ...grpc.CallOption) (*GetAllOnlineUserResp, error) {
	out := new(GetAllOnlineUserResp)
	err := c.cc.Invoke(ctx, "/pb.online/GetAllOnlineUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OnlineServer is the server API for Online service.
// All implementations must embed UnimplementedOnlineServer
// for forward compatibility
type OnlineServer interface {
	OnlineUser(context.Context, *OnlineUserReq) (*OnlineUserResp, error)
	BatchOnlineUser(context.Context, *BatchOnlineUserReq) (*BatchOnlineUserResp, error)
	OfflineUser(context.Context, *OfflineUserReq) (*OfflineUserResp, error)
	BatchOfflineUser(context.Context, *BatchOfflineUserReq) (*BatchOfflineUserResp, error)
	KeepAliveUser(context.Context, *KeepAliveUserReq) (*KeepAliveUserResp, error)
	// is not online
	BatchKeepAliveUser(context.Context, *BatchKeepAliveUserReq) (*BatchKeepAliveUserResp, error)
	GetUserGateway(context.Context, *GetUserGatewayReq) (*GetUserGatewayResp, error)
	BatchGetUserGateway(context.Context, *BatchGetUserGatewayReq) (*BatchGetUserGatewayResp, error)
	GetAllOnlineUser(context.Context, *GetAllOnlineUserReq) (*GetAllOnlineUserResp, error)
	mustEmbedUnimplementedOnlineServer()
}

// UnimplementedOnlineServer must be embedded to have forward compatible implementations.
type UnimplementedOnlineServer struct {
}

func (UnimplementedOnlineServer) OnlineUser(context.Context, *OnlineUserReq) (*OnlineUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnlineUser not implemented")
}
func (UnimplementedOnlineServer) BatchOnlineUser(context.Context, *BatchOnlineUserReq) (*BatchOnlineUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchOnlineUser not implemented")
}
func (UnimplementedOnlineServer) OfflineUser(context.Context, *OfflineUserReq) (*OfflineUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OfflineUser not implemented")
}
func (UnimplementedOnlineServer) BatchOfflineUser(context.Context, *BatchOfflineUserReq) (*BatchOfflineUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchOfflineUser not implemented")
}
func (UnimplementedOnlineServer) KeepAliveUser(context.Context, *KeepAliveUserReq) (*KeepAliveUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method KeepAliveUser not implemented")
}
func (UnimplementedOnlineServer) BatchKeepAliveUser(context.Context, *BatchKeepAliveUserReq) (*BatchKeepAliveUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchKeepAliveUser not implemented")
}
func (UnimplementedOnlineServer) GetUserGateway(context.Context, *GetUserGatewayReq) (*GetUserGatewayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserGateway not implemented")
}
func (UnimplementedOnlineServer) BatchGetUserGateway(context.Context, *BatchGetUserGatewayReq) (*BatchGetUserGatewayResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetUserGateway not implemented")
}
func (UnimplementedOnlineServer) GetAllOnlineUser(context.Context, *GetAllOnlineUserReq) (*GetAllOnlineUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllOnlineUser not implemented")
}
func (UnimplementedOnlineServer) mustEmbedUnimplementedOnlineServer() {}

// UnsafeOnlineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OnlineServer will
// result in compilation errors.
type UnsafeOnlineServer interface {
	mustEmbedUnimplementedOnlineServer()
}

func RegisterOnlineServer(s grpc.ServiceRegistrar, srv OnlineServer) {
	s.RegisterService(&Online_ServiceDesc, srv)
}

func _Online_OnlineUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnlineUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineServer).OnlineUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.online/OnlineUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineServer).OnlineUser(ctx, req.(*OnlineUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Online_BatchOnlineUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOnlineUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineServer).BatchOnlineUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.online/BatchOnlineUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineServer).BatchOnlineUser(ctx, req.(*BatchOnlineUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Online_OfflineUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OfflineUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineServer).OfflineUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.online/OfflineUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineServer).OfflineUser(ctx, req.(*OfflineUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Online_BatchOfflineUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchOfflineUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineServer).BatchOfflineUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.online/BatchOfflineUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineServer).BatchOfflineUser(ctx, req.(*BatchOfflineUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Online_KeepAliveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeepAliveUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineServer).KeepAliveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.online/KeepAliveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineServer).KeepAliveUser(ctx, req.(*KeepAliveUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Online_BatchKeepAliveUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchKeepAliveUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineServer).BatchKeepAliveUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.online/BatchKeepAliveUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineServer).BatchKeepAliveUser(ctx, req.(*BatchKeepAliveUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Online_GetUserGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserGatewayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineServer).GetUserGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.online/GetUserGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineServer).GetUserGateway(ctx, req.(*GetUserGatewayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Online_BatchGetUserGateway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetUserGatewayReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineServer).BatchGetUserGateway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.online/BatchGetUserGateway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineServer).BatchGetUserGateway(ctx, req.(*BatchGetUserGatewayReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Online_GetAllOnlineUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllOnlineUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OnlineServer).GetAllOnlineUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.online/GetAllOnlineUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OnlineServer).GetAllOnlineUser(ctx, req.(*GetAllOnlineUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Online_ServiceDesc is the grpc.ServiceDesc for Online service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Online_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.online",
	HandlerType: (*OnlineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnlineUser",
			Handler:    _Online_OnlineUser_Handler,
		},
		{
			MethodName: "BatchOnlineUser",
			Handler:    _Online_BatchOnlineUser_Handler,
		},
		{
			MethodName: "OfflineUser",
			Handler:    _Online_OfflineUser_Handler,
		},
		{
			MethodName: "BatchOfflineUser",
			Handler:    _Online_BatchOfflineUser_Handler,
		},
		{
			MethodName: "KeepAliveUser",
			Handler:    _Online_KeepAliveUser_Handler,
		},
		{
			MethodName: "BatchKeepAliveUser",
			Handler:    _Online_BatchKeepAliveUser_Handler,
		},
		{
			MethodName: "GetUserGateway",
			Handler:    _Online_GetUserGateway_Handler,
		},
		{
			MethodName: "BatchGetUserGateway",
			Handler:    _Online_BatchGetUserGateway_Handler,
		},
		{
			MethodName: "GetAllOnlineUser",
			Handler:    _Online_GetAllOnlineUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "online.proto",
}
