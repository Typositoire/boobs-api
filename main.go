package main

import (
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var redisStringSplit = strings.Split(os.Getenv("REDIS_URL"), "@")
var redisAddr = redisStringSplit[1]
var redisPassword = strings.Split(redisStringSplit[0], ":")[2]

var redisClient = redis.NewClient(&redis.Options{
	Addr:     redisAddr,
	Password: redisPassword,
	DB:       0,
})

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
