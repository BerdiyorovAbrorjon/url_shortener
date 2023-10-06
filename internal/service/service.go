package service

import "github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"

type Service struct {
	User domain.UserService
	Url  domain.UrlService
}
