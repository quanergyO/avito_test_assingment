package types

type GetModelBannerInput struct {
	TagIds    []int `json:"tag_ids" binding:"required"`
	FeatureId int   `json:"feature_id" binding:"required"`
}
