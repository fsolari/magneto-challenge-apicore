package app

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"io"
)

func loggerConfig() {
	logfile, err := os.Create(LogFilePath)
	if err != nil {

		log.Fatalln("Failed to create request log file:", err)
	}
	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(gin.DefaultWriter)
}
