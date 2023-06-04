package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mrb-haqee/go-login-Oauth/model"
	"golang.org/x/oauth2"
)

type Controller struct {
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) HandleMain(lo *gin.Context) {
	// var htmlIndex = `
	// <html>
	// 	<body>
	// 		<a href="/login">Google Log In</a>
	// 	</body>
	// </html>`

	lo.HTML(http.StatusOK, "login.html", gin.H{"login": "/login"})
}

func (c *Controller) HandleGoogleLogin(lo *gin.Context) {
	url := model.GoogleOauthConfig.AuthCodeURL(model.OauthStateString)
	lo.Redirect(http.StatusTemporaryRedirect, url)
}

func (c *Controller) HandleGoogleCallback(lo *gin.Context) {
	state := lo.Request.FormValue("state")
	// state := r.FormValue("state")
	if state != model.OauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", model.OauthStateString, state)
		lo.Redirect(http.StatusTemporaryRedirect, "/")
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	code := lo.Request.FormValue("code")
	// code := r.FormValue("code")
	token, err := model.GoogleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("Code exchange failed with '%s'\n", err.Error())
		lo.Redirect(http.StatusTemporaryRedirect, "/")
		// http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, _ := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	contents, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var data model.LoginOauth
	json.Unmarshal(contents, &data)

	log.Println(data)

	lo.String(200,"Hello %s!",data.Name)

	// lo.Redirect(http.StatusSeeOther, "/")

}
