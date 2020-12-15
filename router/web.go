package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (w WEB) Listen() {
	apiGroup := r.Group(w.Prefix)
	{
		apiGroup.GET("/", w.test)
	}
}

func (_ WEB) test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"aaa": "bbb"})
}
