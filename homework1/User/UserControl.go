package User

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/labstack/echo"
	"strconv"
	"strings"
	"time"
)
//注册
func Signup (c echo.Context)error{
	inf := User{}
	err := c.Bind(&inf)
	if err != nil {
		return c.JSON(300,err.Error())
	}
	if inf.password == ""{
		return c.JSON(300,"密码不能为空")
	}
	err = CreateNew(&inf)
	if err != nil{
		return err
	}
	return c.JSON(200,"注册成功")
}

//登录
func Login (c echo.Context)error{
	inf := User{}
	err := c.Bind(&inf)
	if err != nil {
		return err
	}
	mod := SelectUser(&inf)
	if mod.password != inf.password {
		return c.JSON(300,"密码错误" )
	}
	//生成token
	token := Ceatetoken(&mod)
	return c.JSON(200,token)
}


//查询用户信息
func FindUser (c echo.Context) error {
	inf := c.Param("username")
	mod := SelectUserByUsername(inf)
	return c.JSON(200,mod)

}

//修改用户信息
func UpdateUser (c echo.Context)error{
	inf := User{}
	err := c.Bind(&inf)
	if err != nil {
		return err
	}
	mod := Update(&inf)
	return c.JSON(200,mod)
}

//jwt

type Jwt struct {
}

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

func NewHeader() Header {
	return Header{
		Alg: "HS256",
		Typ: "JWT",
	}
}

type Payload struct {
	Iss      string `json:"iss"`
	Exp      string `json:"exp"`
	Iat      string `json:"iat"`
	Username string `json:"username"`
	Uid      uint
}

//生成token
func Ceatetoken (mod *User)string{
	header :=NewHeader()
	payload := Payload{
		Iss:      "liuxinyu",
		Exp:      strconv.FormatInt(time.Now().Add(3*time.Hour).Unix(), 10),
		Iat:      strconv.FormatInt(time.Now().Unix(), 10),
		Username: mod.Username,
		Uid:      mod.ID,
	}
	//生成json字符串
	h, _ := json.Marshal(header)
	p, _ := json.Marshal(payload)
	//生成
	headerbase64 :=base64.StdEncoding.EncodeToString(h)
	payloadbase64 :=base64.StdEncoding.EncodeToString(p)
	res := strings.Join([]string{headerbase64, payloadbase64}, ".")
	//加盐
	key := "123"
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(res))
	s := mac.Sum(nil)
	//生成签名
	signature :=base64.StdEncoding.EncodeToString(s)
	//生成token
	token := res + "." + signature
	return token
}

//验证token
func Checktoken(token string) error{
	arr := strings.Split(token, ".")
	if len(arr) != 3 {
		err := errors.New("token error")
		return err
	}
	_, err := base64.StdEncoding.DecodeString(arr[0])
	if err != nil {
		err := errors.New("token error")
		return err
	}
	pay, err := base64.StdEncoding.DecodeString(arr[1])
	if err != nil {
		err = errors.New("token error")
		return err
	}
	sign, err := base64.StdEncoding.DecodeString(arr[2])
	if err != nil {
		err = errors.New("token error")
		return err
	}

	res := arr[0] + "." + arr[1]
	key := []byte("redrock")
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(res))
	s := mac.Sum(nil)
	if res := bytes.Compare(sign, s); res != 0 {
		fmt.Println("test")
		err = errors.New("token error")
		return err
	}

	var payload Payload
	json.Unmarshal(pay,&payload)
	uid :=payload.Uid
	username :=payload.Username
	result := User{}
	result.Username = username
	result.ID = uid
	return nil
}