package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	// 目标 URL 80端口也可省略
	url := "http://127.0.0.1:80/ping"

	// 创建一个新的 HTTP GET 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("创建请求失败: %v", err)
	}

	// 使用默认的 HTTP 客户端发送请求
	client := &http.Client{}
	for i := 0; i < 50; i++ {
		resp, err := client.Do(req)
		if err != nil {
			log.Fatalf("发送请求失败: %v", err)
		}
		defer resp.Body.Close()

		// 检查响应状态码
		if resp.StatusCode != http.StatusOK {
			log.Fatalf("请求失败，状态码: %d", resp.StatusCode)
		}

		// 读取响应体
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("读取响应体失败: %v", err)
		}

		// 打印响应内容
		fmt.Println("响应内容:")
		fmt.Println(string(body))
		time.Sleep(time.Second * 2)
	}
}
