package repository

import (
	"context"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"
	"github.com/BerdiyorovAbrorjon/url-shortener/internal/repository/pgstore"
)

type userRepository struct {
	store *pgstore.PgStore
}

func NewUserRepository(store *pgstore.PgStore) domain.UserRepository {
	return &userRepository{
		store: store,
	}
}

func (ur *userRepository) Create(ctx context.Context, email, hashedPassword, name string) (domain.User, error) {
	return ur.store.CreateUser(ctx, email, hashedPassword, name)
}

func (ur *userRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	return ur.store.GetUserByEmail(ctx, email)
}

func (ur *userRepository) ListUserUrls(ctx context.Context, userId, limit, offset int32) ([]domain.Url, error) {
	return ur.store.ListUserUrls(ctx, userId, limit, offset)
}
