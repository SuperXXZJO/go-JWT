package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"homework1/Router"
	"homework1/User"
)

var (
	DB *gorm.DB
)

func init(){
	//连接数据库
	db, err := gorm.Open("mysql", "root:root@/homework1?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	DB =db
	if err != nil {
		panic(err)
	}

	//创建表
	err =DB.AutoMigrate(&User.User{}).Error
	if err != nil {
		panic(err)
	}

}
func main(){
	Router.RUN()
}
