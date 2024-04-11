package cache

import (
	"avito_test_assingment/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
	"time"
)

type RedisCache struct {
	cli *redis.Client
}

const (
	TTLCache = 15 * time.Minute
)

type Config struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func NewRedis(cfg Config) (*RedisCache, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Host + ":" + cfg.Port,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	_, err := rdb.Ping(context.Background()).Result()

	return &RedisCache{cli: rdb}, err
}

func (r *RedisCache) WriteBanner(data types.BannerGet200ResponseInner) error {
	bannerKey := r.configureRedisKey(data.FeatureId, data.TagIds)
	err := r.cli.Set(context.Background(), bannerKey, data, TTLCache).Err()
	slog.Info("REDIS: Save banner in cache")

	return err
}

func (r *RedisCache) ReadBanner(input types.GetModelBannerInput) (types.BannerGet200ResponseInner, error) {
	var banner types.BannerGet200ResponseInner
	bannerKey := r.configureRedisKey(input.FeatureId, input.TagIds)
	res, err := r.cli.Get(context.Background(), bannerKey).Result()
	if err != nil {
		return types.BannerGet200ResponseInner{}, err
	}

	err = json.Unmarshal([]byte(res), &banner)
	if err != nil {
		return types.BannerGet200ResponseInner{}, err
	}

	slog.Info("REDIS: read banner from cache:")

	return banner, err
}

func (r *RedisCache) configureRedisKey(featureId int, tagIds []int) string {
	bannerKey := fmt.Sprintf("featureId=%d", featureId)
	for i, tag := range tagIds {
		bannerKey += fmt.Sprintf("tagId%d=%d", i+1, tag)
	}
	return bannerKey
}

func (r *RedisCache) IsBannerExists(key string) (bool, error) {
	notExists, err := r.cli.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	return notExists == 1, nil
}
