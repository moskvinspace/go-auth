package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moskvinspace/go-auth/pkg/api/handlers"
	"log"
	"os"
)

func Start() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowAllOrigins:  true,
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Access-Control-Allow-Origin"},
	}))

	r.POST("/api/register", handlers.Register)
	r.POST("/api/login", handlers.Login)
	r.GET("/api/logout", handlers.Logout)
	r.GET("/api/current_user", handlers.CurrentUser)

	if err := r.Run(); err != nil {
		log.Fatal("Failed to run server. \n", err)
		os.Exit(2)
	} // listen and serve on 0.0.0.0:8080
}
