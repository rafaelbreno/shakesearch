package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a API) Listen() {
	apiGroup := r.Group(a.Prefix)
	{
		apiGroup.GET("/", a.test)
	}
}

func (_ API) test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"aaa": "bbb"})
	return
}
