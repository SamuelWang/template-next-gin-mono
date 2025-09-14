package main

import (
	"backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    // Setup routes
    routes.SetupRoutes(r)

    // Start the server
    r.Run(":8080") // listen and serve on localhost:8080
}