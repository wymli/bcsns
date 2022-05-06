package pool

import (
	"errors"
	"math/rand"
	"strings"
)

type ConnPool struct {
	connMap map[string][]interface{}
}

type NewFunc func(endpoint string) (interface{}, error)

func MustNewConnPool(f NewFunc, c Config) ConnPool {
	connMap := make(map[string][]interface{}, len(c.Endpoints))
	for _, endpoint := range c.Endpoints {
		endpoint, err := standardizeAddr(endpoint)
		if err != nil {
			panic(err)
		}

		connList := make([]interface{}, c.Size)
		for i := 0; i < c.Size; i++ {
			conn, err := f(endpoint)
			if err != nil {
				panic(err)
			}
			connList[i] = conn
		}
		connMap[endpoint] = connList
	}
	return ConnPool{
		connMap: connMap,
	}
}

func standardizeAddr(endpoint string) (string, error) {
	s := strings.Split(endpoint, ":")
	switch len(s) {
	case 0:
		return "", errors.New("unrecognized address")
	case 1:
		return "localhost:" + s[0], nil
	case 2:
		ip := s[0]
		port := s[1]
		if ip == "127.0.0.1" {
			ip = "localhost"
		}
		return ip + ":" + port, nil
	default:
		return "", errors.New("unrecognized address")
	}
}

func (cp *ConnPool) Get(endpoint string) (interface{}, error) {
	addr, err := standardizeAddr(endpoint)
	if err != nil {
		return nil, err
	}
	connList, ok := cp.connMap[addr]
	if !ok {
		return nil, errors.New("Not found")
	}
	return connList[rand.Intn(len(connList))], nil
}
