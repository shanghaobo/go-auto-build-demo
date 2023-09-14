package main

import (
	"fmt"
	"go-auto-build-demo/model"
	"time"
)

func main() {
	fmt.Println("hello world")
	time.Sleep(10 * time.Second)
	model.InitDb()
}
