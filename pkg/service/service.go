package service

import (
	"avito_test_assingment/pkg/repository"
	"avito_test_assingment/types"
)

type Authorization interface {
	CreateUser(user types.UserType) (int, error)
	CheckAuthData(username, password string) (types.UserType, error)
	GenerateToken(user types.UserType) (string, error)
	ParserToken(accessToken string) (*types.TokenClaims, error)
}

type Banner interface {
	BannerGet(token string, featureId int, tagId int, limit int, offset int) ([]types.BannerGet200ResponseInner, error)
	BannerIdDelete(id int, token string) error
	BannerIdPatch(id int, data types.BannerIdPatchRequest, token string) error
	BannerPost(data types.BannerPostRequest, token string) (int, error)
	UserBannerGet(tagId int, featureId int, useLastRevision bool, token string) (types.BannerGet200ResponseInner, error)
}

type Service struct {
	Authorization
	Banner
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		Banner:        NewBannerService(repos),
	}
}
