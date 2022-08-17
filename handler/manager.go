package handler

import (
	"log"
	"mailgo/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func Manager() {
	e := echo.New()

	routes.Mail(e)

	PORT := os.Getenv("PORT")

	log.Fatal(e.Start(":" + PORT))
}
