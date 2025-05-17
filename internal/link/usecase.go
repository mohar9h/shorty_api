package link

import (
	"context"
	"errors"
	"math/rand"
	"shorty_api/internal/common/model"
	"time"
)

const (
	linkExpireDuration = 3 * time.Hour
	maxLinksPerIPDaily = 5
)

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) Usecase {
	return &usecase{repository: repository}
}

func (u *usecase) Shorten(ctx context.Context, longURL, ip string) (string, error) {
	// every ip can make 5 short links in everyday
	count, err := u.repository.CountTodayLinksByIP(ctx, ip)
	if err != nil {
		return "", err
	}
	if count >= maxLinksPerIPDaily {
		return "", errors.New("rate limit exceeded: max 5 links per day per IP")
	}

	code := generateCode(6)

	linkRecord := &model.Link{
		Code:      code,
		Url:       longURL,
		IPAddress: ip,
		ExpireAt:  time.Now().Add(linkExpireDuration),
	}

	if err := u.repository.SaveLink(ctx, linkRecord); err != nil {
		return "", err
	}

	return code, nil
}

func (u *usecase) GetOriginalURL(ctx context.Context, code string) (string, error) {
	linkRecord, err := u.repository.GetLink(ctx, code)
	if err != nil {
		return "", err
	}

	return linkRecord.Url, nil
}

func generateCode(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
