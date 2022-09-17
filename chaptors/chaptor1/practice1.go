// 打印命令的名字
package main

import (
	"fmt"
	"os"
)

func echo() {
	fmt.Println(os.Args[0])
}

func main() {
	echo()
}