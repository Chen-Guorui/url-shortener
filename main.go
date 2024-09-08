package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"url-shortener/config"
	"url-shortener/handler"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.POST("/shorten", handler.Shorten)
	r.GET("/redirect/:shortUrl", handler.Redirect)

	fmt.Printf("Serve at port: %d\n", config.Config.Port)
	r.Run(fmt.Sprintf("0.0.0.0:%d", config.Config.Port))
}
