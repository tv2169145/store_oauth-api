package rest

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/tv2169145/store_oauth-api/src/domain/users"
	"github.com/tv2169145/store_oauth-api/src/utils/errors"
	"gopkg.in/go-resty/resty.v2"
	"time"
)

var (
	usersRestClient = rest.RequestBuilder{
		BaseURL: "https://api.store.com",
		Timeout: 100 * time.Millisecond,
	}
	restyClient = resty.New()
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
}

type usersRepository struct {
	client *resty.Client
}

func init() {
	restyClient.SetTimeout(1 * time.Minute)
	restyClient.SetHostURL("http://localhost:8081")
}

func NewRepository() RestUsersRepository {
	return &usersRepository{}
}

func (r *usersRepository) LoginUser(email, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email: email,
		Password: password,
	}

	// 使用 "github.com/mercadolibre/golang-restclient/rest"
	//response := usersRestClient.Post("/users/login", request)
	//if response == nil || response.Response == nil  {
	//	return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	//}
	//
	//if response.StatusCode > 299 {
	//	var  restErr errors.RestErr
	//	err := json.Unmarshal(response.Bytes(), &restErr)
	//	if err != nil {
	//		return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
	//	}
	//	return nil, &restErr
	//}
	//var user users.User
	//if err := json.Unmarshal(response.Bytes(), &user); err != nil {
	//	return nil, errors.NewInternalServerError("error when trying to unmarshal response")
	//}
	//return &user, nil

	// 使用 resty

	successResult := users.User{}
	errorResult := errors.RestErr{}
	response, err := restyClient.R().SetBody(request).SetResult(&successResult).SetError(&errorResult).Post("/users/login")
	//fmt.Println(response, err, successResult, errorResult)
	if err != nil {
		return nil, errors.NewInternalServerError("invalid restclient response when trying to login user")
	}
	if response.StatusCode() > 299 {
		return nil, &errorResult
	}
	return &successResult, nil
}
