package main

import (
	"fmt"
	"github.com/bingodfok/freshguard/application/resources/internal/alc"
	"github.com/bingodfok/freshguard/application/resources/internal/handler"
)

// 外部通用资源(文件、短信等)管理
func main() {
	ctx, err := alc.NewApplicationContext()
	if err != nil {
		fmt.Println(err)
		return
	}
	ser := handler.NewHttpServer(ctx)
	ser.Run()
}
