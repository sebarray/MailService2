package main

import (
	conf "mailgo/config"
	"mailgo/handler"
)

// GmailService : Gmail client for sending email

func main() {
	conf.Loadenv()
	handler.Manager()
}
