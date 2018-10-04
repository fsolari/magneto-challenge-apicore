package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/mercadolibre/test/magneto-challenge-apicore/service"
)

func GetStats(c *gin.Context) {

	stats, err := service.CalculateStats()

	if err != nil {
		log.Printf("Error calculating stats %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusOK, stats)

}

