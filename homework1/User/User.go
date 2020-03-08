package User

import (
	"github.com/jinzhu/gorm"
)
var (
	DB *gorm.DB
)

type User struct {
	gorm.Model
	Username string
	password string
	Sign     string
}
//添加用户 （注册）
func CreateNew (mod *User)error{
	u :=User{
		Username: mod.Username,
		password: mod.password,

	}
	DB.Create(&u)

	return nil
}

//查询用户 登录
func SelectUser(mod *User) User{
	Res := User{}
	DB.Where("Username=?",mod.Username).First(&Res)
	return Res
}

//根据用户名查找用户
func SelectUserByUsername(username string) User{
	Res := User{}
	DB.Where("Username=?",username).First(&Res)
	return Res
}

//修改用户信息
func Update (mod *User) User {
	Res :=User{}
	DB.Model(&Res).Updates(User{Username: mod.Username,password:mod.password,Sign:mod.Sign})
	return Res
}