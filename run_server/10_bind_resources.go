package run_server

import appdb "tim_presse/tim_ms_loggerPublicRepos/app_database"

func BindResources() {
	appdb.InitAppDatabase(TimServLoggerConf.DB_SQLApplschemaName, TimServLoggerConf.DB_UserPwd, TimExtParams.DBServer, TimExtParams.DBPort)
	return
}
