package main

// Importing fmt
import (
	"github.com/gin-gonic/gin"
	"github.com/jonathanmorais/endoenginnering/t1d-api/controllers/calculate"
	"github.com/jonathanmorais/endoenginnering/t1d-api/controllers/healthcheck"
)

// Calling main
func main() {

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.GET("/health", healthcheck.Ok)
	router.POST("/calculate", calculate.Calculate)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	//router.Run(":3000")
}
