package main

import (
	"fmt"
	"github.com/bingodfok/freshguard/application/fridge/internal/context"
	"strconv"
)

// 冰箱相关管理
func main() {
	applicationContext, err := context.NewApplicationContext()
	if err != nil {
		panic(err)
	}
	fmt.Printf(strconv.FormatBool(applicationContext.Config.Mysql.ShowSql))
}
