package router

import (
	"github.com/gin-gonic/gin"
)

/* With two structs it's possible
 * to define differents structures
 * for each one, without interfering
 * the others.
**/

type API struct {
	Prefix string
}

type WEB struct {
	Prefix string
}

type Router interface {
	Listen()
}

var r *gin.Engine

func Listen() {
	// Defining r global variable
	r = gin.Default()

	// Web Routes
	web := WEB{"/"}
	web.Listen()

	// API Routes
	api := API{"/api"}
	api.Listen()

	r.Run(":8080")
}
