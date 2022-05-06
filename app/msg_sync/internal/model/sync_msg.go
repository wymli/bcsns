package model

import (
	"fmt"

	"github.com/gocql/gocql"
)

type SyncMessage struct {
	Uid         uint64 // recv id
	SendUid     uint64 // send id
	RoomId      uint64 // room id
	ServerMsgId int64
	SendMsgId   int64
	ContentType string
	Data        []byte
}

func (sm *SyncMessage) Clear() {
	sm.Uid = 0
	sm.SendUid = 0
	sm.RoomId = 0
	sm.ServerMsgId = 0
	sm.SendMsgId = 0
	sm.ContentType = ""
	sm.Data = nil
}

func (sm *SyncMessage) Store(sess *gocql.Session, table string) error {
	if sm == nil {
		return fmt.Errorf("sync message model is nil")
	}

	return sess.Query(`INSERT INTO ? (uid, send_uid, room_id, server_msg_id, send_msg_id, content_type, data) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		table, sm.Uid, sm.SendUid, sm.RoomId, sm.ServerMsgId, sm.SendMsgId, sm.ContentType, sm.Data).Exec()
}

type SyncMessages []SyncMessage

func (sms SyncMessages) BatchStore(sess *gocql.Session, table string) error {
	batch := sess.NewBatch(gocql.UnloggedBatch)
	for _, sm := range sms {
		batch.Query(`INSERT INTO ? (uid, send_uid, room_id, server_msg_id, send_msg_id, content_type, data) VALUES (?, ?, ?, ?, ?, ?, ?)`,
			table, sm.Uid, sm.SendUid, sm.RoomId, sm.ServerMsgId, sm.SendMsgId, sm.ContentType, sm.Data)
	}
	return sess.ExecuteBatch(batch)
}

func QuerySyncMsgAll(sess *gocql.Session, table string, uid uint64) ([]SyncMessage, error) {
	sms := []SyncMessage{}
	iter := sess.Query("SELECT uid, send_uid, room_id, server_msg_id, send_msg_id, content_type, data FROM ? WHERE uid = ? ORDER BY server_msg_id DESC", table, uid).Iter()

	sm := SyncMessage{}
	for iter.Scan(&sm.Uid, &sm.SendUid, &sm.RoomId, &sm.ServerMsgId, &sm.SendMsgId, &sm.ContentType, &sm.Data) {
		sms = append(sms, sm)
		sm.Clear()
	}

	return sms, nil
}

func QuerySyncMsgLatest(sess *gocql.Session, table string, uid uint64, latestMsgId int64) ([]SyncMessage, error) {
	sms := []SyncMessage{}
	iter := sess.Query("SELECT uid, send_uid, room_id, server_msg_id, send_msg_id, content_type, data FROM ? WHERE uid = ? AND server_msg_id > ?", table, uid, latestMsgId).Iter()

	sm := SyncMessage{}
	for iter.Scan(&sm.Uid, &sm.SendUid, &sm.RoomId, &sm.ServerMsgId, &sm.SendMsgId, &sm.ContentType, &sm.Data) {
		sms = append(sms, sm)
		sm.Clear()
	}

	return sms, nil
}
