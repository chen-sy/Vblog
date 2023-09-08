package hello

import (
	"fmt"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		//TODO 默认欢迎语改为可配置
		name = "Hello,World!"
	}
	fmt.Fprintf(w, "%s\n", name)
}
