package service

import (
	"context"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository/pgstore"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/hashing"
)

type urlService struct {
	urlRepo domain.UrlRepository
}

func NewUrlService(urlRepo domain.UrlRepository) domain.UrlService {
	return &urlService{
		urlRepo: urlRepo,
	}
}

func (us *urlService) CreateUrl(ctx context.Context, orgUrl string, userId int32) (domain.Url, error) {
	len := 6
	for {
		shortUrl := hashing.GenerateShortURL(orgUrl, len)
		url, err := us.urlRepo.CreateUrl(ctx, shortUrl, orgUrl, userId)
		if err == nil {
			return url, nil
		}
		if pgstore.ErrorCode(err) != pgstore.UniqueViolation {
			return url, err
		}
		len++
	}
}

func (us *urlService) GetOrgUrlByShort(ctx context.Context, shortUrl string) (string, error) {
	return us.urlRepo.GetOrgUrlByShort(ctx, shortUrl)
}

func (us *urlService) GetUrlById(ctx context.Context, id int64) (domain.Url, error) {
	return us.urlRepo.GetUrlById(ctx, id)
}

func (us *urlService) UpdateOrgUrl(ctx context.Context, id int64, newOrgUrl string) (domain.Url, error) {
	return us.urlRepo.UpdateOrgUrl(ctx, id, newOrgUrl)
}

func (us *urlService) DeleteUrl(ctx context.Context, id int64) error {
	return us.urlRepo.DeleteUrl(ctx, id)
}
