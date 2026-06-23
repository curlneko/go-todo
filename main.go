package main

import (
	"gin-todo/routes"
)

func main() {
	r := routes.SetupRouter()

	r.Run(":8080")
}
