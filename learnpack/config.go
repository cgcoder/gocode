package main

import (
    "os"
    "errors"
    "fmt"
)

type Config interface {
    GetDbConnString() string
}

type config struct {
    dbConnString string
}

var AppConfig = config{}

func InitConfig() error {
    AppConfig.dbConnString = os.Getenv(string(DB_CONN))
    if len(AppConfig.dbConnString) <= 0 {
        return errors.New(fmt.Sprintf("missing ENV var %s", string(DB_CONN)))
    }

    return nil
}

func (c config) GetDbConnString() string {
    return c.dbConnString
}
