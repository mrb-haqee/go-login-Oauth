package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

type API struct {
	r *gin.Engine
}

func NewApi() API {
	api := API{r: NewRouter()}

	return api
}

func (api *API) Start() {
	log.Println("Server running at: http://localhost:8080")
	log.Fatal(api.r.Run(":8080"))
}
