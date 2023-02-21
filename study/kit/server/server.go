package server

import (
	"fmt"
	"golang/study/kit/server/endpoint"
	"golang/study/kit/server/service"
	"golang/study/kit/server/transport"
	"io"
	"net/http"
	"os"

	transporthttp "github.com/go-kit/kit/transport/http"
	kitlog "github.com/go-kit/log"
	"golang.org/x/time/rate"
)

func Server(port int) error {
	//初始化服务
	var userService = service.UserService{}
	var point = endpoint.UserEndPoint(userService)
	//添加日志中间件
	var logger kitlog.Logger
	{
		logger = kitlog.NewLogfmtLogger(os.Stdout)
		logger = kitlog.WithPrefix(logger, "mykit", "1.0")
		logger = kitlog.WithPrefix(logger, "time", kitlog.DefaultTimestampUTC) //加上前缀时间
		logger = kitlog.WithPrefix(logger, "caller", kitlog.DefaultCaller)     //加上前缀，日志输出时的文件和第几行代码
	}
	var loggerPoint = endpoint.LoggerMiddleware(logger)(point)
	//将服务绑定到节点添加限流中间件
	limit := rate.NewLimiter(1, 5) //附加限流 最大容量10个 每秒生成一个
	var userEndPoint = endpoint.RateLimit(limit)(loggerPoint)
	//添加自定义错误
	var options = []transporthttp.ServerOption{transporthttp.ServerErrorEncoder(transport.UserError)}
	//将节点绑定到服务器
	server := transporthttp.NewServer(userEndPoint, transport.UserDecodeRequest, transport.UserEncodeResponse, options...)
	//确定路由
	http.Handle("/user", server)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "连接正常")
	})
	//启动服务器
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
