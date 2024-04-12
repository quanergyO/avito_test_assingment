package types

type DeleteModelBannerInput struct {
	TagIds    []int `json:"tag_ids" binding:"required"`
	FeatureId int   `json:"feature_id" binding:"required"`
}
