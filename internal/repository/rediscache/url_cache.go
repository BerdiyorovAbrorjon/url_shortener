package rediscache

import (
	"context"
	"encoding/json"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"
)

func (r *RedisCache) CreateUrl(ctx context.Context, url domain.Url) (string, error) {
	urlCache := domain.UrlCache{
		ID:       url.ID,
		OrgUrl:   url.OrgUrl,
		ShortUrl: url.ShortUrl,
	}

	val, err := json.Marshal(urlCache)
	if err != nil {
		return "", err
	}

	err = r.redisCl.Set(ctx, url.ShortUrl, val, r.expiration).Err()
	return url.ShortUrl, err
}

func (r *RedisCache) GetUrl(ctx context.Context, shortUrl string) (domain.UrlCache, error) {
	urlCache := domain.UrlCache{}

	val, err := r.redisCl.Get(ctx, shortUrl).Bytes()
	if err != nil {
		return urlCache, err
	}

	err = json.Unmarshal(val, &urlCache)
	return urlCache, err
}
