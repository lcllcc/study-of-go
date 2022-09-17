// 比较+字符串和join字符串的性能区别

package main

import (
	"fmt"
	"time"
	"strings"
	"os"
)

func main() {
	var s string
	start := time.Now()
	for i := 0;i < len(os.Args);i++ {
		s = s + os.Args[i]  
	}
	cost := time.Since(start).Seconds()
	fmt.Println("s:", s)
	fmt.Println("+ string:", cost)

	start = time.Now()
	s = strings.Join(os.Args[0:], "")
	cost = time.Since(start).Seconds();
	fmt.Println("s:", s)
	fmt.Println("join string:", cost)
}