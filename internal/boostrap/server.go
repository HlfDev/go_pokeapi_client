package boostrap

import (
	"fmt"
	"log"
	"pokeapi/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func Initialize() {

	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}

	e := gin.Default()
	routes(e)

	err := e.Run(":8080")

	if err != nil {
		panic(err)
	}

	fmt.Println("Server Started")

}

func routes(e *gin.Engine) {
	g := e.Group("api/v1")
	{
		g.GET("/pokemon/random", handlers.GetRandomPokemon)
	}
}
