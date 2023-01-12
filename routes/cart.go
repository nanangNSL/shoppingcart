package routes

import (
	"API-CART/controllers"

	"github.com/gin-gonic/gin"
)

func CartRoutes(router *gin.Engine){
	router.GET("/", controllers.GetCart)
	router.GET("/cart", controllers.GetFilterProduct)
	router.POST("/", controllers.CreateCart)
	router.DELETE("/:id", controllers.DeleteCart)
}