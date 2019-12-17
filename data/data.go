package data

import (
	"config"
	"crypto/sha1"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/satori/go.uuid"
	"log"
	"time"
)

// should this be a global variable?
var Db *sql.DB

const (
	timeFormat = "Jan 2, 2006 at 3:04pm"
)

func init() {
	dataSourceName := config.AppConfig.DbConfig.UserName +
		":" + config.AppConfig.DbConfig.Password +
		"@tcp(" + config.AppConfig.DbConfig.IP + ")/" +
		config.AppConfig.DbConfig.DbName +
		"?charset=utf8&parseTime=true"

	var err error
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func formatTimeToString(t time.Time) string {
	return t.Format(timeFormat)
}

func createUUID() (string, error) {
	u4, err := uuid.NewV4()
	if err != nil {
		return "", fmt.Errorf("error while creating UUID:%s", err.Error())
	}
	return u4.String(), nil
}

func Encrypt(str string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(str)))
}
