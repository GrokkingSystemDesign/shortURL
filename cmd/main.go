package main

import (
	"fmt"
	"github.com/GrokkingSystemDesign/shortURL/config"
	"github.com/GrokkingSystemDesign/shortURL/service"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	err := config.LoadConf()
	if err != nil {
		log.Fatalf("LoadTConf error|%v", err)
		return
	}
	err = config.InitDB()
	if err != nil {
		fmt.Println(err)
		log.Fatalf("InitDB db error|%v", err)
		return
	}
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.POST("/", service.HandleURLShorten)
	router.GET("/:url", service.HandleRedirect)
	router.Run(":8080")
}
