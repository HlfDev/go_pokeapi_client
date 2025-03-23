package bootstrap

import (
	"net/http"
	"os"
	"time"

	"github.com/HlfDev/go_pokeapi_client/internal/cache"
	"github.com/HlfDev/go_pokeapi_client/internal/handlers"
	"github.com/HlfDev/go_pokeapi_client/internal/logger"
	"github.com/HlfDev/go_pokeapi_client/internal/repositories"
	"github.com/HlfDev/go_pokeapi_client/internal/routes"
	"github.com/HlfDev/go_pokeapi_client/internal/services"

	// Import the docs
	_ "github.com/HlfDev/go_pokeapi_client/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Initialize sets up the application.
func Initialize() *gin.Engine {
	godotenv.Load()
	logger.Init()
	httpClient := &http.Client{Timeout: 10 * time.Second}
	c := cache.NewCache(10 * time.Minute)
	repo := repositories.NewPokemonRepository(httpClient, c)
	svc := services.NewPokemonService(repo)
	h := handlers.NewPokemonHandler(svc)
	r := gin.Default()
	r.Use(gin.Logger(), gin.Recovery())

	// Setup Swagger at http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupRoutes(r, h)
	logger.Log.Info("Server running on port 8080")
	if os.Getenv("TEST_MODE") != "true" {
		r.Run(":8080")
	}
	return r
}
