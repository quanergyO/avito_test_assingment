package service

import (
	"avito_test_assingment/internal/cache"
	"avito_test_assingment/internal/repository"
	"avito_test_assingment/types"
)

type Authorization interface {
	CreateUser(user types.UserType) (int, error)
	CheckAuthData(username, password string) (types.UserType, error)
	GenerateToken(user types.UserType) (string, error)
	ParserToken(accessToken string) (*types.TokenClaims, error)
}

type Banner interface {
	BannerGet(featureId int, tagId []int, limit int, offset int) ([]types.BannerGet200ResponseInner, error)
	BannerIdDelete(id int) error
	BannerIdPatch(id int, data types.BannerIdPatchRequest) error
	BannerPost(data types.BannerPostRequest) (int, error)
	UserBannerGet(tagId []int, featureId int, useLastRevision bool) (types.BannerGet200ResponseInner, error)
	DeleteBannerByFeatureAndTags(tagId []int, featureId int) error
}

type Service struct {
	Authorization
	Banner
}

func NewService(repos *repository.Repository, cacheInstance cache.Cache) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Banner:        NewBannerService(repos, cacheInstance),
	}
}
