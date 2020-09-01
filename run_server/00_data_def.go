package run_server

import "github.com/gin-gonic/gin"

type TimServLoggerConfStruct struct {
	DB_UserPwd           string
	DB_SQLApplschemaName string
	LocationErrLog       string
	WriteErrToFileSys    bool
}

type TimServLoggerExtParams struct {
	ConfigFile       string
	TimMicroServices TimMicroservicesStruct
	DBServer         string
	DBPort           string
}

type TimMicroservicesStruct struct {
	NameTimLogServer string
	PortTimLogServer string
}

type ExceptionStruct struct {
	Occured bool
	ErrTxt  string
}

var (
	Router               *gin.Engine
	TimExtParams         TimServLoggerExtParams
	TimServLoggerConf    TimServLoggerConfStruct
	TimMicroservices     TimMicroservicesStruct
	exceptionReadSetting ExceptionStruct
)
