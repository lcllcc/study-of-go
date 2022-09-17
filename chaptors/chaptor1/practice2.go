// 打印运行命令时指定的参数
package main

import (
	"fmt"
	"os"
)

func main()  {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
}