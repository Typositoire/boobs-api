package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func index(c *gin.Context) {
	boobsCounterNsfw, err := redisClient.Get("boobs_counter_nsfw").Result()
	if err == redis.Nil {
		boobsCounterNsfw = "0"
	} else if err != nil {
		fmt.Println(err)
	}

	boobsCounterSfw, err := redisClient.Get("boobs_counter_sfw").Result()
	if err == redis.Nil {
		boobsCounterSfw = "0"
	} else if err != nil {
		fmt.Println(err)
	}

	c.HTML(
		http.StatusOK,
		"index.html",
		gin.H{
			"title":              "Boobs like a BaaS",
			"boobs_counter_nsfw": boobsCounterNsfw,
			"boobs_counter_sfw":  boobsCounterSfw,
		},
	)
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
