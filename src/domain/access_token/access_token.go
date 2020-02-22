package access_token

import (
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
	Expires     int64  `json:"expire"`
}

func GetNewAccessToken() AccessToken {

	return AccessToken{
		Expires: time.Now().In(zone).Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	now := time.Now().In(zone)
	expiredTime := time.Unix(at.Expires, 0).In(zone)
	//return time.Unix(at.Expires, 0).Before(time.Now().In(zone))
	return now.After(expiredTime)
}
