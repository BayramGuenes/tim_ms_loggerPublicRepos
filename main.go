package main

import (
	"log"
	run_server "tim_presse/tim_ms_loggerPublicRepos/run_server"
	utils "tim_presse/tim_ms_loggerPublicRepos/utils"
)

var (
	timServLoggerConf run_server.TimServLoggerConfStruct
	timMicroservices  run_server.TimMicroservicesStruct
	timExtParams      run_server.TimServLoggerExtParams
	exception         run_server.ExceptionStruct
)

func main() {

	utils.StdOut("Version 0.0.1-rc01")

	timExtParams, timServLoggerConf, timMicroservices, exception = run_server.ReadSettings()

	if exception.Occured {
		log.Fatal(exception.ErrTxt)
		return
	}

	run_server.BindResources()
	if exception.Occured {
		log.Fatal(exception.ErrTxt)
		return
	}
	run_server.UseGinAsRouter()
	run_server.Router.Use(run_server.CORSMiddleware())
	run_server.SetRoutes()
	run_server.ListenAndServe()

}
