package main

import (
	"math/rand"
	"net/http"
	"strconv"
	"time"

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
	sfw := false

	if c.Request.URL.Query().Get("sfw") == "1" {
		sfw = true
	}

	_int, err := strconv.Atoi(c.Param("amount"))

	if _int > 500000 {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"message": "Too Many Boobs, the limit is currently 500000.",
		})

		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"boobs": genBoobs(_int, sfw),
	})
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
