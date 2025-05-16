package link

import (
	"context"
	"gorm.io/gorm"
	"shorty_api/internal/model"
)

type Repository interface {
	SaveLink(ctx context.Context, link *model.Link) error
	GetLink(ctx context.Context, id string) (*model.Link, error)
}

type linkRepository struct {
	db *gorm.DB
}

func NewLinkRepository(db *gorm.DB) Repository {
	return &linkRepository{db: db}
}

func (l linkRepository) SaveLink(ctx context.Context, link *model.Link) error {
	return l.db.WithContext(ctx).Create(link).Error
}

func (l linkRepository) GetLink(ctx context.Context, code string) (*model.Link, error) {
	var link model.Link

	err := l.db.WithContext(ctx).Where("code  = ?", code).First(&link).Error
	if err != nil {
		return nil, err
	}

	return &link, nil
}
