package router

import (
	"fmt"
	"os"

	"pulley.com/shakesearch/handler"
)

func InitAPI() {
	api := API{fmt.Sprintf("%s", os.Getenv("API_PREFIX"))}
	api.Listen()
}

func (a API) Listen() {
	apiGroup := r.Group(a.Prefix)
	{
		apiGroup.POST("/", handler.Search)
	}
}
