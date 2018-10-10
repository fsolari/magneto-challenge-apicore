package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/magneto-challenge-apicore/src/handler"
)

func mapUrls(router *gin.Engine) {

	router.GET("/ping", handler.Ping)
	router.POST("/mutant", handler.GetDNATest)
	router.GET("/stats", handler.GetDNAStats)
	router.GET("/", handler.Home)
}