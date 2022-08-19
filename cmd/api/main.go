package main

import (
	"log"
	"os"

	"github.com/ccallazans/twitter-video-downloader/internal/config"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

type Application struct {
	logger *log.Logger
	server *echo.Echo
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	config.RequestHeader.Set("authorization", os.Getenv("authorization"))
	config.RequestHeader.Set("x-guest-token", os.Getenv("x-guest-token"))

	logger := log.Default()

	app := Application{
		logger: logger,
	}

	app.NewRouter()

	err = app.server.Start(":5000")
	if err != nil {
		log.Fatalln(err)
	}
}