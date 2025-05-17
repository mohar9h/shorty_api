package link

import (
	"context"
	"time"

	"gorm.io/gorm"
	"shorty_api/internal/common/model"
)

type linkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) Repository {
	return &linkRepository{db: db}
}

func (r *linkRepository) SaveLink(ctx context.Context, link *model.Link) error {
	return r.db.WithContext(ctx).Create(link).Error
}

func (r *linkRepository) GetLink(ctx context.Context, code string) (*model.Link, error) {
	var link model.Link
	err := r.db.WithContext(ctx).
		Where("code = ?", code).
		First(&link).Error

	if err != nil {
		return nil, err
	}

	// check expire at -- expire after 3 hours
	if time.Now().After(link.ExpireAt) {
		return nil, gorm.ErrRecordNotFound
	}

	return &link, nil
}

func (r *linkRepository) CountTodayLinksByIP(ctx context.Context, ip string) (int64, error) {
	var count int64
	startOfDay := time.Now().Truncate(24 * time.Hour)

	err := r.db.WithContext(ctx).
		Model(&model.Link{}).
		Where("ip_address = ? AND created_at >= ?", ip, startOfDay).
		Count(&count).Error

	return count, err
}
