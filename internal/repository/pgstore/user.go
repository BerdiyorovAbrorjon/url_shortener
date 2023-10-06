package pgstore

import (
	"context"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository/pgstore/sqlc"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/mapper"
)

func (store *PgStore) CreateUser(ctx context.Context, email, hashedPassword, name string) (domain.User, error) {
	user := domain.User{}

	dbUser, err := store.querier(ctx).CreateUser(ctx, sqlc.CreateUserParams{
		Name:           name,
		HashedPassword: hashedPassword,
		Email:          email,
	})
	if err != nil {
		return user, err
	}

	err = mapper.Map(dbUser, &user)
	return user, err
}

func (store *PgStore) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	user := domain.User{}

	dbUser, err := store.querier(ctx).GetUserByEmail(ctx, email)
	if err != nil {
		return user, err
	}

	err = mapper.Map(dbUser, &user)
	return user, err
}
