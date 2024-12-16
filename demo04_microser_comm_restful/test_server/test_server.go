package main

import (
	"fmt"
	"github.com/jiebozeng/golangutils/convert"
	"github.com/jiebozeng/golangutils/mathutils"
	"io"
	"net/http"
	"net/url"
	"time"
)

func main() {
	targetUrl := "http://127.0.0.1:8188/user/"
	//get user by userId
	go func() {
		for i := 0; i < 100; i++ {
			userId := mathutils.RandInterval(1, 15)
			rsp, err := http.Get(targetUrl + "info/" + convert.ToString(userId))
			if err != nil {
				panic(err)
			}
			defer rsp.Body.Close()
			println(rsp.Status)
			body, _ := io.ReadAll(rsp.Body)
			fmt.Println(string(body))
			time.Sleep(time.Second * 5)
		}
	}()

	u, _ := url.ParseRequestURI(targetUrl + "list")
	data := url.Values{}
	data.Set("pageNum", "1")
	data.Set("pageSize", "5")

	u.RawQuery = data.Encode()
	rspList, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	defer rspList.Body.Close()
	println(rspList.Status)
	bodyList, _ := io.ReadAll(rspList.Body)
	fmt.Println(string(bodyList))
	fmt.Scanln()
}
