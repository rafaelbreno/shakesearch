package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitAPI() {
	api := API{fmt.Sprintf("%s", os.Getenv("API_PREFIX"))}
	api.Listen()
}

func (a API) Listen() {
	apiGroup := r.Group(a.Prefix)
	{
		apiGroup.POST("/", a.test)
	}
}

func (_ API) test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"aaa": "bbb"})
	return
}
