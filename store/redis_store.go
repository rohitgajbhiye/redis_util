package store

import (
	"context"
)

type RedisStore interface {
	Get(context.Context, string) (interface{}, error)
	Set(context.Context, string, interface{}) error
}

type RedisConfig struct {
}

type redisStore struct {
}

func NewRedisStore() redisStore {
	return redisStore{}
}
func (rs redisStore) Get(context.Context, string) (interface{}, error) {
	return map[string]string{
		"test": "test",
	}, nil
}

func (rs redisStore) Set(context.Context, string, interface{}) error {
	return nil
}
