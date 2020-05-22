package main

import (
		"gojenkis/router"
)

func main() {

	router := router.InitRouter()
	router.Run(":8090")
}