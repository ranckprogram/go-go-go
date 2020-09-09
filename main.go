package main

import (
	// "net/http"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	router := gin.New()
	router.Use(TlsHandler())

	// router.StaticFS("/static", http.Dir("/static"))
	router.Static("/", "./static")
	// router.GET("/persons", 	func(c *gin.Context) {
  //   c.JSON(200, gin.H{
  //     "message": "pong",
  //   })
  // })

	// router.Run(":8000")
	router.RunTLS(":8080", "ca/tls_cert.crt", "ca/tls_key.key")

}

func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
			secureMiddleware := secure.New(secure.Options{
					SSLRedirect: true,
					SSLHost:     "localhost:8080",
			})
			err := secureMiddleware.Process(c.Writer, c.Request)

			// If there was an error, do not continue.
			if err != nil {
					return
			}

			c.Next()
	}
}

