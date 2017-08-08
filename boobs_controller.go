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
	_int, err := strconv.Atoi(c.Param("amount"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"boobs": genBoobs(_int),
	})
}

func genBoobs(amount int) []string {
	rand.Seed(time.Now().Unix())

	var boob string
	var boobs []string

	for i := 0; i < amount; i++ {
		boob = boobList[rand.Intn(len(boobList))]
		boobs = append(boobs, boob)
	}

	return boobs
}
