package repository

import (
	"avito_test_assingment/internal/repository/postgres"
	"avito_test_assingment/types"
	"github.com/jmoiron/sqlx"
)

type Banner interface {
	BannerGet(featureId int, tagId []int, limit int, offset int) ([]types.BannerGet200ResponseInner, error)
	BannerIdDelete(id int) error
	BannerIdPatch(id int, data types.BannerIdPatchRequest) error
	BannerPost(data types.BannerPostRequest) (int, error)
	UserBannerGet(tagId []int, featureId int) (types.BannerGet200ResponseInner, error)
}

type Authorization interface {
	CreateUser(user types.UserType) (int, error)
	GetUser(username, password string) (types.UserType, error)
}

type Repository struct {
	Authorization
	Banner
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuth(db),
		Banner:        postgres.NewBanner(db),
	}
}
