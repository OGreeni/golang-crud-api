package routes

import (
	"example.com/ogreeni/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/users", controllers.GetUsers())
}
