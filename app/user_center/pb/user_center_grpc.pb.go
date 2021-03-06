// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.15.8
// source: user_center.proto

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

// UserCenterClient is the client API for UserCenter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserCenterClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*RegisterUserResp, error)
	LoginUser(ctx context.Context, in *LoginUserReq, opts ...grpc.CallOption) (*LoginUserResp, error)
	Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowResp, error)
	Unfollow(ctx context.Context, in *UnfollowReq, opts ...grpc.CallOption) (*UnfollowResp, error)
	MakeGroup(ctx context.Context, in *MakeGroupReq, opts ...grpc.CallOption) (*MakeGroupResp, error)
	JoinGroup(ctx context.Context, in *JoinGroupReq, opts ...grpc.CallOption) (*JoinGroupResp, error)
	LeaveGroup(ctx context.Context, in *LeaveGroupReq, opts ...grpc.CallOption) (*LeaveGroupResp, error)
	GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error)
	GetGroupInfo(ctx context.Context, in *GetGroupInfoReq, opts ...grpc.CallOption) (*GetGroupInfoResp, error)
	GetGroupMemberUid(ctx context.Context, in *GetGroupMemberUidReq, opts ...grpc.CallOption) (*GetGroupMemberUidResp, error)
	GetMyFriends(ctx context.Context, in *GetMyFriendsReq, opts ...grpc.CallOption) (*GetMyFriendsResp, error)
	GetMyFollows(ctx context.Context, in *GetMyFollowsReq, opts ...grpc.CallOption) (*GetMyFollowsResp, error)
	GetMyFans(ctx context.Context, in *GetMyFansReq, opts ...grpc.CallOption) (*GetMyFansResp, error)
	GetMyFansUid(ctx context.Context, in *GetMyFansUidReq, opts ...grpc.CallOption) (*GetMyFansUidResp, error)
	GetMyGroups(ctx context.Context, in *GetMyGroupsReq, opts ...grpc.CallOption) (*GetMyGroupsResp, error)
}

type userCenterClient struct {
	cc grpc.ClientConnInterface
}

func NewUserCenterClient(cc grpc.ClientConnInterface) UserCenterClient {
	return &userCenterClient{cc}
}

