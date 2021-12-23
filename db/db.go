package db

import (
	"fmt"
	"gin.forum.com/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

var engine *gorm.DB

type conn struct {
	DbHost string
	DbUser string
	DbPass string
	DbPort	string
	DbName	string
}
// MyConf 库选择
func InitDB(MyConf string) error {
	var err error
	conf := config.GetConfig()

	DBconn := conn{}

	switch MyConf {
	case "db_default": //默认数据库
		DBconn.DbHost = conf.DBConfig.DbHost
		DBconn.DbUser = conf.DBConfig.DbUser
		DBconn.DbPass = conf.DBConfig.DbPassword
		DBconn.DbPort = conf.DBConfig.DbPort
		DBconn.DbName = conf.DBConfig.DbName
	case "ads_config": //测试库1
		DBconn.DbHost = conf.ADSConfig.DbHost
		DBconn.DbUser = conf.ADSConfig.DbUser
		DBconn.DbPass = conf.ADSConfig.DbPassword
		DBconn.DbPort = conf.ADSConfig.DbPort
		DBconn.DbName = conf.ADSConfig.DbName
	}

	fmt.Println(MyConf)

	engine, err = gorm.Open("mysql", DBconn.DbUser+":"+DBconn.DbPass+"@tcp("+DBconn.DbHost+":"+DBconn.DbPort+")/"+DBconn.DbName+"?charset=utf8")
	if err != nil {
		err = errors.Wrap(err, "初始化数据库失败")
		return err
	}
	defer engine.Close()
	return nil
}



func GetDb(MyConf string) *gorm.DB  {
	InitDB(MyConf)
	return engine
}