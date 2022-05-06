package logic

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/wymli/bcsns/app/bc_proxy/internal/svc"
	"github.com/wymli/bcsns/common/logx"
	bcsns "github.com/wymli/bcsns/smart_contract"
)

type EventWatchLogic struct {
	logx.Logger
	ctx      context.Context
	svcCtx   *svc.ServiceContext
	notifyCh chan interface{}
}

func NewEventWatchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EventWatchLogic {
	return &EventWatchLogic{
		Logger: logx.WithTraceCtx(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EventWatchLogic) Watch() {
	go l.watchMomentsPersistedEvent()
	go l.watchRoomMsgPersistedEvent()
	go l.watchMomentsPersistedEvent()
}

func (l *EventWatchLogic) watchUserMsgPersistedEvent() {
	msink := make(chan *bcsns.BcsnsUserMsgPersistedEvent, 100)
	subs, err := l.svcCtx.BcsnsClient.WatchUserMsgPersistedEvent(nil, msink, nil, nil)
	if err != nil {
		l.Fatal().Msgf("failed to watch user msg persisted event, err:%v", err)
	}

	for {
		select {
		case err := <-subs.Err():
			l.Fatal().Msgf("failed to subscribe an user msg persisted event stream, err:%v", err)
		case e := <-msink:
			l.notifyCh <- e
		}
	}
}

func (l *EventWatchLogic) watchRoomMsgPersistedEvent() {
	msink := make(chan *bcsns.BcsnsRoomMsgPersistedEvent, 100)
	subs, err := l.svcCtx.BcsnsClient.WatchRoomMsgPersistedEvent(nil, msink, nil)
	if err != nil {
		l.Fatal().Msgf("failed to watch room msg persisted event, err:%v", err)
	}

	for {
		select {
		case err := <-subs.Err():
			l.Fatal().Msgf("failed to subscribe an room msg persisted event stream, err:%v", err)
		case e := <-msink:
			l.notifyCh <- e
		}
	}
}

func (l *EventWatchLogic) watchMomentsPersistedEvent() {
	msink := make(chan *bcsns.BcsnsMomentsPersistedEvent, 100)
	subs, err := l.svcCtx.BcsnsClient.WatchMomentsPersistedEvent(nil, msink, nil)
	if err != nil {
		l.Fatal().Msgf("failed to watch room msg persisted event, err:%v", err)
	}

	for {
		select {
		case err := <-subs.Err():
			l.Fatal().Msgf("failed to subscribe an room msg persisted event stream, err:%v", err)
		case e := <-msink:
			l.notifyCh <- e
		}
	}
}

func (l *EventWatchLogic) GetAllRoomMsg(room_id uint64) {
	// bind.FilterOpts{Start: 0} means start block
	iter, err := l.svcCtx.BcsnsClient.FilterRoomMsgPersistedEvent(&bind.FilterOpts{Start: 0}, []uint64{room_id})
	if err != nil {
		l.Fatal().Msgf("failed to filter room msg persisted event, err:%v", err)
	}

	for iter.Next() {
		// iter.Event.Message
	}
}
