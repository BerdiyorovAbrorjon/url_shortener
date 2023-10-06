package service

import (
	"context"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/util"
)

type userService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) domain.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) ListUserUrls(ctx context.Context, userId, limit, offset int32) ([]domain.Url, error) {
	return us.userRepo.ListUserUrls(ctx, userId, limit, offset)
}

func (us *userService) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	return us.userRepo.GetByEmail(ctx, email)
}

func (us *userService) CreateUser(ctx context.Context, email, password, name string) (domain.User, error) {

	user := domain.User{}

	hashedPassword, err := util.HashPassword(password)
	if err != nil {
		return user, err
	}

	return us.userRepo.Create(ctx, email, hashedPassword, name)
}
