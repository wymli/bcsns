package svc

import (
	"github.com/wymli/bcsns/app/seq_service/internal/config"
	"github.com/wymli/bcsns/app/seq_service/internal/seqmgr"

	mylogx "github.com/wymli/bcsns/common/logx"
	zerologx "github.com/zeromicro/go-zero/core/logx"
)

type ServiceContext struct {
	Config     config.Config
	UserSeqMgr *seqmgr.UserSeqMgr
}

func NewServiceContext(c config.Config) *ServiceContext {
	zerologx.Disable()
	mylogx.Init(c.Logx)

	return &ServiceContext{
		Config:     c,
		UserSeqMgr: seqmgr.MustNewUserSeqMgr(c.Seq),
	}
}
