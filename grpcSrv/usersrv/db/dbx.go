package db

import (
	"fmt"

	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"

	"github.com/tietang/dbx"
	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

var database *dbx.Database

func init() {
	file := kvs.GetCurrentFilePath("../config/config.ini", 1)
	fmt.Println("filepath:", file)
	conf := ini.NewIniFileConfigSource(file)
	dbx := new(DbxDatabaseStarter)
	dbx.Setup(conf)
}

//DbxDatabase get database instance
func DbxDatabase() *dbx.Database {
	return database
}

//DbxDatabaseStarter starter
type DbxDatabaseStarter struct {
}

//Setup dbx
func (dbStarter *DbxDatabaseStarter) Setup(conf kvs.ConfigSource) {

	setting := dbx.Settings{}
	err := kvs.Unmarshal(conf, &setting, "mysql")
	if err != nil {
		panic(err)
	}
	fmt.Println("mysql:", setting.ShortDataSourceName())
	db, err := dbx.Open(setting)
	if err != nil {
		panic(err)
	}
	fmt.Println("mysqlxxxxxxxxxxxxx:", db)
	database = db
}
