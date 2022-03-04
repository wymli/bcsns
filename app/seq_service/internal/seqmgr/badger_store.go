package seqmgr

import (
	"fmt"
	"strconv"

	badger "github.com/dgraph-io/badger/v3"
)

type BadgerStore struct {
	BadgerClient *badger.DB
	Key          []byte
}

func (bs *BadgerStore) Read() (uint64, error) {
	var dst []byte
	err := bs.BadgerClient.View(func(txn *badger.Txn) error {
		item, err := txn.Get(bs.Key)
		if err != nil {
			return err
		}

		dst, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		return nil
	})

	switch err {
	case badger.ErrKeyNotFound:
		return 0, nil
	case nil:
		return strconv.ParseUint(string(dst), 10, 64)
	default:
		return 0, err
	}
}

func (bs *BadgerStore) Store(v uint64) error {
	err := bs.BadgerClient.Update(func(txn *badger.Txn) error {
		return txn.Set(bs.Key, []byte(fmt.Sprintf("%d", v)))
	})
	if err != nil {
		return err
	}
	return nil
}
