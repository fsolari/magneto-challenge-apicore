package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(c *gin.Context) {
	c.JSON(http.StatusOK, "Welcome to Magneto API. For usage see: https://github.com/fsolari/magneto-challenge-apicore")
}
