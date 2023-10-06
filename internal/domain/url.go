package domain

import (
	"context"
	"time"
)

// db model
type Url struct {
	ID        int64     `json:"id"`
	UserID    int32     `json:"user_id"`
	OrgUrl    string    `json:"org_url"`
	ShortUrl  string    `json:"short_url"`
	Clicks    int64     `json:"clicks"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// cache model
type UrlCache struct {
	ID       int64  `json:"id"`
	OrgUrl   string `json:"org_url"`
	ShortUrl string `json:"short_url"`
}

type CreateUrlRequest struct {
	OrgUrl string `json:"org_url" binding:"required"`
}

type GetUrlByIdRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type RederectRequest struct {
	ShortUrl string `uri:"short_url"`
}

type UpdateOrgUrlRequest struct {
	ID        int64  `json:"id" binding:"required,min=1"`
	NewOrgUrl string `json:"new_org_url" binding:"required"`
}

type DeleteUrlRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type UrlRepository interface {
	CreateUrl(ctx context.Context, shortUrl, orgUrl string, userId int32) (Url, error)
	GetOrgUrlByShort(ctx context.Context, shortUrl string) (string, error)
	GetUrlById(ctx context.Context, id int64) (Url, error)
	UpdateOrgUrl(ctx context.Context, id int64, newOrgUrl string) (Url, error)
	DeleteUrl(ctx context.Context, id int64) error
}

type UrlService interface {
	CreateUrl(ctx context.Context, orgUrl string, userId int32) (Url, error)
	GetOrgUrlByShort(ctx context.Context, shortUrl string) (string, error)
	GetUrlById(ctx context.Context, id int64) (Url, error)
	UpdateOrgUrl(ctx context.Context, id int64, newOrgUrl string) (Url, error)
	DeleteUrl(ctx context.Context, id int64) error
}
