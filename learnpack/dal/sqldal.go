package dal

import (
	"context"
	"fmt"
	"os"

	"github.com/gocode/learnpack/config"
	"github.com/jackc/pgx/v4/pgxpool"
)

var pool *pgxpool.Pool

// InitSql setup sql
func InitSQL() error {
	var err error
	pool, err = pgxpool.Connect(context.Background(), config.AppConfig.GetDbConnString())
	if err != nil {
		fmt.Fprintf(os.Stderr, "unable to init db %v", err)
		return err
	}

	return nil
}

// UninitSql tear down sql
func UninitSQL() {
	if pool != nil {
		pool.Close()
	}
}

// GetPool get sql conn
func GetPool() *pgxpool.Pool {
	return pool
}
