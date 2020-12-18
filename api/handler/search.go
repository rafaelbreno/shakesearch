package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"pulley.com/shakesearch/cmd/search"
)

type Query struct {
	Query string `json:"query"`
}

var w *search.Works
var load_error error

func init() {
	w, load_error = search.Load("completeworks.txt")
}

func Search(c *gin.Context) {
	var q Query

	if load_error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": load_error.Error()})
		return
	}

	if err := c.BindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	r, err := w.Search(q.Query)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"resp":     r,
		"contents": w.Contents,
	})
}
