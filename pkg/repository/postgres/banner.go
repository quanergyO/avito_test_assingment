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

func (r *Banner) BannerGet(featureId int, tagId int, limit int, offset int) ([]types.BannerGet200ResponseInner, error) {
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

func (r *Banner) UserBannerGet(tagId int, featureId int, useLastRevision bool) (types.BannerGet200ResponseInner, error) {
	return types.BannerGet200ResponseInner{}, fmt.Errorf("not implemented")
}
