package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

  router.Static("/static", "./static")
	router.GET("/test", 	func(c *gin.Context) {
    c.JSON(200, gin.H{
      "message": "ok",
    })
  })

	router.Run(":8000")

}



