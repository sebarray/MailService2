package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"mailgo/model"
	"mailgo/service/mail"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Mail(c echo.Context) error {
	var Mail model.Mail
	body, err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		c.Error(err)
	}
	err = json.Unmarshal(body, &Mail)
	if err != nil {
		c.Error(err)
	}
	status, err := mail.SendEmailOAUTH2(os.Getenv("TO"), Mail, "template.txt")
	if err != nil {
		log.Println(err)
	}
	if status {
		log.Println("Email sent successfully using OAUTH")
	}
	return c.JSON(http.StatusOK, Mail)
}
