package postgress

import (
	"SOCIAL_MEDIA_APP/pkg/config"
	"SOCIAL_MEDIA_APP/pkg/logger"
	"context"
	"fmt"
	"sync"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	postgresDB *sqlx.DB
	pgOnce     sync.Once
	pgInitErr  error
)

func InitializePostgresDB() {
	DbConfig := config.GetDBConfig()
	logs := logger.GetLogger()

	ip := DbConfig.Postgress.IP
	username := DbConfig.Postgress.Username
	password := DbConfig.Postgress.Password
	port := DbConfig.Postgress.Port
	dbname := DbConfig.Postgress.DbName

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		username, password, ip, port, dbname)

	postgresDB, pgInitErr = sqlx.Connect("postgres", connStr)
	if pgInitErr != nil {
		pgInitErr = fmt.Errorf("unable to connect to PostgreSQL database: %w", pgInitErr)
		return
	}

	pgInitErr = postgresDB.PingContext(context.Background())
	if pgInitErr != nil {
		pgInitErr = fmt.Errorf("unable to ping PostgreSQL database: %w", pgInitErr)
		return
	}
	logs.Info().Msg("Successfully connected to the PostgreSQL database!")
}

func GetPostgressDBInstance() (*sqlx.DB, error) {
	logs := logger.GetLogger()

	pgOnce.Do(InitializePostgresDB)
	if pgInitErr != nil {
		logs.Error().Err(pgInitErr).Msg("postgress error:")
		return nil, pgInitErr
	}

	return postgresDB, pgInitErr
}
