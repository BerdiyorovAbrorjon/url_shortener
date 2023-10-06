package repository

import (
	"context"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository/pgstore"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository/rediscache"
)

type urlRepository struct {
	store *pgstore.PgStore
	cache *rediscache.RedisCache
}

func NewUrlRepository(pgStore *pgstore.PgStore, cache *rediscache.RedisCache) domain.UrlRepository {
	return &urlRepository{
		store: pgStore,
		cache: cache,
	}
}

func (ur *urlRepository) CreateUrl(ctx context.Context, shortUrl, orgUrl string, userId int32) (domain.Url, error) {
	return ur.store.CreateUrl(ctx, orgUrl, shortUrl, userId)
}

func (ur *urlRepository) GetOrgUrlByShort(ctx context.Context, shortUrl string) (string, error) {
	urlCache, err := ur.cache.GetUrl(ctx, shortUrl)
	if err != nil {
		dbUrl, err := ur.store.GetOrgUrlByShort(ctx, shortUrl)
		if err != nil {
			return "", err
		}
		ur.cache.CreateUrl(ctx, dbUrl)
		return dbUrl.OrgUrl, nil
	}

	err = ur.store.IncrementClick(ctx, urlCache.ID)
	if err != nil {
		return "", err
	}
	return urlCache.OrgUrl, err
}

func (ur *urlRepository) GetUrlById(ctx context.Context, id int64) (domain.Url, error) {
	return ur.store.GetUrlById(ctx, id)
}

func (ur *urlRepository) UpdateOrgUrl(ctx context.Context, id int64, newOrgUrl string) (domain.Url, error) {
	return ur.store.UpdateOrgUrl(ctx, id, newOrgUrl)
}

func (ur *urlRepository) DeleteUrl(ctx context.Context, id int64) error {
	return ur.store.DeleteUrl(ctx, id)
}
