package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/moskvinspace/go-auth/api/handlers"
	"github.com/moskvinspace/go-auth/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	api := r.Group("/api")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
		api.GET("/logout", handlers.Logout)
		api.GET("/current_user", handlers.CurrentUser)
	}

	docs.SwaggerInfo.Title = "moskvinspace/go-auth"
	docs.SwaggerInfo.Description = "Golang Authentication API with gin, postgres and JWT"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := r.Run(); err != nil {
		log.Fatal("Failed to run server. \n", err)
		os.Exit(2)
	} // listen and serve on 0.0.0.0:8080
}
