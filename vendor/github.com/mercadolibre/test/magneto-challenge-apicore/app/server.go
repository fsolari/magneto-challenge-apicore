package app

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

func Start() {
	SetupRouter()
	loggerConfig()
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	mapUrls(router)
	port := ":8080"
	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	s.ListenAndServe()
	return router
}