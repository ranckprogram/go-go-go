package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.StaticFS("/static", http.Dir("/static"))
	router.GET("/test", 	func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "ok",
    })
  })

	router.Run(":8000")

}



