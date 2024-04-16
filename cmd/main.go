package main

import (
	"online-ordering-app/internal/router"
)

func main() {
	r := router.SetupRouter()
	r.Run()
}