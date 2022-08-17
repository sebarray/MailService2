package handler

import (
	"log"
	"mailgo/routes"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Manager() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*", "*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	routes.Mail(e)

	PORT := os.Getenv("PORT")

	log.Fatal(e.Start(":" + PORT))
}
