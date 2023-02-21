package transport

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"golang/study/kit/server/endpoint"
)

/*
传输层：负责与传输协议HTTP，GRPC，THRIFT等相关的逻辑
*/

// 对请求进行解码 将传递过来的url请求参数转换成节点层定义的格式
func UserDecodeRequest(ctx context.Context, request *http.Request) (interface{}, error) {
	//uri参数解析
	tmp := request.URL.Query().Get("id")
	var id int = 0
	if tmp != "" {
		id, _ = strconv.Atoi(tmp)
	}
	method := request.Method
	log.Printf(": %s %s %s %d\n", request.RemoteAddr, request.RequestURI, method, id)
	return endpoint.UserRequest{Id: id, Method: method}, nil
}

// 对响应进行编码
func UserEncodeResponse(c context.Context, rw http.ResponseWriter, response interface{}) error {
	rw.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(rw).Encode(response)
}

func UserError(ctx context.Context, err error, w http.ResponseWriter) {
	contentType, body := "text/plain; charset=utf-8", []byte(err.Error())
	var code = 500
	w.Header().Set("Content-Type", contentType)
	if ur, ok := err.(endpoint.UserError); ok {
		code = ur.Code
	}
	w.WriteHeader(code)
	w.Write(body)
}
