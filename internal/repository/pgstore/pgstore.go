package pgstore

import (
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository/pgstore/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgStore struct {
	pool *pgxpool.Pool
	q    sqlc.Querier
}

func NewPgStore(pool *pgxpool.Pool) *PgStore {
	return &PgStore{
		pool: pool,
		q:    sqlc.New(pool),
	}
}
