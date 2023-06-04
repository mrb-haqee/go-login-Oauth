package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mrb-haqee/go-login-Oauth/controllers/controller"
)

func NewRouter() *gin.Engine {
	var c = controller.NewController()
	r := gin.Default()
	r.LoadHTMLGlob("view/*")

	r.GET("/", c.HandleMain)
	r.GET("/login", c.HandleGoogleLogin)
	r.GET("/callback", c.HandleGoogleCallback)
	// r.GET("/LoginWithOauth")
	return r
}
