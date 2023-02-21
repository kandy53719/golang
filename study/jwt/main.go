package main

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	//加密测试
	var user = User{Id: 123, Name: "张三"} //测试加密对象
	var sec = []byte("123abc")           //加密使用到的密钥
	token, err := Des(user, sec)
	fmt.Printf("token: %v\n", token)
	fmt.Printf("err: %v\n", err)

	//解密测试
	u, err2 := UnDes(token, sec)
	fmt.Printf("u: %v\n", u)
	fmt.Printf("err2: %v\n", err2)
}

type User struct {
	Id   int
	Name string
	jwt.StandardClaims
}

type UserClaim struct {
	Uname              string `json:"username"`
	jwt.StandardClaims        //嵌套了这个结构体就实现了Claim接口
}

// 对称加密
func Des(user User, sec []byte) (token string, err error) {
	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, user).SignedString(sec) //加密得到密文

	return
}

// 对称解密
func UnDes(token string, sec []byte) (user User, err error) {
	//验证
	// var getToken, _ = jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
	// 	return sec, nil //这里是对称加密，所以只要有人拿到了这个sec就可以进行访问不安全
	// })
	var getToken *jwt.Token
	getToken, err = jwt.ParseWithClaims(token, &user, func(t *jwt.Token) (interface{}, error) {
		return sec, nil
	})
	if getToken.Valid {
		fmt.Println(getToken.Claims.(*User).Name)
	}
	return
}
