package service

import (
	"avito_test_assingment/internal/cache"
	"avito_test_assingment/internal/repository"
	"avito_test_assingment/types"
)

type BannerService struct {
	repo  *repository.Repository
	cache cache.Cache
}

func NewBannerService(repo *repository.Repository, cacheInstance cache.Cache) *BannerService {
	return &BannerService{
		repo:  repo,
		cache: cacheInstance,
	}
}

func (s *BannerService) BannerGet(featureId int, tagId []int, limit int, offset int) ([]types.BannerGet200ResponseInner, error) {
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

func (s *BannerService) UserBannerGet(tagId []int, featureId int, useLastRevision bool) (types.BannerGet200ResponseInner, error) {
	if useLastRevision {
		data, err := s.cache.ReadBanner(types.GetModelBannerInput{
			TagIds:    tagId,
			FeatureId: featureId,
		})
		if err == nil {
			return data, nil
		}
	}

	data, err := s.repo.UserBannerGet(tagId, featureId)
	if err != nil {
		return data, err
	}

	err = s.cache.WriteBanner(data)

	return data, err
}

func (s *BannerService) DeleteBannerByFeatureAndTags(tagId []int, featureId int) error {
	return s.repo.DeleteBannerByFeatureAndTags(tagId, featureId)
}
