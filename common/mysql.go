package common

import (
	"KlotskiWeb/config"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var MysqlDB *sql.DB

func InitMysqlDB() error {
	conf := config.NewMysqlConfig()

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.DB)

	var err error
	MysqlDB, err = sql.Open("mysql", connStr)
	if err != nil {
		return err
	}

	err = MysqlDB.Ping()
	if err != nil {
		return err
	}
	return nil
}

func CloseMysqlDB() error {
	if MysqlDB != nil {
		return MysqlDB.Close()
	}
	return nil
}
