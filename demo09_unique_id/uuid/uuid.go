package main

import (
	"fmt"
	"github.com/google/uuid"
)

func main() {
	id := uuid.New()
	fmt.Println("生成唯一id成功")
	fmt.Println("uint32 id", id.ID(), id.Version())
	fmt.Println("string id", id.String())
}
