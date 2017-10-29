package main

import (
	"github.com/leetpy/rum/conf"
	"github.com/leetpy/rum/router"
)

func main() {
	// init config
	configFilePath := ""

	conf.InitConfig(configFilePath)

	// init router
	router.InitRouter()
}
