package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) createList(c *gin.Context) {

	userId, _ := c.Get(userCtx)
	c.JSON(http.StatusOK,map[string]interface{}{
		"id": userId,
	})

}
func (h *Handler) getAllLists(c *gin.Context) {
	
}
func (h *Handler) getListById(c *gin.Context) {
	
}
func (h *Handler) updateList(c *gin.Context) {
	
}
func (h *Handler) deleteList(c *gin.Context) {
	
}