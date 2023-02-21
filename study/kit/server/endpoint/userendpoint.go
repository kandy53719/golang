package endpoint

import (
	"context"
	"errors"
	"golang/study/kit/server/service"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/log"
	"golang.org/x/time/rate"
)

/*
节点层：传输层需要用到的数据格式 负责request、response的转换，以及拦截相关
*/

type UserRequest struct {
	Id     int `json:"id"`
	Method string
}

type UserResponse struct {
	Result string `json:"result"`
}

type UserError struct {
	Code    int
	Message string
}

func (ur UserError) Error() string {
	return ur.Message
}

func NewUserError(code int, message string) error {
	return UserError{code, message}
}

// 用户服务节点
func UserEndPoint(us service.IUserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		if ur, ok := request.(UserRequest); !ok {
			return response, errors.New("非法请求")
		} else {
			//RESTful接口 这里做路由
			switch ur.Method {
			case http.MethodGet:
				return UserResponse{Result: us.GetName(ur.Id)}, nil
			case http.MethodPost:
				return UserResponse{Result: strconv.Itoa(us.GetAge(ur.Id))}, nil
			case http.MethodPut:
				return UserResponse{Result: us.GetName(ur.Id)}, nil
			case http.MethodDelete:
				return UserResponse{Result: strconv.Itoa(us.GetAge(ur.Id))}, nil
			default:
				return UserResponse{Result: us.GetName(ur.Id)}, nil
			}
		}
	}
}

// 限流中间件 用来包裹节点 提供限流服务
func RateLimit(limit *rate.Limiter) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			if !limit.Allow() { //确认当前是有用容量
				return nil, NewUserError(404, "当前请求数目过多，已限流")
			} else {
				return next(ctx, request)
			}
		}
	}
}

// 日志中间件
func LoggerMiddleware(logger log.Logger) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			var r = request.(UserRequest)
			logger.Log("method", r.Method, "id", r.Id)
			return next(ctx, request)
		}
	}
}
