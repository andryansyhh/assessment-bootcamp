package main

import (
	"assesment/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.PassRouter(r)
	routes.UserRouter(r)

	r.Run(":1244")
}
