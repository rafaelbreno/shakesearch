package app

import (
	"github.com/joho/godotenv"
	"pulley.com/shakesearch/router"
)

func Init() {
	setOS()

	router.Listen()
}

func setOS() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
