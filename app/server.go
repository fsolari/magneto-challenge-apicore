package app

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	"os"
	"log"
)

func Start() {
	SetupRouter()
	loggerConfig()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	mapUrls(router)

	port := os.Getenv("PORT") || 5000

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	s.ListenAndServe()
	return router
}