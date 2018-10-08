package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/mercadolibre/magneto-challenge-apicore/domain"
	"github.com/mercadolibre/magneto-challenge-apicore/service"
)

func IsMutant(c *gin.Context) {
	var dna domain.DNA
	bindErr := c.BindJSON(&dna)

	if bindErr != nil || !domain.IsDNAValid(dna) {
		errMSG := "Failed to process request: DNA format invalid."
		c.JSON(http.StatusBadRequest, errMSG)
		return
	}

	isMutant, err := service.DNATest(dna)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	MSG := "DNA: "
	if !isMutant {
		c.JSON(http.StatusBadRequest, MSG + "Human.")
	} else {
		c.JSON(http.StatusOK, MSG + "Mutant.")
	}

}
