package seqmgr

import (
	"fmt"

	"github.com/dgraph-io/badger/v3"
)

// todo: 可以使用mmap申请大内存自己维护,而不是用badger
type UserSeqMgr struct {
	Partitions []SeqGenerator
	config     SeqGenConfig
}

func NewUserSeqMgr(config SeqGenConfig) (*UserSeqMgr, error) {
	db, err := badger.Open(badger.DefaultOptions(config.FileNameOptional))
	if err != nil {
		return nil, fmt.Errorf("failed to open badger, err:%v", err)
	}

	ptr := make([]SeqGenerator, config.UserPartitionSize)
	for i := 0; i < len(ptr); i++ {
		if err := ptr[i].Init(&BadgerStore{
			BadgerClient: db,
			Key:          []byte(fmt.Sprintf("%d", i)),
		}, config.Step); err != nil {
			return nil, fmt.Errorf("failed to init seqGenerator, err:%v", err)
		}
	}

	return &UserSeqMgr{
		Partitions: ptr,
	}, nil
}

func MustNewUserSeqMgr(config SeqGenConfig) *UserSeqMgr {
	usm, err := NewUserSeqMgr(config)
	if err != nil {
		panic(err)
	}
	return usm
}

func (usm *UserSeqMgr) GetNextId(userId uint64) (uint64, error) {
	partitionId := userId % uint64(len(usm.Partitions))
	return usm.Partitions[partitionId].GetNextId(usm.config.Step)
}
