package main

import (
	"fmt"
	"github.com/jiebozeng/golangutils/convert"
	"github.com/jiebozeng/golangutils/mathutils"
	"io"
	"net/http"
)

func main() {
	targetUrl := "http://127.0.0.1:8188/red/"
	go func() {
		for i := 0; i < 1000; i++ {
			userId := mathutils.RandInterval(1, 15)
			redpackId := mathutils.RandInterval(1, 15)
			rsp, err := http.Get(targetUrl + convert.ToString(userId) + "/" + convert.ToString(redpackId))
			if err != nil {
				panic(err)
			}
			defer rsp.Body.Close()
			println(rsp.Status)
			body, _ := io.ReadAll(rsp.Body)
			fmt.Println(string(body))
		}
	}()
	fmt.Scanln()
}
