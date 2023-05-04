package api

import (
	"github.com/gin-gonic/gin"
	"github.com/moskvinspace/simple-web-app/pkg/api/handlers"
	"log"
	"os"
)

func Start() {
	r := gin.Default()

	r.POST("/sign-up", handlers.SignUp)
	r.POST("/sign-in", handlers.SignIn)

	if err := r.Run(); err != nil {
		log.Fatal("Failed to run server. \n", err)
		os.Exit(2)
	} // listen and serve on 0.0.0.0:8080
}
