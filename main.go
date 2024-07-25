package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, 你访问的URL路径是 %s", r.URL.Path)
}

func main() {
	http.HandleFunc("/", handler)  // 对根路径的请求进行处理
	http.ListenAndServe(":8888", nil) // 在8080端口启动服务
}
