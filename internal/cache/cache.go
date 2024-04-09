package cache

import (
	"avito_test_assingment/types"
)

type Cache interface {
	WriteBanner(data types.BannerGet200ResponseInner) error
	ReadBanner(input types.GetModelBannerInput) (types.BannerGet200ResponseInner, error)
}
