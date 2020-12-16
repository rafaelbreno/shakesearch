package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Query struct {
	Query string `json:"query"`
}

//type SearchResult struct {
//}

func Search(c *gin.Context) {
	var q Query

	if err := c.BindJSON(&q); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"resp": q})
}
