package main

import (
	"AnalyzingApp/app/controllers"
	"AnalyzingApp/config"
	"AnalyzingApp/config/utils"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)
	controllers.StreamIngestionData()
	controllers.StartWebServer()
}
