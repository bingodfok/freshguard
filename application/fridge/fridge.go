package main

import (
	"fmt"
	"github.com/bingodfok/freshguard/application/fridge/internal/alc"
	"strconv"
)

// 冰箱相关管理
func main() {
	applicationContext, err := alc.NewApplicationContext()
	if err != nil {
		panic(err)
	}
	fmt.Printf(strconv.FormatBool(applicationContext.Config.Mysql.ShowSql))
}
