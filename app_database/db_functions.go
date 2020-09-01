package app_database

import (
	"database/sql"
	"time"
)

func InitAppDatabase(iDBSchemaName string, iDBUsrPwd string, iDBServer, iDBPort string) {
	MysqlApplschemaName = iDBSchemaName

	MysqldbserverRef = iDBUsrPwd + "@tcp(" +
		iDBServer + ":" + iDBPort + ")"
	Databaselocation = MysqldbserverRef + "/" + MysqlApplschemaName

}

func CreateLogRecord(iRecord LogRecord) (eException ExceptionStruct) {
	eException = ExceptionStruct{}
	println("CreateLogRecord called")
	db, err := sql.Open("mysql", Databaselocation)
	if err != nil {
		eException.Occured = true
		eException.ErrTxt = err.Error()
		return
	}
	defer db.Close()
	insstmt := "INSERT timlog SET" +
		" transactkey=?, apptransact=?, appclient=?, applogging=?, svnapptransact=?, svnapplogging=?," +
		" step=?, stepresult=?, stepdatetime=?," +
		" transactstatus=? ,  uname=?, stepcontext=?"
	stmt, err := db.Prepare(insstmt)
	if err != nil {
		println("CreateLogRecord called err:" + err.Error())

		eException.Occured = true
		eException.ErrTxt = err.Error()
		return
	}
	currentTime := time.Now()
	timePostfix := currentTime.Format("20060102150405")

	_, err = stmt.Exec(
		iRecord.TransactKey,
		iRecord.AppTransact,
		iRecord.AppClient,
		iRecord.AppLogging,
		iRecord.SvnAppTransact,
		iRecord.SvnAppLogging,
		iRecord.Step,
		iRecord.StepResult,
		timePostfix,
		iRecord.TransactStatus,
		iRecord.UName,
		iRecord.StepContext,
	)
	if err != nil {
		eException.Occured = true
		println("CreateLogRecord called err:" + err.Error())

		eException.ErrTxt = err.Error()
		return
	}

	return
}
