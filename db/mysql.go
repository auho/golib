package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"fmt"
	"database/sql"
)

type MysqlDriver struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	Charset  string
}

func (o MysqlDriver) dns() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", o.User, o.Password, o.Host, o.Port, o.Database, o.Charset)
}

func InitMysql(driver MysqlDriver) *MysqlDb {
	mysql := NewDb()
	mysql.Connect(driver)

	return mysql
}

type MysqlDb struct {
	connection *xorm.Engine
}

func NewDb() *MysqlDb {
	return &MysqlDb{}
}

func (db *MysqlDb) Connect(config MysqlDriver) {
	engine, err := xorm.NewEngine("mysql", config.dns())
	if err != nil {
		panic(err)
	}

	db.connection = engine
}

func (db *MysqlDb) Query(sql string) ([]map[string][]byte, error) {
	return db.connection.Query(sql)
}

func (db *MysqlDb) Execute(sql string) (sql.Result, error) {
	return db.connection.Exec(sql)
}

func (db *MysqlDb) Count(sql string) (int64, error) {
	return db.connection.Count(sql)
}

func (db *MysqlDb) Close() {
	db.connection.Close()
}
