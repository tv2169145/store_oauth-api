package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime)
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "token should not be expired")
	assert.EqualValues(t, "", at.AccessToken, "new token should not have defined access token id")
	assert.EqualValues(t, 0, at.UserId, "new token should not have an user id")
}

func TestIsExpired(t *testing.T) {
	at := AccessToken{}
	assert.True(t, at.IsExpired(), "empty token should be expired by default")

	at.Expires = time.Now().In(zone).Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "token expiring 3 hours from now should not be expired")
}
