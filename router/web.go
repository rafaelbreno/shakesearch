package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitWEB() {
	web := WEB{fmt.Sprintf("%s", os.Getenv("WEB_PREFIX"))}
	web.Listen()
}

func (w WEB) Listen() {
	apiGroup := r.Group(w.Prefix)
	{
		apiGroup.GET("/", w.test)
	}
}

func (_ WEB) test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"aaa": "bbb"})
}
