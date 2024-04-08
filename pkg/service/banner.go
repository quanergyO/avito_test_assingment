package service

import (
	"avito_test_assingment/pkg/repository"
	"avito_test_assingment/types"
	"log/slog"
)

type BannerService struct {
	repo repository.Banner
}

func NewBannerService(repo repository.Banner) *BannerService {
	return &BannerService{repo: repo}
}

func (s *BannerService) BannerGet(token string, featureId int, tagId int, limit int, offset int) ([]types.BannerGet200ResponseInner, error) {
	slog.Info("Token:", token)
	return s.repo.BannerGet(featureId, tagId, limit, offset)
}

func (s *BannerService) BannerIdDelete(id int, token string) error {
	slog.Info("Token:", token)
	return s.repo.BannerIdDelete(id)
}

func (s *BannerService) BannerIdPatch(id int, data types.BannerIdPatchRequest, token string) error {
	slog.Info("Token:", token)
	return s.repo.BannerIdPatch(id, data)
}

func (s *BannerService) BannerPost(data types.BannerPostRequest, token string) (int, error) {
	slog.Info("Token:", token)
	return s.repo.BannerPost(data)
}

func (s *BannerService) UserBannerGet(tagId int, featureId int, useLastRevision bool, token string) (types.BannerGet200ResponseInner, error) {
	slog.Info("Token:", token)
	return s.repo.UserBannerGet(tagId, featureId, useLastRevision)
}
