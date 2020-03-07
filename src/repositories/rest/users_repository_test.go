package rest

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	//rest.StartMockupServer()
	//httpmock.Activate()
	os.Exit(m.Run())
}

//func TestLoginUserTimeoutFromApi(t *testing.T) {
//	rest.FlushMockups()
//	rest.AddMockups(&rest.Mock{
//		URL: "https://api.store.com",
//		HTTPMethod: http.MethodPost,
//		ReqBody: `{"email":"email@gmail.com","password":"password"}`,
//		RespHTTPCode:-1,
//		RespBody:`{"message": "invalid login credentials", "status": "404", "error": "not_found"}`,
//	})
//	repository := usersRepository{}
//	user, err := repository.LoginUser("email@gmail.com", "password")
//	fmt.Println(user)
//	fmt.Println(err)
//
//	assert.Nil(t, user)
//	assert.NotNil(t, err)
//	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
//	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)
//}
//
//func TestLoginUserInvalidErrorInterface(t *testing.T) {
//	rest.FlushMockups()
//	rest.AddMockups(&rest.Mock{
//		URL: "http://api.store.com",
//		HTTPMethod: http.MethodPost,
//		ReqBody: `{"email":"email@gmail.com","password":"password"}`,
//		RespHTTPCode: http.StatusNotFound,
//		RespBody:`{"message": "invalid login credentials", "status": "404", "error": "not_found"}`,
//	})
//	repository := usersRepository{}
//	user, err := repository.LoginUser("email@gmail.com", "password")
//	fmt.Println(user)
//	fmt.Println(err)
//	assert.Nil(t, user)
//	assert.NotNil(t, err)
//	assert.EqualValues(t, http.StatusNotFound, err.Status)
//	assert.EqualValues(t, "invalid error interface when trying to login user", err.Message)
//}
//
//func TestLoginUserInvalidLoginCredentials(t *testing.T) {
//	rest.FlushMockups()
//	rest.AddMockups(&rest.Mock{
//		URL: "http://api.store.com",
//		HTTPMethod: http.MethodPost,
//		ReqBody: `{"email":"email@gmail.com","password":"password"}`,
//		RespHTTPCode: http.StatusNotFound,
//		RespBody:`{"message": "invalid login credentials", "status": "404", "error": "not_found"}`,
//	})
//	repository := usersRepository{}
//	user, err := repository.LoginUser("email@gmail.com", "password")
//	fmt.Println(user)
//	fmt.Println(err)
//	assert.Nil(t, user)
//	assert.NotNil(t, err)
//	assert.EqualValues(t, http.StatusNotFound, err.Status)
//	assert.EqualValues(t, "invalid login credentials", err.Message)
//}
//
//func TestLoginUserInvalidUserJsonResponse(t *testing.T) {
//	rest.FlushMockups()
//	rest.AddMockups(&rest.Mock{
//		URL: "http://api.store.com",
//		HTTPMethod: http.MethodPost,
//		ReqBody: `{"email":"email@gmail.com","password":"password"}`,
//		RespHTTPCode: http.StatusOK,
//		RespBody:`	{"id": "1", "first_name": "lin", "last_name": "moby", "email": "moby@gmail.com", "date_created": "2020-02-22 02:20:02", "status": "active", "password": "$2a$10$zLp5gpKBXnBol3NZw5U8Ke6qaeKy58ClxGae4UfdUlwTmEc4DPgO6"}`,
//	})
//	repository := usersRepository{}
//	user, err := repository.LoginUser("email@gmail.com", "password")
//	fmt.Println(user)
//	fmt.Println(err)
//	assert.Nil(t, user)
//	assert.NotNil(t, err)
//	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
//	assert.EqualValues(t, "error when trying to unmarshal response", err.Message)
//}
//
//func TestLoginUserNoError(t *testing.T) {
//	rest.FlushMockups()
//	rest.AddMockups(&rest.Mock{
//		URL: "http://api.store.com",
//		HTTPMethod: http.MethodPost,
//		ReqBody: `{"email":"email@gmail.com","password":"password"}`,
//		RespHTTPCode: http.StatusOK,
//		RespBody:`	{"id": 1, "first_name": "lin", "last_name": "moby", "email": "moby@gmail.com", "date_created": "2020-02-22 02:20:02", "status": "active", "password": "$2a$10$zLp5gpKBXnBol3NZw5U8Ke6qaeKy58ClxGae4UfdUlwTmEc4DPgO6"}`,
//	})
//	repository := usersRepository{}
//	user, err := repository.LoginUser("email@gmail.com", "password")
//	fmt.Println(user)
//	fmt.Println(err)
//	assert.Nil(t, user)
//	assert.NotNil(t, err)
//	assert.EqualValues(t, http.StatusOK, err.Status)
//}

// 使用 resty and httpmock
func TestUsersRepositoryUseRestyLoginUserServerError(t *testing.T) {
	defer httpmock.DeactivateAndReset()
	fixture := `{"message":"invalid restclient response when trying to login user","status": "500","error":"internal server error"}`
	responder := httpmock.NewStringResponder(200, fixture)
	fakeUrl := "https://api.store.com"
	httpmock.RegisterResponder("POST", fakeUrl, responder)
	repository := usersRepository{}
	//restyClient.SetHostURL("http://localhost:9999")
	res, err := repository.LoginUser("moby@gmail.com", "1234")

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, 500, err.Status)
	assert.EqualValues(t, "invalid restclient response when trying to login user", err.Message)
}

func TestUsersRepositoryUseRestyNotFoundUser(t *testing.T) {
	repository := usersRepository{}
	//restyClient.SetHostURL("http://localhost:8080")
	res, err := repository.LoginUser("oby@gmail.com", "1234")

	assert.Nil(t, res)
	assert.NotNil(t, err)
	assert.EqualValues(t, 404, err.Status)
	assert.EqualValues(t, "no user found in db", err.Message)
}

func TestUsersRepositoryLoginUserSuccess(t *testing.T) {
	repository := usersRepository{}
	//restyClient.SetHostURL("http://localhost:8080")
	res, err := repository.LoginUser("moby@gmail.com", "1234")

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.EqualValues(t, "moby@gmail.com", res.Email)
}

