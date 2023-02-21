package service

/*
服务层：专注于业务逻辑，就是我们的业务类、接口等相关信息存放。
*/

type IUserService interface {
	GetName(id int) string
	GetAge(id int) int
}

type UserService struct{}

func (us UserService) GetName(id int) (name string) {
	switch id {
	case 123:
		name = "张三"
	case 456:
		name = "李四"
	case 789:
		name = "王五"
	default:
		name = "匿名用户"
	}
	return
}

func (us UserService) GetAge(id int) (age int) {
	switch id {
	case 123:
		age = 18
	case 456:
		age = 19
	case 789:
		age = 20
	default:
		age = 0
	}
	return
}
