package handler

import (
	"todo-app/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New() // инициализируем роутер

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up",h.signUp)
		auth.POST("/sign-in",h.signIn)
	}

	api := router.Group("/api",h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/",h.createList)
			lists.GET("/",h.getAllLists)
			lists.GET("/:id",h.getListById)
			lists.PUT("/:id",h.updateItem)
			lists.DELETE("/:id",h.deleteItem)
		}
		items := lists.Group(":id/items")
		{
			items.POST("/",h.createList)
			items.GET("/",h.getAllLists)
			items.GET("/:id",h.getListById)
			items.PUT("/:id",h.updateList)
			items.DELETE("/:id",h.deleteList)
		}
	}
	return router
}
