package dal

import (
    "database/sql"
)

var pool *sql.DB

func InitSql() {

}

func UninitSql() {
    if pool {
        pool.Close()
    }
}
