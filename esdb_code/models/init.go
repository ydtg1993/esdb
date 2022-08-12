package models

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var DB *sql.DB

func init() {
	//使用mysql连接池
	mysqlhost := os.Getenv("mysqlhost")
	mysqluser := os.Getenv("mysqluser")
	mysqlpassword := os.Getenv("mysqlpass")
	mysqldbname := os.Getenv("mysqldbname")

	mysqllifetime, _ := strconv.Atoi(os.Getenv("mysqllifetime"))
	mysqlidletime, _ := strconv.Atoi(os.Getenv("mysqlidletime"))
	mysqlmaxconn, _ := strconv.Atoi(os.Getenv("mysqlmaxconn"))

	DB = poolInitMysql(mysqlhost, mysqluser, mysqlpassword, mysqldbname)
	DB.SetConnMaxIdleTime(time.Duration(mysqlidletime) * time.Second) //最大空闲时间
	DB.SetConnMaxLifetime(time.Duration(mysqllifetime) * time.Second) //连接最长生命周期时间
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(mysqlmaxconn) //最大连接数
}

//mysql连接池
func poolInitMysql(server, username, password, dbname string) *sql.DB {
	sMysql := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, server, dbname)
	db, err := sql.Open("mysql", sMysql)
	if err != nil {
		panic("connect mysql error->" + err.Error())
		return db
	}

	//设置utf-8
	err = db.Ping()
	if err != nil {
		panic("test mysql error->" + err.Error())
		return db
	}

	return db
}

//uuid
func uuid4() string {
	u4 := uuid.New()
	return u4.String()
}
