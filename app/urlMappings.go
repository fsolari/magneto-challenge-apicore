package app

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/test/magneto-challenge-apicore/handler"
)

func mapUrls(router *gin.Engine){
	router.GET("/ping", handler.Ping)
	router.POST("/mutant", handler.IsMutant)
	router.GET("/stats", handler.GetStats)
}
