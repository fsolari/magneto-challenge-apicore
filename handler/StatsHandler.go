package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"github.com/mercadolibre/magneto-challenge-apicore/service"
)

func GetDNAStats(c *gin.Context) {

	stats, err := service.CalculateDNAStats()

	if err != nil {
		log.Printf("[StatsHandler.GetDNAStats] Error calculating DNA stats %v\n", err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, stats)

}