func (c *userCenterClient) RegisterUser(ctx context.Context, in *RegisterUserReq, opts ...grpc.CallOption) (*RegisterUserResp, error) {
	out := new(RegisterUserResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) LoginUser(ctx context.Context, in *LoginUserReq, opts ...grpc.CallOption) (*LoginUserResp, error) {
	out := new(LoginUserResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/LoginUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) Follow(ctx context.Context, in *FollowReq, opts ...grpc.CallOption) (*FollowResp, error) {
	out := new(FollowResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/Follow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) Unfollow(ctx context.Context, in *UnfollowReq, opts ...grpc.CallOption) (*UnfollowResp, error) {
	out := new(UnfollowResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/Unfollow", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) MakeGroup(ctx context.Context, in *MakeGroupReq, opts ...grpc.CallOption) (*MakeGroupResp, error) {
	out := new(MakeGroupResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/MakeGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) JoinGroup(ctx context.Context, in *JoinGroupReq, opts ...grpc.CallOption) (*JoinGroupResp, error) {
	out := new(JoinGroupResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/JoinGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) LeaveGroup(ctx context.Context, in *LeaveGroupReq, opts ...grpc.CallOption) (*LeaveGroupResp, error) {
	out := new(LeaveGroupResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/LeaveGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetUserInfo(ctx context.Context, in *GetUserInfoReq, opts ...grpc.CallOption) (*GetUserInfoResp, error) {
	out := new(GetUserInfoResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/GetUserInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetGroupInfo(ctx context.Context, in *GetGroupInfoReq, opts ...grpc.CallOption) (*GetGroupInfoResp, error) {
	out := new(GetGroupInfoResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/GetGroupInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetGroupMemberUid(ctx context.Context, in *GetGroupMemberUidReq, opts ...grpc.CallOption) (*GetGroupMemberUidResp, error) {
	out := new(GetGroupMemberUidResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/GetGroupMemberUid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetMyFriends(ctx context.Context, in *GetMyFriendsReq, opts ...grpc.CallOption) (*GetMyFriendsResp, error) {
	out := new(GetMyFriendsResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/GetMyFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetMyFollows(ctx context.Context, in *GetMyFollowsReq, opts ...grpc.CallOption) (*GetMyFollowsResp, error) {
	out := new(GetMyFollowsResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/GetMyFollows", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetMyFans(ctx context.Context, in *GetMyFansReq, opts ...grpc.CallOption) (*GetMyFansResp, error) {
	out := new(GetMyFansResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/GetMyFans", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetMyFansUid(ctx context.Context, in *GetMyFansUidReq, opts ...grpc.CallOption) (*GetMyFansUidResp, error) {
	out := new(GetMyFansUidResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/GetMyFansUid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userCenterClient) GetMyGroups(ctx context.Context, in *GetMyGroupsReq, opts ...grpc.CallOption) (*GetMyGroupsResp, error) {
	out := new(GetMyGroupsResp)
	err := c.cc.Invoke(ctx, "/pb.UserCenter/GetMyGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserCenterServer is the server API for UserCenter service.
// All implementations must embed UnimplementedUserCenterServer
// for forward compatibility
type UserCenterServer interface {
	RegisterUser(context.Context, *RegisterUserReq) (*RegisterUserResp, error)
	LoginUser(context.Context, *LoginUserReq) (*LoginUserResp, error)
	Follow(context.Context, *FollowReq) (*FollowResp, error)
	Unfollow(context.Context, *UnfollowReq) (*UnfollowResp, error)
	MakeGroup(context.Context, *MakeGroupReq) (*MakeGroupResp, error)
	JoinGroup(context.Context, *JoinGroupReq) (*JoinGroupResp, error)
	LeaveGroup(context.Context, *LeaveGroupReq) (*LeaveGroupResp, error)
	GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error)
	GetGroupInfo(context.Context, *GetGroupInfoReq) (*GetGroupInfoResp, error)
	GetGroupMemberUid(context.Context, *GetGroupMemberUidReq) (*GetGroupMemberUidResp, error)
	GetMyFriends(context.Context, *GetMyFriendsReq) (*GetMyFriendsResp, error)
	GetMyFollows(context.Context, *GetMyFollowsReq) (*GetMyFollowsResp, error)
	GetMyFans(context.Context, *GetMyFansReq) (*GetMyFansResp, error)
	GetMyFansUid(context.Context, *GetMyFansUidReq) (*GetMyFansUidResp, error)
	GetMyGroups(context.Context, *GetMyGroupsReq) (*GetMyGroupsResp, error)
	mustEmbedUnimplementedUserCenterServer()
}

// UnimplementedUserCenterServer must be embedded to have forward compatible implementations.
type UnimplementedUserCenterServer struct {
}

func (UnimplementedUserCenterServer) RegisterUser(context.Context, *RegisterUserReq) (*RegisterUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedUserCenterServer) LoginUser(context.Context, *LoginUserReq) (*LoginUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
func (UnimplementedUserCenterServer) Follow(context.Context, *FollowReq) (*FollowResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Follow not implemented")
}
func (UnimplementedUserCenterServer) Unfollow(context.Context, *UnfollowReq) (*UnfollowResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unfollow not implemented")
}
func (UnimplementedUserCenterServer) MakeGroup(context.Context, *MakeGroupReq) (*MakeGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeGroup not implemented")
}
func (UnimplementedUserCenterServer) JoinGroup(context.Context, *JoinGroupReq) (*JoinGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGroup not implemented")
}
func (UnimplementedUserCenterServer) LeaveGroup(context.Context, *LeaveGroupReq) (*LeaveGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveGroup not implemented")
}
func (UnimplementedUserCenterServer) GetUserInfo(context.Context, *GetUserInfoReq) (*GetUserInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserInfo not implemented")
}
func (UnimplementedUserCenterServer) GetGroupInfo(context.Context, *GetGroupInfoReq) (*GetGroupInfoResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupInfo not implemented")
}
func (UnimplementedUserCenterServer) GetGroupMemberUid(context.Context, *GetGroupMemberUidReq) (*GetGroupMemberUidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupMemberUid not implemented")
}
func (UnimplementedUserCenterServer) GetMyFriends(context.Context, *GetMyFriendsReq) (*GetMyFriendsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyFriends not implemented")
}
func (UnimplementedUserCenterServer) GetMyFollows(context.Context, *GetMyFollowsReq) (*GetMyFollowsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyFollows not implemented")
}
func (UnimplementedUserCenterServer) GetMyFans(context.Context, *GetMyFansReq) (*GetMyFansResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyFans not implemented")
}
func (UnimplementedUserCenterServer) GetMyFansUid(context.Context, *GetMyFansUidReq) (*GetMyFansUidResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyFansUid not implemented")
}
func (UnimplementedUserCenterServer) GetMyGroups(context.Context, *GetMyGroupsReq) (*GetMyGroupsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyGroups not implemented")
}
func (UnimplementedUserCenterServer) mustEmbedUnimplementedUserCenterServer() {}

// UnsafeUserCenterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserCenterServer will
// result in compilation errors.
type UnsafeUserCenterServer interface {
	mustEmbedUnimplementedUserCenterServer()
}

func RegisterUserCenterServer(s grpc.ServiceRegistrar, srv UserCenterServer) {
	s.RegisterService(&UserCenter_ServiceDesc, srv)
}

func _UserCenter_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).RegisterUser(ctx, req.(*RegisterUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_LoginUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).LoginUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/LoginUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).LoginUser(ctx, req.(*LoginUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_Follow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).Follow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/Follow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).Follow(ctx, req.(*FollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_Unfollow_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UnfollowReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).Unfollow(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/Unfollow",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).Unfollow(ctx, req.(*UnfollowReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_MakeGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MakeGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).MakeGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/MakeGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).MakeGroup(ctx, req.(*MakeGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_JoinGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).JoinGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/JoinGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).JoinGroup(ctx, req.(*JoinGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_LeaveGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).LeaveGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/LeaveGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).LeaveGroup(ctx, req.(*LeaveGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetUserInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetUserInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/GetUserInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetUserInfo(ctx, req.(*GetUserInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetGroupInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetGroupInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/GetGroupInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetGroupInfo(ctx, req.(*GetGroupInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetGroupMemberUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMemberUidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetGroupMemberUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/GetGroupMemberUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetGroupMemberUid(ctx, req.(*GetGroupMemberUidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetMyFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyFriendsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetMyFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/GetMyFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetMyFriends(ctx, req.(*GetMyFriendsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetMyFollows_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyFollowsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetMyFollows(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/GetMyFollows",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetMyFollows(ctx, req.(*GetMyFollowsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetMyFans_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyFansReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetMyFans(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/GetMyFans",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetMyFans(ctx, req.(*GetMyFansReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetMyFansUid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyFansUidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetMyFansUid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/GetMyFansUid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetMyFansUid(ctx, req.(*GetMyFansUidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserCenter_GetMyGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyGroupsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserCenterServer).GetMyGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserCenter/GetMyGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserCenterServer).GetMyGroups(ctx, req.(*GetMyGroupsReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserCenter_ServiceDesc is the grpc.ServiceDesc for UserCenter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserCenter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserCenter",
	HandlerType: (*UserCenterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _UserCenter_RegisterUser_Handler,
		},
		{
			MethodName: "LoginUser",
			Handler:    _UserCenter_LoginUser_Handler,
		},
		{
			MethodName: "Follow",
			Handler:    _UserCenter_Follow_Handler,
		},
		{
			MethodName: "Unfollow",
			Handler:    _UserCenter_Unfollow_Handler,
		},
		{
			MethodName: "MakeGroup",
			Handler:    _UserCenter_MakeGroup_Handler,
		},
		{
			MethodName: "JoinGroup",
			Handler:    _UserCenter_JoinGroup_Handler,
		},
		{
			MethodName: "LeaveGroup",
			Handler:    _UserCenter_LeaveGroup_Handler,
		},
		{
			MethodName: "GetUserInfo",
			Handler:    _UserCenter_GetUserInfo_Handler,
		},
		{
			MethodName: "GetGroupInfo",
			Handler:    _UserCenter_GetGroupInfo_Handler,
		},
		{
			MethodName: "GetGroupMemberUid",
			Handler:    _UserCenter_GetGroupMemberUid_Handler,
		},
		{
			MethodName: "GetMyFriends",
			Handler:    _UserCenter_GetMyFriends_Handler,
		},
		{
			MethodName: "GetMyFollows",
			Handler:    _UserCenter_GetMyFollows_Handler,
		},
		{
			MethodName: "GetMyFans",
			Handler:    _UserCenter_GetMyFans_Handler,
		},
		{
			MethodName: "GetMyFansUid",
			Handler:    _UserCenter_GetMyFansUid_Handler,
		},
		{
			MethodName: "GetMyGroups",
			Handler:    _UserCenter_GetMyGroups_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_center.proto",
}
