package pgstore

import (
	"context"
	"testing"
	"time"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/hashing"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/util"
	"github.com/stretchr/testify/require"
)

func CreateRandomUrl(t *testing.T, user domain.User) domain.Url {
	orgUrl := util.RandomOrgUrl()
	shortUrl := hashing.GenerateShortURL(orgUrl, int(util.RandomNumber(6, 10)))

	url, err := testPgStore.CreateUrl(context.Background(), orgUrl, shortUrl, user.ID)
	require.NoError(t, err)
	require.NotEmpty(t, url)

	require.NotEmpty(t, url.ID)
	require.Equal(t, url.OrgUrl, orgUrl)
	require.Equal(t, url.ShortUrl, shortUrl)
	require.Equal(t, url.UserID, user.ID)
	require.Equal(t, url.Clicks, int64(0))

	url1, err1 := testPgStore.CreateUrl(context.Background(), orgUrl, shortUrl, user.ID)
	require.Error(t, err1)
	require.Empty(t, url1)

	return url
}

func TestCreateUrl(t *testing.T) {
	user := CreateRandomUser(t)
	CreateRandomUrl(t, user)
}

func TestGetUrlByID(t *testing.T) {
	user := CreateRandomUser(t)
	url := CreateRandomUrl(t, user)

	gUrl, err := testPgStore.GetUrlById(context.Background(), url.ID)
	require.NoError(t, err)
	require.NotEmpty(t, gUrl)

	require.Equal(t, url.ID, gUrl.ID)
	require.Equal(t, url.OrgUrl, gUrl.OrgUrl)
	require.Equal(t, url.ShortUrl, gUrl.ShortUrl)
	require.Equal(t, url.Clicks, gUrl.Clicks)
	require.Equal(t, url.UserID, gUrl.UserID)

	require.WithinDuration(t, url.CreatedAt, gUrl.CreatedAt, time.Second)
	require.WithinDuration(t, url.UpdatedAt, gUrl.UpdatedAt, time.Second)

	gUrl, err = testPgStore.GetUrlById(context.Background(), -1)
	require.Error(t, err)
	require.Empty(t, gUrl)
}

func TestGetOrgUrlByShort(t *testing.T) {
	user := CreateRandomUser(t)
	url := CreateRandomUrl(t, user)

	gUrl, err := testPgStore.GetOrgUrlByShort(context.Background(), url.ShortUrl)
	require.NoError(t, err)
	require.NotEmpty(t, gUrl)

	require.Equal(t, url.ID, gUrl.ID)
	require.Equal(t, url.OrgUrl, gUrl.OrgUrl)
	require.Equal(t, url.ShortUrl, gUrl.ShortUrl)
	require.Equal(t, url.Clicks, gUrl.Clicks)
	require.Equal(t, url.UserID, gUrl.UserID)

	require.WithinDuration(t, url.CreatedAt, gUrl.CreatedAt, time.Second)
	require.WithinDuration(t, url.UpdatedAt, gUrl.UpdatedAt, time.Second)

	gUrl, err = testPgStore.GetOrgUrlByShort(context.Background(), util.RandomString(8))
	require.Error(t, err)
	require.Empty(t, gUrl)
}

func TestIncrementClick(t *testing.T) {
	user := CreateRandomUser(t)
	url := CreateRandomUrl(t, user)

	err := testPgStore.IncrementClick(context.Background(), url.ID)
	require.NoError(t, err)

	updUrl, err := testPgStore.GetUrlById(context.Background(), url.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updUrl)

	require.Equal(t, url.Clicks+1, updUrl.Clicks)
}

func TestListUserUrls(t *testing.T) {
	user := CreateRandomUser(t)
	for i := 0; i < 5; i++ {
		CreateRandomUrl(t, user)
	}

	urls, err := testPgStore.ListUserUrls(context.Background(), user.ID, 5, 0)
	require.NoError(t, err)
	require.NotEmpty(t, urls)
	require.Equal(t, len(urls), 5)

	for _, u := range urls {
		require.NotEmpty(t, u)
		require.Equal(t, user.ID, u.UserID)
	}

	urls, err = testPgStore.ListUserUrls(context.Background(), user.ID, -1, 0)
	require.Empty(t, urls)

	urls, err = testPgStore.ListUserUrls(context.Background(), user.ID, 5, -1)
	require.Empty(t, urls)

	urls, err = testPgStore.ListUserUrls(context.Background(), -1, 5, 0)
	require.Empty(t, urls)
}

func TestUpdateOrgUrl(t *testing.T) {
	user := CreateRandomUser(t)
	url := CreateRandomUrl(t, user)
	newOrlUrl := util.RandomOrgUrl()

	updUrl, err := testPgStore.UpdateOrgUrl(context.Background(), -1, newOrlUrl)
	require.Error(t, err)
	require.Empty(t, updUrl)

	updUrl, err = testPgStore.UpdateOrgUrl(context.Background(), url.ID, newOrlUrl)
	require.NoError(t, err)
	require.NotEmpty(t, updUrl)

	require.Equal(t, url.ID, updUrl.ID)
	require.Equal(t, url.Clicks, updUrl.Clicks)
	require.Equal(t, url.ShortUrl, updUrl.ShortUrl)
	require.Equal(t, newOrlUrl, updUrl.OrgUrl)
	require.Equal(t, url.UserID, updUrl.UserID)

	require.WithinDuration(t, url.CreatedAt, updUrl.CreatedAt, time.Second)
	require.NotEqualValues(t, url.UpdatedAt, updUrl.UpdatedAt)
}

func TestDeleteUrl(t *testing.T) {
	user := CreateRandomUser(t)
	url := CreateRandomUrl(t, user)

	err := testPgStore.DeleteUrl(context.Background(), url.ID)
	require.NoError(t, err)

	gUrl, err := testPgStore.GetUrlById(context.Background(), url.ID)
	require.Error(t, err)
	require.Empty(t, gUrl)
}
