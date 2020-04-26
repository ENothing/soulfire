package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"soulfire/pkg/logging"
)

type Database struct {
	Self *gorm.DB
	//Docker *gorm.DB
}

var DB *Database


func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
	}
}

func GetSelfDB() *gorm.DB {

	host := viper.GetString("Mysql.Host")
	port := viper.GetString("Mysql.Port")
	database := viper.GetString("Mysql.Database")
	user := viper.GetString("Mysql.Username")
	password := viper.GetString("Mysql.Password")


	return connect(user, password, host+":"+port, database)
}


func connect(user, password, host, database string) *gorm.DB {

	conf := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, database,
	)

	db, err := gorm.Open("mysql", conf)
	if err != nil {
		logging.Logging(logging.ERR, "数据库连接出现错误,请检查数据库连接信息")
		logging.Logging(logging.ERR, err)
	}

	db.LogMode(true)
	db.DB().SetMaxIdleConns(0)

	return db
}

func (db *Database) Close() {
	DB.Self.Close()
	//db.Docker.Close()
}
