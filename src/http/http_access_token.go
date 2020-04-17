package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suvamsingh/bookstore_oauth-api/src/domain/access_token"
)

//AccessTokenHandler ...
type AccessTokenHandler interface {
	GetByID(*gin.Context)
}
type accessTokenHandler struct {
	service access_token.Service
}

//NewHandler ...
func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetByID(c *gin.Context) {
	accessTokenID := c.Param("access_token_id")
	accessToken, err := handler.service.GetByID(accessTokenID)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)

	c.JSON(http.StatusNotImplemented, "implement me now !!!")
}
