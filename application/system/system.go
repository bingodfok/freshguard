package main

import (
	"github.com/bingodfok/freshguard/application/system/internal/alc"
)

// 系统基础管理

func main() {
	_, err := alc.NewApplicationContext()
	if err != nil {
		panic(err)
	}
}
