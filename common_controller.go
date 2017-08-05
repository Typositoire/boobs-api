package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Boobs as a Service ©2017",
	})
}

//
//
//

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

//
//
//

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"message": "Not Found",
	})
}
