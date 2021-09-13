package main

import (
	"books-api/Routes"
)

func main() {
	router := Routes.SetupRouter()
	router.Run()
}
