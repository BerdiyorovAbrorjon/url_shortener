package pgstore

import (
	"context"
	"testing"
	"time"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUser(t *testing.T) domain.User {
	hashedPass, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)
	email := util.RandomEmail()
	name := util.RandomName()

	user, err := testPgStore.CreateUser(context.Background(), email, hashedPass, name)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.NotEmpty(t, user.ID)
	require.Equal(t, user.Email, email)
	require.Equal(t, user.HashedPassword, hashedPass)
	require.Equal(t, user.Name, name)

	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	user1, err1 := testPgStore.CreateUser(context.Background(), email, hashedPass, name)
	require.Error(t, err1)
	require.Empty(t, user1)

	return user
}

func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

func TestGetUserByEmail(t *testing.T) {
	user := CreateRandomUser(t)

	gUser, err := testPgStore.GetUserByEmail(context.Background(), user.Email)
	require.NoError(t, err)
	require.NotEmpty(t, gUser)

	require.Equal(t, user.Email, gUser.Email)
	require.Equal(t, user.Name, gUser.Name)
	require.Equal(t, user.HashedPassword, gUser.HashedPassword)
	require.Equal(t, user.ID, gUser.ID)

	require.WithinDuration(t, user.CreatedAt, gUser.CreatedAt, time.Second)
	require.WithinDuration(t, user.UpdatedAt, gUser.UpdatedAt, time.Second)

	gUser, err = testPgStore.GetUserByEmail(context.Background(), util.RandomEmail())
	require.Error(t, err)
	require.Empty(t, gUser)
}
