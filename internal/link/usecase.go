package link

import (
	"context"
	"math/rand"
	"shorty_api/internal/model"
	"time"
)

type Usecase interface {
	Shorten(ctx context.Context, longURL string) (string, error)
	Resolve(ctx context.Context, longURL string) (string, error)
}

type usecase struct {
	repository Repository
}

func NewUsecase(repository Repository) Usecase {
	return &usecase{repository: repository}
}

func (u usecase) Shorten(ctx context.Context, longURL string) (string, error) {
	code := generateCode(6)

	linkRecord := &model.Link{
		Code: code,
		Url:  longURL,
	}
	if err := u.repository.SaveLink(ctx, linkRecord); err != nil {
		return "", err
	}
	return code, nil
}

func (u usecase) Resolve(ctx context.Context, longURL string) (string, error) {
	linkRecord, err := u.repository.GetLink(ctx, longURL)
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
