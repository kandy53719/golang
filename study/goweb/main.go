package main

import (
	"fmt"
	"net/http"
)

func testHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test,world", r.URL.Path)
}

func helloHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello,world", r.URL.Path)
}

func main() {
	// http.HandleFunc("/hello", helloHandle)
	// http.HandleFunc("/test", testHandle)
	// http.ListenAndServe(":8080", nil)

	//var myHandler = MyHandler{}
	// var server = http.Server{
	// 	Addr:    ":8080",
	// 	Handler: &MyHandler{},
	// }
	// server.ListenAndServe()

	mux := http.NewServeMux()
	mux.Handle("/api/", &MyHandler{})
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Welcome to the home page!")
	})
	http.ListenAndServe(":8080", mux)
}

type MyHandler struct {
}

func (mh *MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello,world", r.URL.Path)
}
