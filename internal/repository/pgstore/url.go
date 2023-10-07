package pgstore

import (
	"context"
	"time"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository/pgstore/sqlc"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/mapper"
)

func (store *PgStore) CreateUrl(ctx context.Context, orgUrl, shortUrl string, userId int32) (domain.Url, error) {
	url := domain.Url{}

	dbUrl, err := store.querier(ctx).CreateUrl(ctx, sqlc.CreateUrlParams{
		OrgUrl:   orgUrl,
		ShortUrl: shortUrl,
		UserID:   userId,
	})
	if err != nil {
		return url, err
	}

	err = mapper.Map(dbUrl, &url)
	return url, err
}

func (store *PgStore) GetUrlById(ctx context.Context, id int64) (domain.Url, error) {
	url := domain.Url{}

	dbUrl, err := store.querier(ctx).GetUrlById(ctx, id)
	if err != nil {
		return url, err
	}

	err = mapper.Map(dbUrl, &url)
	return url, err
}

func (store *PgStore) GetOrgUrlByShort(ctx context.Context, shortUrl string) (domain.Url, error) {
	url := domain.Url{}

	dbUrl, err := store.querier(ctx).GetUrlByShort(ctx, shortUrl)
	if err != nil {
		return url, err
	}

	err = mapper.Map(dbUrl, &url)
	return url, err
}

func (store *PgStore) IncrementClick(ctx context.Context, id int64) error {
	return store.querier(ctx).IncrementClick(ctx, id)
}

func (store *PgStore) ListUserUrls(ctx context.Context, userId, limit, offset int32) ([]domain.Url, error) {
	urls := []domain.Url{}

	dbUrls, err := store.querier(ctx).ListUserUrls(ctx, sqlc.ListUserUrlsParams{
		UserID: userId,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return urls, nil
	}

	err = mapper.Map(dbUrls, &urls)
	return urls, err
}

func (store *PgStore) UpdateOrgUrl(ctx context.Context, id int64, newOrgUrl string) (domain.Url, error) {
	url := domain.Url{}

	dbUrl, err := store.querier(ctx).UpdateOrgUrl(ctx, sqlc.UpdateOrgUrlParams{
		ID:        id,
		OrgUrl:    newOrgUrl,
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return url, err
	}

	err = mapper.Map(dbUrl, &url)
	return url, err
}

func (store *PgStore) DeleteUrl(ctx context.Context, id int64) error {
	return store.querier(ctx).DeleteUrl(ctx, id)
}
