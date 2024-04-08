package cache

import (
	"avito_test_assingment/types"
	"context"
)

type Cache interface {
	WriteBanner(ctx context.Context, data types.BannerPostRequest) error
	ReadBanner(ctx context.Context, input types.GetModelBannerInput) (types.BannerGet200ResponseInner, error)
}
