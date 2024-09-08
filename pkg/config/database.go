package config

import (
	"os"
	"strconv"
)

type DbConfig struct {
	Postgress PostgressConfig
	Redis     RedisConfig
}

type PostgressConfig struct {
	IP       string
	Port     int
	Username string
	Password string
	DbName   string
}

type RedisConfig struct {
	IP       string
	Port     int
	Password string
}

var dbConf = &DbConfig{}

func GetDBConfig() *DbConfig {
	return dbConf
}

func LoadDbConfig() {
	dbConf.Postgress.IP = os.Getenv("PG_IP")
	dbConf.Postgress.Port = 5432
	dbConf.Postgress.Port, _ = strconv.Atoi(os.Getenv("PG_PORT"))
	dbConf.Postgress.Username = os.Getenv("PG_USER")
	dbConf.Postgress.Password = os.Getenv("PG_PASSWORD")
	dbConf.Postgress.DbName = os.Getenv("PG_DATABASE")
	dbConf.Redis.IP = os.Getenv("REDIS_IP")
	dbConf.Redis.Port = 6379
	dbConf.Redis.Port, _ = strconv.Atoi(os.Getenv("REDIS_PORT"))
	dbConf.Redis.Password = os.Getenv("REDIS_PASSWORD")
}
