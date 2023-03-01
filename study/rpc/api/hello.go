package api

import "fmt"

type Hello struct{}

// req 用来接收远程参数，res用来返回对方值,必须可序列化
func (h Hello) Hello(req string, res *string) error {
	*res = "你好，" + req
	return nil
}

// 以下是多值参数及返回
type UserRequest struct {
	Name    string
	Message string
}

type UserResponse struct {
	Code    int
	Message string
}

func (h Hello) HelloUser(req UserRequest, res *UserResponse) error {
	fmt.Println("客户端：", req.Name, req.Message)
	*res = UserResponse{
		Code:    200,
		Message: "欢迎:" + req.Name,
	}
	return nil
}
