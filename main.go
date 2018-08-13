package main

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/berto/shorty/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")

	app := api.SetupRouter()
	app.Run(":" + port)
}
