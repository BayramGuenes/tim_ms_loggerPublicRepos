package app_model

import (
	"os"
	"time"

	logdb "tim_presse/tim_ms_loggerPublicRepos/app_database"
	"tim_presse/tim_ms_loggerPublicRepos/utils"

	timLOG "github.com/BayramGuenes/tim_utils_log"
	uid "github.com/lithammer/shortuuid"
)

func generateTransactionKeyShort() string {
	//nanosec := t.UnixNano()
	//millisec := t.UnixNano() / int64(time.Millisecond)

	lUid := uid.New()
	lKey := lUid

	////println("lKey:"+lKey)
	return lKey

}
func createDir(iDir string) {
	if _, err := os.Stat(iDir); os.IsNotExist(err) {

		os.Mkdir(iDir, 0777)
	}
}

func createErrFilePath(iLocationErrLog string, iLogErrTransactPath timLOG.TimLogTransactPath) string {
	t := time.Now()

	thisdate := t.Format("20060102")

	filepath := ""
	filepath = iLocationErrLog +
		thisdate + string(os.PathSeparator) +
		iLogErrTransactPath.TransActApp + string(os.PathSeparator) +
		iLogErrTransactPath.TransActName + string(os.PathSeparator)

	subdir := iLocationErrLog + thisdate
	createDir(subdir)
	subdir += string(os.PathSeparator) + iLogErrTransactPath.TransActApp
	createDir(subdir)
	subdir += string(os.PathSeparator) + iLogErrTransactPath.TransActName
	createDir(subdir)

	return filepath

}

func StartLogTransaction(iInput timLOG.InputParamStartTransact) (eOutput timLOG.OutputParamStartTransact) {
	eOutput = timLOG.OutputParamStartTransact{}
	eOutput.LogTrans.TransAppName = iInput.TransAppName
	eOutput.LogTrans.TransName = iInput.TransName
	eOutput.LogTrans.UName = iInput.UName
	eOutput.LogTrans.ClientAppName = iInput.ClientAppName
	transKeyshort := generateTransactionKeyShort() //(iInput.TransName)
	eOutput.LogTrans.TransKey = transKeyshort

	lLogRecord := logdb.LogRecord{}
	lLogRecord.TransactKey = transKeyshort
	lLogRecord.AppTransact = eOutput.LogTrans.TransAppName
	lLogRecord.AppClient = eOutput.LogTrans.ClientAppName
	lLogRecord.AppLogging = eOutput.LogTrans.TransAppName
	lLogRecord.SvnAppTransact = iInput.TransName
	lLogRecord.SvnAppLogging = iInput.TransName
	lLogRecord.Step = "[START TRANSACTION ] "
	lLogRecord.StepResult = ""
	t := time.Now()
	lLogRecord.StepDateTime = t.Format("20060102150405")
	lLogRecord.UName = eOutput.LogTrans.UName
	lLogRecord.StepContext = ""

	logdb.CreateLogRecord(lLogRecord)

	return
}

func StartLogService(iInput timLOG.InputParamStartTransact) (eOutput timLOG.OutputParamStartTransact) {
	eOutput = timLOG.OutputParamStartTransact{}
	eOutput.LogTrans.TransAppName = iInput.TransAppName
	eOutput.LogTrans.TransName = iInput.TransName
	eOutput.LogTrans.UName = iInput.UName
	eOutput.LogTrans.ClientAppName = iInput.ClientAppName
	transKeyshort := iInput.TransKey
	eOutput.LogTrans.TransKey = transKeyshort

	lLogRecord := logdb.LogRecord{}
	lLogRecord.TransactKey = transKeyshort
	lLogRecord.AppTransact = eOutput.LogTrans.TransAppName
	lLogRecord.AppClient = eOutput.LogTrans.ClientAppName
	lLogRecord.AppLogging = iInput.LoggingAppName
	lLogRecord.SvnAppTransact = iInput.TransName
	lLogRecord.SvnAppLogging = iInput.ServiceName
	lLogRecord.Step = "<Start Remote Service >"
	lLogRecord.StepResult = ""
	t := time.Now()
	lLogRecord.StepDateTime = t.Format("20060102150405")
	lLogRecord.UName = eOutput.LogTrans.UName
	lLogRecord.StepContext = ""

	logdb.CreateLogRecord(lLogRecord)
	return
}

func LogTransStep(iInput timLOG.InputParamLogStep) (eException timLOG.ExceptionStruct) {
	utils.StdOut("app_model/functions.go: LogTransStep called")
	eException = timLOG.ExceptionStruct{}
	lLogRecord := logdb.LogRecord{}
	lLogRecord.TransactKey = iInput.LogTransHeader.TransKey
	lLogRecord.AppTransact = iInput.LogTransHeader.TransAppName
	lLogRecord.AppClient = iInput.LogTransHeader.ClientAppName
	lLogRecord.AppLogging = iInput.AppLogging
	lLogRecord.SvnAppTransact = iInput.LogTransHeader.TransName
	lLogRecord.SvnAppLogging = iInput.AppSVName
	lLogRecord.Step = iInput.StepName
	lLogRecord.StepResult = ""
	t := time.Now()
	lLogRecord.StepDateTime = t.Format("20060102150405")
	lLogRecord.UName = iInput.LogTransHeader.UName
	lLogRecord.StepContext = iInput.Context

	logdb.CreateLogRecord(lLogRecord)

	return
}

