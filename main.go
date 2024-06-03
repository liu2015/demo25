package main

import (
	"demo25/global"
	"demo25/initalize"
	"fmt"
)

func main() {
	// fmt.print("启动")
	fmt.Println("启动")
	global.Gorm()
	run := initalize.Router()
	run.Run()
}
