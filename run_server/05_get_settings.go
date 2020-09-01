package run_server

import (
	"encoding/json"
	"flag"
	"os"
	"strings"
	"tim_presse/tim_ms_loggerPublicRepos/utils"
)

func ReadSettings() (
	etimExtParams TimServLoggerExtParams,
	etimServLoggerConf TimServLoggerConfStruct,
	eTimMicroservices TimMicroservicesStruct,
	//etimServExtParams TimLoggerConf
	eExceptionOP ExceptionStruct) {

	etimExtParams = TimServLoggerExtParams{}
	exceptionReadSetting = ExceptionStruct{}

	lTimExtParams := loadFromOSFlags()
	lTimExtParams = loadFromOSEnv(lTimExtParams)
	lTimExtParams = loadFromOSArgs(lTimExtParams)

	TimMicroservices = lTimExtParams.TimMicroServices

	TimExtParams = lTimExtParams
	//timMicroservices.NameTimLogServer = "tim-log-server"
	//timMicroservices.PortTimLogServer = "4080"

	if len(lTimExtParams.ConfigFile) == 0 {
		eExceptionOP.Occured = true
		eExceptionOP.ErrTxt = //"Achtung !!!  Service konnte nicht gestartet werden." +
			" Bitte Konfigurationsdatei-Namen  über Umgebungsparameter" +
				" 'configfile' bzw als Startargument(osparam) 'configfile=<path>' angeben."
		return
	}
	TimServLoggerConf, eExceptionOP = getConf(lTimExtParams.ConfigFile)
	utils.StdOut("conf:" + TimServLoggerConf.DB_SQLApplschemaName)
	utils.StdOut("dbusr:" + TimServLoggerConf.DB_UserPwd)

	if len(TimServLoggerConf.LocationErrLog) == 0 {
		eExceptionOP.Occured = true
		eExceptionOP.ErrTxt = //"Achtung !!!  Service konnte nicht gestartet werden." +
			" Bitte Pfad für Err-Loggingdateien in der Konfigurations-Datei angeben."

		return
	}
	etimExtParams = TimExtParams
	etimServLoggerConf = TimServLoggerConf
	eTimMicroservices = TimMicroservices
	eExceptionOP = exceptionReadSetting

	return
}
func loadFromOSFlags() (eTimServLoggerExtParams TimServLoggerExtParams) {

	eTimServLoggerExtParams = TimServLoggerExtParams{}
	pointerConfigFile := flag.String("configfile", "", "./config.conf")
	pointerNameTimLogServer := flag.String("SVCTimLogServer", "127.0.0.1", "Service LocationID TimLogServer")
	pointerPortTimLogServer := flag.String("PortTimLogServer", "4080", "Port TimLogServer")
	pointerNameDBServer := flag.String("SVCDBServer", "127.0.0.1", "Service LocationID DBServer")
	pointerPortDBServer := flag.String("PortDBServer", "3036", "Port DBServer")

	flag.Parse()

	eTimServLoggerExtParams.ConfigFile = *pointerConfigFile
	eTimServLoggerExtParams.TimMicroServices.NameTimLogServer = *pointerNameTimLogServer
	eTimServLoggerExtParams.TimMicroServices.PortTimLogServer = *pointerPortTimLogServer
	eTimServLoggerExtParams.DBServer = *pointerNameDBServer
	eTimServLoggerExtParams.DBPort = *pointerPortDBServer

	return
}

func loadFromOSEnv(iTimServLoggerExtParams TimServLoggerExtParams) (eTimServLoggerExtParams TimServLoggerExtParams) {

	eTimServLoggerExtParams = iTimServLoggerExtParams

	environList := os.Environ()
	leng := len(environList)
	for i := 0; i < leng; i++ {
		if i > 0 {
			osenvparamval := environList[i]
			splittedString := strings.Split(osenvparamval, "=")
			paramname := splittedString[0]
			paramval := splittedString[1]

			if paramname == "configfile" || paramname == "configFile" || paramname == "CONFIGFILE" {
				eTimServLoggerExtParams.ConfigFile = paramval
			}
			if paramname == "svctimlogserver" || paramname == "SVCTimLogServer" || paramname == "SVCTIMLOGSERVER" {
				eTimServLoggerExtParams.TimMicroServices.NameTimLogServer = paramval
			}
			if paramname == "porttimlogserver" || paramname == "PortTimLogServer" || paramname == "PORTTIMLOGSERVER" {
				eTimServLoggerExtParams.TimMicroServices.PortTimLogServer = paramval
			}
			if paramname == "svcdbserver" || paramname == "SVCDBServer" || paramname == "SVCDBSERVER" {
				eTimServLoggerExtParams.DBServer = paramval
			}
			if paramname == "portdbserver" || paramname == "PortDBServer" || paramname == "PORTDBSERVER" {
				eTimServLoggerExtParams.DBPort = paramval
			}

		}
	}
	return
}

func loadFromOSArgs(iTimServLoggerExtParams TimServLoggerExtParams) (eTimServLoggerExtParams TimServLoggerExtParams) {

	eTimServLoggerExtParams = iTimServLoggerExtParams
	leng := len(os.Args)
	for i := 0; i < leng; i++ {
		if i > 0 {
			osparam := os.Args[i]
			splittedString := strings.Split(osparam, "=")
			var namevalues []string
			namevalues = append(namevalues, splittedString...)
			leng := len(namevalues)
			if leng > 1 {
				paramname := splittedString[0]
				paramval := splittedString[1]

				if paramname == "configfile" || paramname == "configFile" || paramname == "CONFIGFILE" {
					eTimServLoggerExtParams.ConfigFile = paramval
				}
				if paramname == "SVCTimLogServer" || paramname == "SVCTIMLOGSERVER" {
					eTimServLoggerExtParams.TimMicroServices.NameTimLogServer = paramval
				}
				if paramname == "PortTimLogServer" || paramname == "PORTTIMLOGSERVER" {
					eTimServLoggerExtParams.TimMicroServices.PortTimLogServer = paramval
				}
				if paramname == "svcdbserver" || paramname == "SVCDBServer" || paramname == "SVCDBSERVER" {
					eTimServLoggerExtParams.DBServer = paramval
				}
				if paramname == "portdbserver" || paramname == "PortDBServer" || paramname == "PORTDBSERVER" {
					eTimServLoggerExtParams.DBPort = paramval
				}

			}
		}
	}

	return
}

func getConf(iConfLocation string) (eTimServLoggerConf TimServLoggerConfStruct, eExceptionOP ExceptionStruct) {
	TimServLoggerConf = TimServLoggerConfStruct{}
	eExceptionOP = ExceptionStruct{}
	//utils.StdOut("getConf called")
	filePathAndName := iConfLocation
	utils.StdOut("filePathAndName:" + filePathAndName)

	file, err := os.Open(filePathAndName)
	if err != nil {
		utils.StdOut("err.Error():" + err.Error())
		eExceptionOP.Occured = true
		eExceptionOP.ErrTxt = err.Error()
		return TimServLoggerConf, eExceptionOP
	}
	// guinvi_server.exe
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&TimServLoggerConf)
	if err != nil {
		utils.StdOut("err.Error():" + err.Error())
		eExceptionOP.Occured = true
		eExceptionOP.ErrTxt = err.Error()
		return TimServLoggerConf, eExceptionOP
	}
	eTimServLoggerConf = TimServLoggerConf
	return
}
