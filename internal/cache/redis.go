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
	slog.Info("REDIS: Save banner in cache: ", data)

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

	slog.Info("REDIS: read banner from cache: ", banner)

	return banner, err
}

func (r *RedisCache) configureRedisKey(featureId int, tagIds []int) string {
	bannerKey := fmt.Sprintf("featureId=%d", featureId)
	for i, tag := range tagIds {
		bannerKey += fmt.Sprintf("tagId%d=%d", i+1, tag)
	}
	slog.Info(bannerKey)
	return bannerKey
}

func (r *RedisCache) getAllDataFromCache() ([]types.BannerGet200ResponseInner, error) {
	keys, err := r.cli.Keys(context.Background(), "*").Result()
	if err != nil {
		return nil, err
	}
	dataToWrite := make([]types.BannerGet200ResponseInner, 0)
	for _, key := range keys {
		var tmp types.BannerGet200ResponseInner
		err := r.cli.Get(context.Background(), key).Scan(&tmp)
		if err != nil {
			continue // TODO think
		}
		dataToWrite = append(dataToWrite, tmp)
	}

	return dataToWrite, nil
}

func (r *RedisCache) IsBannerExists(key string) (bool, error) {
	notExists, err := r.cli.Exists(context.Background(), key).Result()
	if err != nil {
		return false, err
	}
	return notExists == 1, nil
}
