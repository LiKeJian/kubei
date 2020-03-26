package main

import (
	log "github.com/sirupsen/logrus"
	"kubei/pkg/config"
	"kubei/pkg/webapp"
)

func initLog(verbose bool) {
	if verbose {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
}

func main() {
	conf := config.LoadConfig()
	scanConfig := config.LoadScanConfig()

	initLog(conf.Verbose)
	kubeiWebapp := webapp.Init(conf, scanConfig)
	kubeiWebapp.Run()
}
