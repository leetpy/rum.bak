package main

import (
	"github.com/leetpy/rum/conf"
	"github.com/leetpy/rum/router"
	"github.com/gin-gonic/gin"
	"os"
	"io"
)

func main() {
	// init config
	configFilePath := ""

	conf.InitConfig(configFilePath)

	// init log
	gin.DisableConsoleColor()
	logfile := conf.Conf.Logfile
	f, _ := os.Create(logfile)
	gin.DefaultWriter = io.MultiWriter(f)

	// init router
	router.InitRouter()
}
