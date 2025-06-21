package config

import (
	"os"
	"strconv"
)

type DbConfig struct {
	MySQL MySQLConfig
}

type MySQLConfig struct {
	IP       string
	Port     int
	Username string
	Password string
	DbName   string
}

var dbConf = &DbConfig{}

func GetDBConfig() *DbConfig {
	return dbConf
}

func LoadDbConfig() {
	dbConf.MySQL.IP = os.Getenv("MYSQL_IP")
	dbConf.MySQL.Port = 3306
	dbConf.MySQL.Port, _ = strconv.Atoi(os.Getenv("MYSQL_PORT"))
	dbConf.MySQL.Username = os.Getenv("MYSQL_USER")
	dbConf.MySQL.Password = os.Getenv("MYSQL_PASSWORD")
	dbConf.MySQL.DbName = os.Getenv("MYSQL_DATABASE")
}
