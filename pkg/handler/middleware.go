package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)


const (
	authorizationHeader = "Authorization"
	userCtx = "userId"

)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized,"empty auth token in header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized,"invalid auth token in header")
	}

	id, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c,http.StatusUnauthorized,err.Error())
		return
	}

	c.Set(userCtx, id)
}

func getUserId(c *gin.Context) (int, error) {
	userIdAnyType, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c,http.StatusInternalServerError, "user ID not found")
		return 0, errors.New("user ID not found")
	}
	idInt, ok := userIdAnyType.(int)
	if !ok {
		newErrorResponse(c,http.StatusInternalServerError, "user ID is of invalid type")
		return 0, errors.New("user ID is of invalid type")
	}
	return idInt, nil
}