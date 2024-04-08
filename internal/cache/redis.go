package cache

import (
	"avito_test_assingment/types"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log/slog"
)

type RedisCache struct {
	cli *redis.Client
}

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

func (r *RedisCache) WriteBanner(ctx context.Context, data types.BannerPostRequest) error {
	bannerKey := r.configureRedisKey(data.FeatureId, data.TagIds)
	jsonString, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = r.cli.Set(context.Background(), bannerKey, jsonString, 0).Err()
	return err
}

func (r *RedisCache) ReadBanner(ctx context.Context, input types.GetModelBannerInput) (types.BannerGet200ResponseInner, error) {
	var banner types.BannerGet200ResponseInner
	bannerKey := r.configureRedisKey(input.FeatureId, input.TagIds)
	val, err := r.cli.Get(context.Background(), bannerKey).Result()
	if err != nil {
		return banner, err
	}
	if err = json.Unmarshal([]byte(val), &banner); err != nil {
		return banner, err
	}

	return banner, nil
}

func (r *RedisCache) configureRedisKey(featureId int, tagIds []int) string {
	bannerKey := fmt.Sprintf("featureId=%d", featureId)
	for i, tag := range tagIds {
		bannerKey += fmt.Sprintf("tagId%d=%d", i+1, tag)
	}
	slog.Info(bannerKey)
	return bannerKey
}
