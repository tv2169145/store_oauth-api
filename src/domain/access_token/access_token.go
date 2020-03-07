package access_token

import (
	"fmt"
	"github.com/tv2169145/store_oauth-api/src/utils/errors"
	"github.com/tv2169145/store_users-api/utils/crypto_utils"
	"strings"
	"time"
)

const (
	expirationTime = 24
	grantTypePassword = "password"
	grandTypeClientCredentials = "client_credentials"
)

var (
	zone *time.Location
)

func init() {
	zone, _ = time.LoadLocation("Asia/Taipei")
}

type AccessTokenRequest struct {
	GrantType    string `json:"grant_type"`
	Scope string `json:"scope"`

	// 使用 password grant_type
	Username     string `json:"username"`
	Password     string `json:"password"`

	// 使用 client_credential grant_type
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func(r *AccessTokenRequest) Validate() *errors.RestErr {
	switch r.GrantType {
	case grantTypePassword:
		break
	case grandTypeClientCredentials:
		break
	default:
		return errors.NewInternalServerError("invalid grant_type")
	}
	//TODO Validate parameters for each grant_type
	return nil
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"` // web frontend - client-id: 123 or Android app - client-id:222
	Expires     int64  `json:"expires"`
}

func GetNewAccessToken(userId int64) AccessToken {
	return AccessToken{
		UserId: userId,
		Expires: time.Now().In(zone).Add(expirationTime * time.Hour).Unix(),
	}
}

func (at *AccessToken) Validate() *errors.RestErr {
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

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
