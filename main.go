package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	log.Printf("starting the r0llm0ps server at http://localhost:9999 ...")

	r := gin.New()

	r.LoadHTMLGlob("html/*.html")

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Foo": "Hi Foo",
		})
	})

	// r.GET("/api/boxes", func(c *gin.Context) {

	// 	c.JSON(http.StatusOK, gin.H{
	// 		"boxes": Boxes,
	// 	})
	// })

	// r.POST("/api/boxes/stop", func(c *gin.Context) {

	// })

	r.Run(":9999")
}
