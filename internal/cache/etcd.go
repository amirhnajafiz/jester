package cache

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type ETCD interface {
	Put(key string, value string) error
	Get(key string) (string, error)
}

func NewCache(cfg Config) (ETCD, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.Endpoints,
		DialTimeout: time.Duration(cfg.Timeout) * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return &cache{
		conn: cli,
	}, nil
}

type cache struct {
	conn *clientv3.Client
}

func (c cache) Put(key string, value string) error {
	return nil
}

func (c cache) Get(key string) (string, error) {
	return "", nil
}
