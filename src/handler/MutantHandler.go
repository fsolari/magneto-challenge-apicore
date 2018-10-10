package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/mercadolibre/magneto-challenge-apicore/src/domain"
	"github.com/mercadolibre/magneto-challenge-apicore/src/service"
)

func GetDNATest(c *gin.Context) {
	var dna domain.DNA
	bindErr := c.BindJSON(&dna)

	if bindErr != nil || !domain.IsDNAValid(dna) {
		err := "[MutantHandler.GetDNATest] Error processing request: DNA format invalid."
		c.JSON(http.StatusBadRequest, err)
		return
	}

	isMutant, err := service.DNATest(dna)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if !isMutant {
		c.JSON(http.StatusForbidden, isMutant)
	} else {
		c.JSON(http.StatusOK, isMutant)
	}

}
