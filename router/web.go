package router

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func InitWEB() {
	web := WEB{fmt.Sprintf("%s", os.Getenv("WEB_PREFIX"))}
	web.Listen()
}

func (w WEB) Listen() {
	webGroup := r.Group(w.Prefix)
	{
		webGroup.GET("/", w.test)
	}

	r.Static("/static", "./static")
}

func (_ WEB) test(c *gin.Context) {
	body, err := ioutil.ReadFile("static/index.html")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.Data(http.StatusOK, "text/html; charset=utf-8", body)
}
