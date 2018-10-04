package app

import (
	"github.com/gin-gonic/gin"
	"time"
	"net/http"
)

func Start() {
	router := gin.Default()
	mapUrls(router)
	loggerConfig()
	port := ":8080"
	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	s.ListenAndServe()
}
