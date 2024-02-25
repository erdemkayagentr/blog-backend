package main

import (
	"erdemkayacomtr/config"
	"erdemkayacomtr/router"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.InitLog()
}

func main() {
	port := os.Getenv("PORT")
	init := config.Init()
	app := router.Init(init)
	app.Run(":" + port)
}
