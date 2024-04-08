package service

import (
	"avito_test_assingment/internal/repository"
	"avito_test_assingment/types"
)

type BannerService struct {
	repo repository.Banner
}

func NewBannerService(repo repository.Banner) *BannerService {
	return &BannerService{repo: repo}
}

func (s *BannerService) BannerGet(featureId int, tagId int, limit int, offset int) ([]types.BannerGet200ResponseInner, error) {
	return s.repo.BannerGet(featureId, tagId, limit, offset)
}

func (s *BannerService) BannerIdDelete(id int) error {
	return s.repo.BannerIdDelete(id)
}

func (s *BannerService) BannerIdPatch(id int, data types.BannerIdPatchRequest) error {
	return s.repo.BannerIdPatch(id, data)
}

func (s *BannerService) BannerPost(data types.BannerPostRequest) (int, error) {
	return s.repo.BannerPost(data)
}

func (s *BannerService) UserBannerGet(tagId int, featureId int, useLastRevision bool) (types.BannerGet200ResponseInner, error) {
	return s.repo.UserBannerGet(tagId, featureId, useLastRevision)
}
