package http

import (
	"github.com/gin-gonic/gin"
	"github.com/tv2169145/store_oauth-api/src/domain/access_token"
	"github.com/tv2169145/store_oauth-api/src/utils/errors"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}

type accessTokenHandler struct {
	service access_token.Service
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessTokenId := strings.TrimSpace(c.Param("access_token_id"))
	accessToken, err := handler.service.GetById(accessTokenId)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var token access_token.AccessToken
	if err := c.ShouldBindJSON(&token); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}
	if err := handler.service.Create(token); err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, token)
}
