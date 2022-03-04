package seqmgr

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

type FileStore struct {
	f *os.File
}

func NewFileStorer(filename string) (StoreKVer, error) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		return nil, err
	}

	return &FileStore{
		f: f,
	}, nil
}

func MustNewFileStorer(filename string) StoreKVer {
	s, err := NewFileStorer(filename)
	if err != nil {
		panic(err)
	}
	return s
}

func (fs *FileStore) Store(v uint64) error {
	if err := fs.f.Truncate(0); err != nil {
		return err
	}
	if _, err := fs.f.WriteAt([]byte(fmt.Sprintf("%d", v)), 0); err != nil {
		return err
	}
	return nil
}

func (fs *FileStore) Read() (uint64, error) {
	bs := make([]byte, 20)

	n, err := fs.f.ReadAt(bs, 0)
	switch err {
	case io.EOF:
		if n == 0 {
			return 0, nil
		}
		return strconv.ParseUint(string(bs[:n]), 10, 64)
	case nil:
		return strconv.ParseUint(string(bs[:n]), 10, 64)
	default:
		return 0, err
	}
}
