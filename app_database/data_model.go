package app_database


type LogRecord struct {
	LogID          int64
	TransactKey    string
	AppTransact    string
	AppClient      string
	AppLogging     string
	SvnAppTransact string
	SvnAppLogging  string
	Step           string
	StepResult     string
	StepDateTime   string
	TransactStatus string
	UName          string
	StepContext    string
}

type ExceptionStruct struct {
	Occured bool
	ErrTxt  string
}

var (
	MysqlApplschemaName string
	MysqldbserverRef    string
	Databaselocation    string
)

