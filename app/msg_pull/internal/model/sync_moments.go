package model

import (
	"fmt"

	"github.com/gocql/gocql"
)

type SyncMoment struct {
	Uid         uint64 // recv id
	SendUid     uint64 // send id
	ServerMsgId int64
	ContentType string
	Data        []byte
}

func (sm *SyncMoment) Clear() {
	sm.Uid = 0
	sm.SendUid = 0
	sm.ServerMsgId = 0
	sm.ContentType = ""
	sm.Data = nil
}

func (sm *SyncMoment) Store(sess *gocql.Session, table string) error {
	if sm == nil {
		return fmt.Errorf("sync message model is nil")
	}

	return sess.Query(`INSERT INTO ? (uid, send_uid, server_msg_id, content_type, data) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		table, sm.Uid, sm.SendUid, sm.ServerMsgId, sm.ContentType, sm.Data).Exec()
}

type SyncMoments []SyncMoment

func (sms SyncMoments) BatchStore(sess *gocql.Session, table string) error {
	batch := sess.NewBatch(gocql.UnloggedBatch)
	for _, sm := range sms {
		batch.Query(`INSERT INTO ? (uid, send_uid, server_msg_id, content_type, data) VALUES (?, ?, ?, ?, ?, ?, ?)`,
			table, sm.Uid, sm.SendUid, sm.ServerMsgId, sm.ContentType, sm.Data)
	}
	return sess.ExecuteBatch(batch)
}

func QuerySyncMomentsAll(sess *gocql.Session, table string, uid uint64) ([]SyncMoment, error) {
	sms := []SyncMoment{}
	iter := sess.Query("SELECT uid, send_uid, server_msg_id, content_type, data FROM ? WHERE uid = ? ORDER BY server_msg_id DESC", table, uid).Iter()

	sm := SyncMoment{}
	for iter.Scan(&sm.Uid, &sm.SendUid, &sm.ServerMsgId, &sm.ContentType, &sm.Data) {
		sms = append(sms, sm)
		sm.Clear()
	}

	return sms, nil
}

func QuerySyncMomentsLatest(sess *gocql.Session, table string, uid uint64, latestMsgId int64) ([]SyncMoment, error) {
	sms := []SyncMoment{}
	iter := sess.Query("SELECT uid, send_uid, server_msg_id, content_type, data FROM ? WHERE uid = ? AND server_msg_id > ?", table, uid, latestMsgId).Iter()

	sm := SyncMoment{}
	for iter.Scan(&sm.Uid, &sm.SendUid, &sm.ServerMsgId, &sm.ContentType, &sm.Data) {
		sms = append(sms, sm)
		sm.Clear()
	}

	return sms, nil
}
