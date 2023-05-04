package api

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

func Start() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	if err := r.Run(); err != nil {
		log.Fatal("Failed to run server. \n", err)
		os.Exit(2)
	} // listen and serve on 0.0.0.0:8080
}
