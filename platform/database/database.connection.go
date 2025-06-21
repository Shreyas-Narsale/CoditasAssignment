package database

import (
	"CODITAS_TASK/platform/database/mysql"
)

func DatabaseConnections() {
	mysql.GetMySQLDBInstance()
}
