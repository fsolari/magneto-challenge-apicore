package handler

import (
	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine) {

	router.GET("/ping", Ping)
	router.POST("/mutant", GetDNATest)
	router.GET("/stats", GetDNAStats)
	router.GET("/", Home)
}