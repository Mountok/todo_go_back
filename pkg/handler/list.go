package handler

import (
	"net/http"
	todoapp "todo-app"

	"github.com/gin-gonic/gin"
)


func (h *Handler) createList(c *gin.Context) {

	userId, err := getUserId(c)
	if err != nil {
		return
	}

	var input todoapp.TodoList
	if err := c.Bind(&input); err != nil {
		newErrorResponse(c,http.StatusBadRequest,err.Error())
		return
	}

	listId, err := h.services.TodoList.Create(userId,input)
	if err != nil {
		newErrorResponse(c,http.StatusInternalServerError,err.Error())
		return

	}

	c.JSON(http.StatusOK,map[string]interface{}{
		"list_id": listId,
		"user_id": userId,
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