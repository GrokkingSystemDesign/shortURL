package main

import (
	"github.com/GrokkingSystemDesign/shortURL/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.POST("/", service.HandleURLShorten)
	router.GET("/:url", service.HandleRedirect)
	router.Run(":8080")
}
