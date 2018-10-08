package app

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
	"os"
	"log"
	"fmt"
)

func Start() {
	SetupRouter()
	loggerConfig()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	mapUrls(router)

	s := &http.Server{
		Addr:           ":" + getPort(),
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
		port = "500"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}