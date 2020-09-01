package run_server

import (
	"encoding/json"
	"net/http"

	//timLOG "mdh.koeln.ivz.cn.ard.de/bitbucket/scm/mdhpres/tim_utils_log_api.git" 
	timLOG       "github.com/BayramGuenes/tim_utils_log"

	model "tim_presse/tim_ms_loggerPublicRepos/app_model"
	utils "tim_presse/tim_ms_loggerPublicRepos/utils"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func SetRoutes() {
	Router.GET("/", HandleHealthCheck)
	Router.GET("/Hello", HandleHealthCheck)
	Router.POST("/StartTransaction", HandleStartTransaction)
	Router.POST("/StartService", HandleStartService)
	Router.POST("/LogServiceStep", HandleLogServiceStep)
	Router.POST("/LogServiceStepResult", HandleLogServiceStepRes)
	Router.POST("/FinishLogTransaction", HandleFinishTransaction)
	Router.POST("/FinishLogService", HandleFinishService)
	Router.POST("/LogFailedIntoFilesys", HandleLogFailedIntoFilesys)
}

func HandleHealthCheck(c *gin.Context) {

	utils.StdOut("run_server/routes.go: HandleHealthCheck called")
	c.JSON(http.StatusOK, "tim_serv_logger (micro-)service is alive.")
}

func HandleStartTransaction(c *gin.Context) {
	utils.StdOut("run_server/routes.go: HandleStartTransaction called")

	rawdata, _ := c.GetRawData()
	lInput := timLOG.InputParamStartTransact{}
	json.Unmarshal(rawdata, &lInput)
	utils.StdOut("HandleStartTransaction:" + lInput.TransAppName + ";" + lInput.TransName)

	lOutput := model.StartLogTransaction(lInput)

	c.JSON(http.StatusOK, lOutput)
}
func HandleStartService(c *gin.Context) {
	utils.StdOut("run_server/routes.go: HandleStartService called")

	rawdata, _ := c.GetRawData()
	lInput := timLOG.InputParamStartTransact{}
	json.Unmarshal(rawdata, &lInput)
	//log.utils.StdOut("HandleStartService:" + lInput.LoggingAppName)

	lOutput := model.StartLogService(lInput)

	c.JSON(http.StatusOK, lOutput)
}

func HandleLogServiceStep(c *gin.Context) {
	utils.StdOut("run_server/routes.go: HandleLogServiceStep called")

	rawdata, _ := c.GetRawData()
	lInput := timLOG.InputParamLogStep{}
	json.Unmarshal(rawdata, &lInput)
	//log.utils.StdOut("LogServiceStep:" + lInput.StepName + ";" + lInput.Context)
	lOutput := model.LogTransStep(lInput)

	c.JSON(http.StatusOK, lOutput)

}

func HandleLogServiceStepRes(c *gin.Context) {
	utils.StdOut("run_server/routes.go: HandleLogServiceStepRes( called")

	rawdata, _ := c.GetRawData()
	lInput := timLOG.InputParamLogStepResult{}
	json.Unmarshal(rawdata, &lInput)
	utils.StdOut("LogServiceStepres:" + lInput.StepName + ";" + lInput.StepResult)
	lOutput := model.LogTransStepResult(lInput)

	c.JSON(http.StatusOK, lOutput)

}

func HandleFinishTransaction(c *gin.Context) {
	utils.StdOut("run_server/routes.go: HandleFinishTransaction called ")
	rawdata, _ := c.GetRawData()
	lInput := timLOG.InputParamFinishTransact{}
	json.Unmarshal(rawdata, &lInput)
	utils.StdOut("HandleFinishTransaction Input.Status:" + lInput.Status)
	lOutput := model.FinishLogTransaction(lInput)
	c.JSON(http.StatusOK, lOutput)

}

func HandleFinishService(c *gin.Context) {
	utils.StdOut("run_server/routes.go: HandleFinishService called ")

	rawdata, _ := c.GetRawData()
	lInput := timLOG.InputParamFinishService{}
	json.Unmarshal(rawdata, &lInput)
	utils.StdOut("HandleFinishService:" + lInput.Status)
	lOutput := model.FinishLogService(lInput, lInput.AppLogging, lInput.ServiceName)

	c.JSON(http.StatusOK, lOutput)

}

func HandleLogFailedIntoFilesys(c *gin.Context) {
	utils.StdOut("run_server/routes.go: HandleLogFailedIntoFilesys called ")

	rawdata, _ := c.GetRawData()

	lInput := timLOG.InputParamFailedToFilesys{}
	json.Unmarshal(rawdata, &lInput)
	lOutput := timLOG.ExceptionStruct{}
	if TimServLoggerConf.WriteErrToFileSys {
		lOutput = model.LogFailedIntoFilesys(lInput, TimServLoggerConf.LocationErrLog)
	} else {
		utils.StdOut("WriteErrToFileSys is disabled by Configuration(config.json)")
	}

	c.JSON(http.StatusOK, lOutput)

}

/*func HandleCheckDoTraceTransaction(c *gin.Context) {
	log.utils.StdOut("HandleCheckDoTraceTransaction called")

	rawdata, _ := c.GetRawData()
	lInput := timLOG.InputParamStartTransact{}
	json.Unmarshal(rawdata, &lInput)
	//log.utils.StdOut("HandleStartTransaction:" + lInput.App + ";" + lInput.TransName)

	lOutput := model.CheckDoTraceTransaction(lInput)

	c.JSON(http.StatusOK, lOutput)
}
func HandleCheckDisableLogMedia(c *gin.Context) {
	log.utils.StdOut("HandleCheckDisableLogMedia called")

	rawdata, _ := c.GetRawData()
	lInput := timLOG.InputParamCheckDisableLogMedia{}
	json.Unmarshal(rawdata, &lInput)
	//log.utils.StdOut("HandleStartTransaction:" + lInput.App + ";" + lInput.TransName)

	lOutput := model.CheckDisableLogMedia(lInput)

	c.JSON(http.StatusOK, lOutput)
}
*/
