package utils

import (
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func MysqlInit() {
	conf := mysql.Config{
		Addr:                 "10.20.69.191",
		DBName:               "hackathon",
		User:                 "dev",
		Passwd:               "1234@abcd",
		Timeout:              2 * time.Second,
		AllowNativePasswords: true,
		Net:                  "tcp",
	}

	url := conf.FormatDSN()
	var err error
	Engine, err = xorm.NewEngine("mysql", url)
	if err != nil {
		panic(err)
	}

	err = Engine.Ping()
	if err != nil {
		panic(err)
	}

}
