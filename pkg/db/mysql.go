package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"soulfire/pkg/config"
	"soulfire/pkg/logging"
)

type Database struct {
	Self *gorm.DB
	//Docker *gorm.DB
}

var (
	DB *Database
)

func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
	}
}

func GetSelfDB() *gorm.DB {

	mysqlConfig, _ := config.Cfg.GetSection("mysql")

	host := mysqlConfig.Key("HOST").String()
	port := mysqlConfig.Key("PORT").String()
	database := mysqlConfig.Key("DATABASE").String()
	user := mysqlConfig.Key("USER").String()
	password := mysqlConfig.Key("PASSWORD").String()

	fmt.Println(mysqlConfig)

	return connect(user, password, host+":"+port, database)
}

//func GetDockerDB() *gorm.DB {
//
//	mysql_config, _ := config.Cfg.GetSection("mysql")
//
//	host := mysql_config.Key("HOST").String()
//	port := mysql_config.Key("PORT").String()
//	database := mysql_config.Key("DATABASE").String()
//	user := mysql_config.Key("USER").String()
//	password := mysql_config.Key("PASSWORD").String()
//
//	return connect(user,password,host+":"+port,database)
//}

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
