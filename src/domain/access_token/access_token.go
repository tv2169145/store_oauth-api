package access_token

import (
	"github.com/tv2169145/store_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

var (
	zone *time.Location
)

func init() {
	zone, _ = time.LoadLocation("Asia/Taipei")
}



type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"` // web frontend - client-id: 123 or Android app - client-id:222
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().In(zone).Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) Validate() *errors.RestErr {
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("invalid token id")
	}
	if at.UserId <= 0 {
		return errors.NewBadRequestError("invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("invalid expiration time")
	}
	return nil
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().In(zone)
	expiredTime := time.Unix(at.Expires, 0).In(zone)
	//return time.Unix(at.Expires, 0).Before(time.Now().In(zone))
	return now.After(expiredTime)
}
