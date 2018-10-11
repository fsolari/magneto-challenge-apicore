package app

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	"os"
	"fmt"
	"github.com/mercadolibre/magneto-challenge-apicore/handler"
)

func Start() {
	setupRouter()
	loggerConfig()
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	handler.MapUrls(router)

	s := &http.Server{
		Addr:           getPort(),
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	s.ListenAndServe()
	return router
}

func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "5000"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}