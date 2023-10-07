package rediscache

import (
	"context"
	"testing"

	"github.com/BerdiyorovAbrorjon/url-shortener/internal/domain"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAndGetUrl(t *testing.T) {
	url := domain.Url{
		ID:       util.RandomNumber(100, 1000),
		ShortUrl: util.RandomString(7),
		OrgUrl:   util.RandomOrgUrl(),
	}

	key, err := testRedisCache.CreateUrl(context.Background(), url)
	require.NoError(t, err)
	require.Equal(t, key, url.ShortUrl)

	cache, err := testRedisCache.GetUrl(context.Background(), key)
	require.NoError(t, err)
	require.NotEmpty(t, cache)

	require.Equal(t, url.ID, cache.ID)
	require.Equal(t, url.ShortUrl, cache.ShortUrl)
	require.Equal(t, url.OrgUrl, cache.OrgUrl)

	cache, err = testRedisCache.GetUrl(context.Background(), util.RandomString(7))
	require.Error(t, err)
	require.Empty(t, cache)
}
