package cache

import (
	"context"
	"errors"
	"time"

	"github.com/amirhnajafiz/jester/pkg"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// ETCD is the key-value storage for
// tracing the published items
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
		return nil, pkg.WrapError("ETCD", "connection error", err)
	}

	return &cache{
		conn:   cli,
		timout: time.Duration(cfg.Timeout) * time.Second,
	}, nil
}

type cache struct {
	conn   *clientv3.Client
	timout time.Duration
}

func (c cache) Put(key string, value string) error {
	ctx, cancel := context.WithTimeout(context.Background(), c.timout)
	defer cancel()

	_, err := c.conn.Put(ctx, key, value)
	if err != nil {
		return pkg.WrapError("ETCD", "put error", err)
	}

	return nil
}

func (c cache) Get(key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timout)
	defer cancel()

	resp, err := c.conn.Get(ctx, key)
	if err != nil {
		return "", pkg.WrapError("ETCD", "get error", err)
	}

	for _, ev := range resp.Kvs {
		return string(ev.Value), nil
	}

	return "", pkg.WrapError("ETCD", "get error", errors.New("item not found"))
}
