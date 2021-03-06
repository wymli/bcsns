// Code generated by goctl. DO NOT EDIT!
// Source: user_center.proto

package server

import (
	"context"

	"github.com/wymli/bcsns/app/user_center/internal/logic"
	"github.com/wymli/bcsns/app/user_center/internal/svc"
	"github.com/wymli/bcsns/app/user_center/pb"
)

type UserCenterServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserCenterServer
}

func NewUserCenterServer(svcCtx *svc.ServiceContext) *UserCenterServer {
	return &UserCenterServer{
		svcCtx: svcCtx,
	}
}

func (s *UserCenterServer) RegisterUser(ctx context.Context, in *pb.RegisterUserReq) (*pb.RegisterUserResp, error) {
	l := logic.NewRegisterUserLogic(ctx, s.svcCtx)
	return l.RegisterUser(in)
}

func (s *UserCenterServer) LoginUser(ctx context.Context, in *pb.LoginUserReq) (*pb.LoginUserResp, error) {
	l := logic.NewLoginUserLogic(ctx, s.svcCtx)
	return l.LoginUser(in)
}

func (s *UserCenterServer) Follow(ctx context.Context, in *pb.FollowReq) (*pb.FollowResp, error) {
	l := logic.NewFollowLogic(ctx, s.svcCtx)
	return l.Follow(in)
}

func (s *UserCenterServer) Unfollow(ctx context.Context, in *pb.UnfollowReq) (*pb.UnfollowResp, error) {
	l := logic.NewUnfollowLogic(ctx, s.svcCtx)
	return l.Unfollow(in)
}

func (s *UserCenterServer) MakeGroup(ctx context.Context, in *pb.MakeGroupReq) (*pb.MakeGroupResp, error) {
	l := logic.NewMakeGroupLogic(ctx, s.svcCtx)
	return l.MakeGroup(in)
}

func (s *UserCenterServer) JoinGroup(ctx context.Context, in *pb.JoinGroupReq) (*pb.JoinGroupResp, error) {
	l := logic.NewJoinGroupLogic(ctx, s.svcCtx)
	return l.JoinGroup(in)
}

func (s *UserCenterServer) LeaveGroup(ctx context.Context, in *pb.LeaveGroupReq) (*pb.LeaveGroupResp, error) {
	l := logic.NewLeaveGroupLogic(ctx, s.svcCtx)
	return l.LeaveGroup(in)
}

func (s *UserCenterServer) GetUserInfo(ctx context.Context, in *pb.GetUserInfoReq) (*pb.GetUserInfoResp, error) {
	l := logic.NewGetUserInfoLogic(ctx, s.svcCtx)
	return l.GetUserInfo(in)
}

func (s *UserCenterServer) GetGroupInfo(ctx context.Context, in *pb.GetGroupInfoReq) (*pb.GetGroupInfoResp, error) {
	l := logic.NewGetGroupInfoLogic(ctx, s.svcCtx)
	return l.GetGroupInfo(in)
}

func (s *UserCenterServer) GetGroupMemberUid(ctx context.Context, in *pb.GetGroupMemberUidReq) (*pb.GetGroupMemberUidResp, error) {
	l := logic.NewGetGroupMemberUidLogic(ctx, s.svcCtx)
	return l.GetGroupMemberUid(in)
}

func (s *UserCenterServer) GetMyFriends(ctx context.Context, in *pb.GetMyFriendsReq) (*pb.GetMyFriendsResp, error) {
	l := logic.NewGetMyFriendsLogic(ctx, s.svcCtx)
	return l.GetMyFriends(in)
}

func (s *UserCenterServer) GetMyFollows(ctx context.Context, in *pb.GetMyFollowsReq) (*pb.GetMyFollowsResp, error) {
	l := logic.NewGetMyFollowsLogic(ctx, s.svcCtx)
	return l.GetMyFollows(in)
}

func (s *UserCenterServer) GetMyFans(ctx context.Context, in *pb.GetMyFansReq) (*pb.GetMyFansResp, error) {
	l := logic.NewGetMyFansLogic(ctx, s.svcCtx)
	return l.GetMyFans(in)
}

func (s *UserCenterServer) GetMyFansUid(ctx context.Context, in *pb.GetMyFansUidReq) (*pb.GetMyFansUidResp, error) {
	l := logic.NewGetMyFansUidLogic(ctx, s.svcCtx)
	return l.GetMyFansUid(in)
}

func (s *UserCenterServer) GetMyGroups(ctx context.Context, in *pb.GetMyGroupsReq) (*pb.GetMyGroupsResp, error) {
	l := logic.NewGetMyGroupsLogic(ctx, s.svcCtx)
	return l.GetMyGroups(in)
}
