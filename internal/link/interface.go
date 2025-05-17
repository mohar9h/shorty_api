package link

import (
	"context"
	"shorty_api/internal/common/model"
)

type Repository interface {
	SaveLink(ctx context.Context, link *model.Link) error
	GetLink(ctx context.Context, id string) (*model.Link, error)
	CountTodayLinksByIP(ctx context.Context, ip string) (int64, error)
}

type Usecase interface {
	Shorten(ctx context.Context, longURL, ip string) (string, error)
	GetOriginalURL(ctx context.Context, code string) (string, error)
}
