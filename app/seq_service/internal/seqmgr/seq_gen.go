package seqmgr

import (
	"fmt"
	"math"
	"sync/atomic"
)

type StoreKVer interface {
	Store(v uint64) error
	Read() (uint64, error) // 注意在实现中,NotFound,EOF等错误都要返回 (0,nil)
}

type SeqGenConfig struct {
	Step              uint64
	FileNameOptional  string
	UserPartitionSize uint64
}

type SeqGenerator struct {
	curSeqId   uint64    // 当前id
	PhaseMaxId uint64    // 下一阶段起始id
	StoreProxy StoreKVer // 持久化存储接口
}

func NewSeqGenerator(storer StoreKVer, step uint64) (*SeqGenerator, error) {
	sg := &SeqGenerator{}
	if err := sg.Init(storer, step); err != nil {
		return nil, err
	}

	return sg, nil
}

func MustNewSeqGenerator(storer StoreKVer, step uint64) *SeqGenerator {
	sg, err := NewSeqGenerator(storer, step)
	if err != nil {
		panic(err)
	}
	return sg
}

func (sg *SeqGenerator) Init(storer StoreKVer, step uint64) error {
	sg.StoreProxy = storer

	lastPersistSeqId, err := sg.StoreProxy.Read()
	if err != nil {
		return fmt.Errorf("failed to get last persistted next-phase seq id, err:%v", err)
	}

	sg.curSeqId = lastPersistSeqId
	sg.PhaseMaxId = lastPersistSeqId + step

	return sg.Persist()
}

func (sg *SeqGenerator) Persist() error {
	return sg.StoreProxy.Store(sg.PhaseMaxId)
}

func (sg *SeqGenerator) GetNextId(step uint64) (uint64, error) {
	if err := sg.maybeNextPhase(step); err != nil {
		return 0, err
	}

	return atomic.AddUint64(&sg.curSeqId, 1), nil
}

func (sg *SeqGenerator) maybeNextPhase(step uint64) error {
	swapped := atomic.CompareAndSwapUint64(&sg.PhaseMaxId, sg.curSeqId+1, sg.PhaseMaxId+step)
	if !swapped {
		return nil
	}

	err := sg.Persist()
	if err != nil {
		return err
	}

	return sg.maybeResetPhase(step)
}

// todo: 目前 reset 的时候,仍然会命中获取id的逻辑,只要reset足够快,或phaseStep足够大,reset期间没有分配超过phaseStep的id即可,否则会出现严重错误
// todo: 比如phaseMaxId 永远得不到更新
func (sg *SeqGenerator) maybeResetPhase(step uint64) error {
	if sg.PhaseMaxId > uint64(math.MaxUint64-step) {
		return sg.reset(step)
	}
	return nil
}

func (sg *SeqGenerator) reset(step uint64) error {
	sg.curSeqId = 0
	sg.PhaseMaxId = step
	return sg.Persist()
}
