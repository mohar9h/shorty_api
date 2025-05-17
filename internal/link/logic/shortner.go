package logic

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"net/url"
	"strings"
)

func GenerateCode(length int) (string, error) {
	b := make([]byte, length)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	code := base64.URLEncoding.EncodeToString(b)
	code = strings.ReplaceAll(code, "-", "")
	code = strings.ReplaceAll(code, "_", "")
	return code[:length], nil
}

func ValidateURL(raw string) error {
	parsed, err := url.ParseRequestURI(raw)
	if err != nil || parsed.Scheme == "" || parsed.Host == "" {
		return errors.New("invalid URL format")
	}
	return nil
}
