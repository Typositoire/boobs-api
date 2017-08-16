package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.html")
	router.Static("/css", "./css")
	router.Static("/img", "./img")

	router.GET("/", index)
	router.GET("/health", health)

	v1Boobs := router.Group("/v1/boobs")
	{
		v1Boobs.GET("/:amount", getBoobs)
	}

	router.NoRoute(notFound)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
}
