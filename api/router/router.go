package router

import (
	"fmt"
	"os"

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
	InitWEB()

	// API Routes
	InitAPI()

	// Running router
	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