func LogTransStepResult(iInput timLOG.InputParamLogStepResult) (eException timLOG.ExceptionStruct) {
	utils.StdOut("app_model/functions.go: LogTransStepResult called")

	eException = timLOG.ExceptionStruct{}
	lLogRecord := logdb.LogRecord{}
	lLogRecord.TransactKey = iInput.LogTransHeader.TransKey
	lLogRecord.AppTransact = iInput.LogTransHeader.TransAppName
	lLogRecord.AppClient = iInput.LogTransHeader.ClientAppName
	lLogRecord.AppLogging = iInput.AppLogging
	lLogRecord.SvnAppTransact = iInput.LogTransHeader.TransName
	lLogRecord.SvnAppLogging = iInput.AppSVName
	lLogRecord.Step = iInput.StepName
	lLogRecord.StepResult = iInput.StepResult
	t := time.Now()
	lLogRecord.StepDateTime = t.Format("20060102150405")
	lLogRecord.UName = iInput.LogTransHeader.UName
	lLogRecord.StepContext = iInput.Context

	logdb.CreateLogRecord(lLogRecord)
	return
}

func FinishLogTransaction(iInput timLOG.InputParamFinishTransact) (eException timLOG.ExceptionStruct) {
	utils.StdOut("app_model/functions.go: FinishLogTransaction called")

	eException = timLOG.ExceptionStruct{}
	lLogRecord := logdb.LogRecord{}
	lLogRecord.TransactKey = iInput.LogTransHeader.TransKey
	lLogRecord.AppTransact = iInput.LogTransHeader.TransAppName
	lLogRecord.AppClient = iInput.LogTransHeader.ClientAppName
	lLogRecord.AppLogging = iInput.LogTransHeader.TransAppName
	lLogRecord.SvnAppTransact = iInput.LogTransHeader.TransName
	lLogRecord.SvnAppLogging = iInput.LogTransHeader.TransName
	lLogRecord.Step = "[END TRANSACTION ]"
	lLogRecord.StepResult = ""
	lLogRecord.TransactStatus = iInput.Status
	t := time.Now()
	lLogRecord.StepDateTime = t.Format("20060102150405")
	lLogRecord.UName = iInput.LogTransHeader.UName
	lLogRecord.StepContext = ""

	logdb.CreateLogRecord(lLogRecord)
	return
}

func FinishLogService(iInput timLOG.InputParamFinishService, iAppLogging string, iSvnAppLogging string) (eException timLOG.ExceptionStruct) {
	utils.StdOut("app_model/functions.go: FinishLogService called")

	eException = timLOG.ExceptionStruct{}
	lLogRecord := logdb.LogRecord{}
	lLogRecord.TransactKey = iInput.LogTransHeader.TransKey
	lLogRecord.AppTransact = iInput.LogTransHeader.TransAppName
	lLogRecord.AppClient = iInput.LogTransHeader.ClientAppName
	lLogRecord.AppLogging = iAppLogging
	lLogRecord.SvnAppTransact = iInput.LogTransHeader.TransName
	lLogRecord.SvnAppLogging = iSvnAppLogging
	lLogRecord.Step = "<End Remote Service >"
	if iInput.Status == timLOG.CoTransStatusFinishedOk {
		lLogRecord.StepResult = iInput.Status
	} else {
		//lLogRecord.StepResult = timLOG.CoResultTypeErr
		lLogRecord.TransactStatus = iInput.Status
	}
	t := time.Now()
	lLogRecord.StepDateTime = t.Format("20060102150405")
	lLogRecord.UName = iInput.LogTransHeader.UName
	lLogRecord.StepContext = ""

	logdb.CreateLogRecord(lLogRecord)
	return
}

func LogFailedIntoFilesys(iInput timLOG.InputParamFailedToFilesys, iLocationErrLog string) (eException timLOG.ExceptionStruct) {
	utils.StdOut("app_model/functions.go: LogFailedIntoFilesys called")

	t := time.Now()
	lUTime := t.Format("150405")

	timErrTransactionPath := timLOG.TimLogTransactPath{}
	timErrTransactionPath.TransActApp = iInput.LogTransHeader.TransAppName
	timErrTransactionPath.TransActName = iInput.LogTransHeader.TransName
	file := createErrFilePath(iLocationErrLog, timErrTransactionPath) + lUTime + "_" + iInput.LogTransHeader.TransKey
	//file := iLocationLogs + "hallo.txt"
	//println("1 Filename:" + file)

	content := "AppTransactionName=" + iInput.LogTransHeader.TransAppName + " | Transaction=" + iInput.LogTransHeader.TransName + " | TransactionKey=" +
		iInput.LogTransHeader.TransKey + "\n\n"
	for i := 0; i < len(iInput.Items); i++ {
		lItem := iInput.Items[i]
		if lItem.ItemType == "step" {
			content = content + "Step=" + lItem.StepName + "; Context:" + lItem.StepContext + "\n"
		}
		if lItem.ItemType == "result" {
			content = content + "ExecResult=" + lItem.StepResult + "; Context:" + lItem.StepContext + "\n"
		}
	}
	content += "\nNähere Informationen zu diesem Fehler finden Sie in der timlog-Datenbank Tabelle=timlog unter dem transactkey=" + iInput.LogTransHeader.TransKey
	appendToFile(file, content)
	eException = timLOG.ExceptionStruct{}
	return
}
