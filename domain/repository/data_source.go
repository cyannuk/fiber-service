package repository

import (
	"database/sql"
	"runtime"

	"fiber-service/config"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/dialects/postgresql"
)

type DataSource struct {
	*reform.DB
}

func (dataSource DataSource) Close() error {
	return dataSource.DBInterface().(*sql.DB).Close()
}

func NewDataSource(config config.DatabaseConfig) (DataSource, error) {
	connConfig, err := pgx.ParseConfig(config.ConnectionString())
	if err != nil {
		return DataSource{}, err
	}
	db := stdlib.OpenDB(*connConfig)
	db.SetMaxOpenConns(runtime.NumCPU() * 4)
	if err = db.Ping(); err != nil {
		return DataSource{}, err
	}
	return DataSource{reform.NewDB(db, postgresql.Dialect, nil)}, nil
}
