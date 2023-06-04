package main

import (
	"github.com/mrb-haqee/go-login-Oauth/controllers/api"
)

func main() {

	api := api.NewApi()
	api.Start()

}
