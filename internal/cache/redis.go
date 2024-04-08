package cache

import (
	"avito_test_assingment/types"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	cli *redis.Client
}

//Addr:     "localhost:6379",
//Password: "1234",
//DB:       0,

func NewRedis()

func (r *RedisCache) WriteBanner(ctx context.Context, data types.BannerPostRequest) error {
	return fmt.Errorf("not implemented")
}

func (r *RedisCache) ReadBanner(ctx context.Context, input types.GetModelBannerInput) (types.BannerGet200ResponseInner, error) {
	return types.BannerGet200ResponseInner{}, fmt.Errorf("not implemented")
}
