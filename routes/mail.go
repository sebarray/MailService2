package routes

import (
	"mailgo/controller"

	"github.com/labstack/echo/v4"
)

func Mail(e *echo.Echo) {
	e.POST("/mail", controller.Mail)
}
