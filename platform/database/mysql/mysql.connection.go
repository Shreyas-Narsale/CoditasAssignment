package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func InitializeMySQLDB() {

}

func GetMySQLDBInstance() (*sqlx.DB, error) {
	return nil, nil
}
