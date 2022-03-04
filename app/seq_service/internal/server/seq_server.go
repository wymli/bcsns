// Code generated by goctl. DO NOT EDIT!
// Source: seq.proto

package server

import (
	"context"

	"github.com/wymli/bcsns/app/seq_service/internal/logic"
	"github.com/wymli/bcsns/app/seq_service/internal/svc"
	"github.com/wymli/bcsns/app/seq_service/pb"
)

type SeqServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedSeqServer
}

func NewSeqServer(svcCtx *svc.ServiceContext) *SeqServer {
	return &SeqServer{
		svcCtx: svcCtx,
	}
}

func (s *SeqServer) GetSeqId(ctx context.Context, in *pb.GetSeqIdReq) (*pb.GetSeqIdResp, error) {
	l := logic.NewGetSeqIdLogic(ctx, s.svcCtx)
	return l.GetSeqId(in)
}