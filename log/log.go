package log

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"

	"url-shortener/config"
)

var Logger *log.Logger

func init() {

	err := os.MkdirAll(config.Config.Log.Path, os.ModePerm)
	if err != nil {
		panic(err)
	}

	logFile, err := os.Create(filepath.Join(config.Config.Log.Path, config.Config.Log.Server))
	if err != nil {
		panic(err)
	}
	Logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)

	f, err := os.Create(filepath.Join(config.Config.Log.Path, config.Config.Log.Request))
	if err != nil {
		panic(err)
	}
	gin.DefaultWriter = io.MultiWriter(f)
}
