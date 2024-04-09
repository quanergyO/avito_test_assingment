package postgres

import (
	"avito_test_assingment/types"
	"encoding/json"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"log/slog"
	"strings"
	"time"
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
	rows := "tag_ids, feature_id, content, is_active, created_at, updated_at"
	query := fmt.Sprintf("SELECT %s FROM banners WHERE 1=1", rows)
	if featureId != 0 {
		query += fmt.Sprintf(" AND feature_id = %d", featureId)
	}
	if len(tagsId) > 0 {
		query += fmt.Sprintf(" AND tag_ids @> ARRAY[%d", tagsId[0])
		for _, tagID := range tagsId[1:] {
			query += fmt.Sprintf(", %d", tagID)
		}
		query += "]"
	}
	if limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}
	if offset > 0 {
		query += fmt.Sprintf(" offset %d", offset)
	}

	slog.Info(query)
	bannersDTO := make([]types.BannerDTO, 0)
	if err := r.db.Select(&bannersDTO, query); err != nil {
		return nil, err
	}

	banners := make([]types.BannerGet200ResponseInner, len(bannersDTO))
	for i, banner := range bannersDTO {
		var contentMap map[string]interface{}
		err := json.Unmarshal(banner.Content, &contentMap)
		if err != nil {
			continue
		}
		banners[i].BannerId = banner.BannerId
		banners[i].TagIds = pqInt64ArrayToIntSlice(banner.TagIds)
		banners[i].FeatureId = banner.FeatureId
		banners[i].Content = contentMap
		banners[i].IsActive = banner.IsActive
		banners[i].CreatedAt = banner.CreatedAt
		banners[i].UpdatedAt = banner.UpdatedAt
	}

	return banners, nil
}

func (r *Banner) BannerIdDelete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", bannerTable)
	slog.Info(query)
	_, err := r.db.Exec(query, id)

	return err
}

func (r *Banner) BannerIdPatch(id int, data types.BannerIdPatchRequest) error {
	setValue := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if data.FeatureId != nil {
		setValue = append(setValue, fmt.Sprintf("feature_id=$%d", argId))
		args = append(args, *data.FeatureId)
		argId++
	}

	if data.TagIds != nil {
		setValue = append(setValue, fmt.Sprintf("tag_ids=$%d", argId))
		args = append(args, pq.Array(*data.TagIds))
		argId++
	}

	if data.Content != nil {
		content, err := json.Marshal(*data.Content)

		if err != nil {
			slog.Info(string(content))
			return err
		}
		setValue = append(setValue, fmt.Sprintf("content=$%d::jsonb", argId))
		args = append(args, content)
		argId++
	}

	if data.IsActive != nil {
		setValue = append(setValue, fmt.Sprintf("is_active=$%d", argId))
		args = append(args, *data.IsActive)
		argId++
	}

	now := time.Now()
	setValue = append(setValue, fmt.Sprintf("updated_at=$%d", argId))
	args = append(args, now)
	argId++

	args = append(args, id)
	setQuery := strings.Join(setValue, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE id=$%d", bannerTable, setQuery, argId)
	slog.Info(query)
	_, err := r.db.Exec(query, args...)

	return err
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
