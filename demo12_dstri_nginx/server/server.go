package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请输入端口号")
		return
	}
	// 获取启动时输入的端口号
	port := os.Args[1]
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong from port: "+port)
	})
	fmt.Println("http server start,listen port ", port)
	http.ListenAndServe(":"+port, nil)
}
