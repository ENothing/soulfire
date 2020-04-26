package db

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"soulfire/pkg/logging"
)

var R redis.Conn

func RedisInit()  {

	var err error
	address := viper.GetString("Redis.address")+":"+viper.GetString("Redis.port")
	R ,err = redis.Dial("tcp",address)

	if err != nil {
		fmt.Println(err.Error())
		logging.Logging(logging.ERR,err.Error())
	}

	defer R.Close()

}
