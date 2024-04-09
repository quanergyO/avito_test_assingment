package postgres

import (
	"avito_test_assingment/types"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"

	"log/slog"
)

type Banner struct {
	db *sqlx.DB
}

func NewBanner(db *sqlx.DB) *Banner {
	return &Banner{
		db: db,
	}
}

func (r *Banner) BannerGet(featureId int, tagsId []int, limit int, offset int) ([]types.BannerGet200ResponseInner, error) {
	return nil, fmt.Errorf("not implemented")
}

func (r *Banner) BannerIdDelete(id int) error {
	return fmt.Errorf("not implemented")
}

func (r *Banner) BannerIdPatch(id int, data types.BannerIdPatchRequest) error {
	return fmt.Errorf("not implemented")
}

func (r *Banner) BannerPost(data types.BannerPostRequest) (int, error) {
	content, err := json.Marshal(data.Content)

	if err != nil {
		slog.Info(string(content))
		return 0, err
	}
	var bannerId int
	query := fmt.Sprintf("INSERT INTO %s(tag_ids, feature_id, content, is_active) VALUES($1, $2, $3::jsonb, $4) RETURNING ID", bannerTable)
	slog.Info(query)
	row := r.db.QueryRow(query, pq.Array(data.TagIds), data.FeatureId, content, data.IsActive)
	if err := row.Scan(&bannerId); err != nil {
		return 0, err
	}

	return bannerId, nil
}

func (r *Banner) UserBannerGet(tagsId []int, featureId int) (types.BannerGet200ResponseInner, error) {
	slog.Info("Repository: UserBannerGet start")
	defer slog.Info("Repository: UserBannerGet end")

	var banner types.BannerDTO
	rows := "tag_ids, feature_id, content, is_active, created_at, updated_at"
	array := "{"
	for _, tag := range tagsId {
		array += fmt.Sprintf("%d,", tag)
	}
	array = array[:len(array)-1]
	array += "}"

	query := fmt.Sprintf("SELECT %s FROM %s WHERE feature_id = $1 AND tag_ids @> $2", rows, bannerTable)
	slog.Info(query)
	err := r.db.Get(&banner, query, featureId, array)
	if err != nil {
		return types.BannerGet200ResponseInner{}, err
	}

	var contentMap map[string]interface{}
	err = json.Unmarshal(banner.Content, &contentMap)
	if err != nil {
		return types.BannerGet200ResponseInner{}, err
	}

	return types.BannerGet200ResponseInner{
		BannerId:  banner.BannerId,
		TagIds:    pqInt64ArrayToIntSlice(banner.TagIds),
		FeatureId: banner.FeatureId,
		Content:   contentMap,
		IsActive:  banner.IsActive,
		CreatedAt: banner.CreatedAt,
		UpdatedAt: banner.UpdatedAt,
	}, nil
}

func pqInt64ArrayToIntSlice(arr pq.Int64Array) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = int(v)
	}
	return result
}
