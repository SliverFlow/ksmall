package biz

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

type EtcdLocker struct {
	logger *zap.Logger
	client *clientv3.Client
}

func NewEtcdLocker(logger *zap.Logger, client *clientv3.Client) *EtcdLocker {
	return &EtcdLocker{
		logger: logger,
		client: client,
	}
}

const prefix = "/ksmall/locker/"

func (l *EtcdLocker) Lock(ctx context.Context, key string) error {
	return nil
}

func (l *EtcdLocker) Unlock(ctx context.Context, key string) error {
	return nil
}
