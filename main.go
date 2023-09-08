package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 初始化配置
	//cfg := config.NewConfig()
	// 绑定路由
	http.HandleFunc("/hello", helloHandler)
	// 启动服务
	err := http.ListenAndServe(":8080", nil) //http://127.0.0.1:8080/hello
	if err != nil {
		log.Println(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		//TODO 默认欢迎语改为可配置
		name = "Hello,World!"
	}
	fmt.Fprintf(w, "%s\n", name)
}
