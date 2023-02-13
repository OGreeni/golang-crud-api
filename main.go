package main

import (
	"example.com/ogreeni/configs"
	"example.com/ogreeni/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(router)

	router.Run("0.0.0.0:3000")
}
