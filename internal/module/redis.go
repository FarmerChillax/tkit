package module

import (
	"context"

	"github.com/FarmerChillax/tkit/config"
	"github.com/FarmerChillax/tkit/pkg/helper"
	"github.com/redis/go-redis/v9"
)

type redisConn struct {
	client *redis.Client
}

func (rc *redisConn) Get(ctx context.Context) *redis.Client {
	return rc.client
}

func NewRedis(conf *config.RedisConfig) (*redisConn, error) {
	rdb, err := helper.NewRedis(conf)
	if err != nil {
		return nil, err
	}

	return &redisConn{
		client: rdb,
	}, nil
}
