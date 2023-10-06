package pgstore

import (
	"context"
	"fmt"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository/pgstore/sqlc"
)

var txContextKey = struct{}{}

func (s *PgStore) WithTX(ctx context.Context, fn func() error) error {
	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}

	txQuery := sqlc.New(tx)
	vCtx := context.WithValue(ctx, txContextKey, txQuery)

	err = fn()
	if err != nil {
		if rbErr := tx.Rollback(vCtx); rbErr != nil {
			return fmt.Errorf("tx error: %w; tx roolback error: %w", err, rbErr)
		}
		return fmt.Errorf("tx error: %w", err)
	}

	return tx.Commit(vCtx)
}

func (s *PgStore) querier(ctx context.Context) sqlc.Querier {
	querier, ok := ctx.Value(txContextKey).(sqlc.Querier)
	if ok {
		return querier
	}
	return s.q
}
