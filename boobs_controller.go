package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/DataDog/datadog-go/statsd"
	"github.com/gin-gonic/gin"
)

var boobList = []string{
	"( o y o )", "( . y . )", "( O Y O )",
	"( O Y o )", "\\./\\./", "(*Y*)", "( . Y . )",
	"(.Y.)", "(。 ㅅ  。)", "(@ㅅ@)", "(•_ㅅ_•)",
	"(o)(o)", "(•)(•)", "(.)(.)(.)", "[○][°]",
	"[°][○]", "( o Y O )", "( + )( + )", "oo",
	"{ O }{ O }", "( ^ )( ^ )", "(Q)(O)", "(O)(Q)",
	"(p)(p)", "\\o/\\o/", "(  -  )(  -  )"}

func getBoobs(c *gin.Context) {
	d, err := statsd.New("127.0.0.1:8125")

	if err != nil {
		fmt.Println(err)
	}

	defer d.Close()

	var key string

	d.Namespace = "boobs-api."
	d.Tags = append(d.Tags, "ENV:"+os.Getenv("ENVIRONMENT"))

	sfw := false
	limit := 5000000

	if len(c.Request.URL.Query().Get("sfw")) > 0 {
		sfw = true
		d.Tags = append(d.Tags, "rating:sfw")
		key = "boobs_counter_sfw"
	} else {
		d.Tags = append(d.Tags, "rating:nsfw")
		key = "boobs_counter_nsfw"
	}

	if len(os.Getenv("BOOBS_LIMIT")) > 0 {
		var err error
		limit, err = strconv.Atoi(os.Getenv("BOOBS_LIMIT"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}
	}

	_int, err := strconv.Atoi(c.Param("amount"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if _int > limit {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"message": "Too Many Boobs, the limit is currently " + strconv.Itoa(limit) + ".",
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"boobs": genBoobs(_int, sfw),
	})

	err = d.Count("amount", int64(_int), nil, 1)

	if err != nil {
		fmt.Println("Unable to send to DD")
	}

	redisClient.IncrBy(key, int64(_int))
}

func genBoobs(amount int, sfw bool) []string {
	rand.Seed(time.Now().Unix())

	var boob string
	var boobs []string

	for i := 0; i < amount; i++ {
		if sfw {
			boob = "(omit)(omit)"
		} else {
			boob = boobList[rand.Intn(len(boobList))]
		}
		boobs = append(boobs, boob)
	}

	return boobs
}
