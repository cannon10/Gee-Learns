package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// net/http 包中用于注册 HTTP 路由的核心函数
	// 将指定的 URL 路径与处理函数关联起来，使得当服务器收到对应路径的请求时，自动调用该处理函数生成响应。
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
