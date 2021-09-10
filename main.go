package main

import (
	"Hulk/router"
)

func main() {
	r := router.InitRouter()
	r.Run(":8080")
}
