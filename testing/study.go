// Description: Go 语言学习测试代码
// go mod tidy 生成 go.sum 文件
// go mod init <module-name> 生成 go.mod 文件

package main

import (
	"fmt"
	"os"
	"runtime"
)

func add(a int, b int) int {
	return a + b
}
func operatingSystem() string {
	hostname, err := os.Hostname()
	// nil 是一个预先声明的标识符，用于表示零值、空值或者是一个空指针 无错误
	if err == nil {
		fmt.Println("主机名：", hostname)
	}
	gopath := os.Getenv("GOPATH")
	fmt.Println("go环境变量GOPATH: ", gopath)
	fmt.Println("系统架构: ", runtime.GOARCH)
	// 声明一个 runtime.MemStats 类型的变量 memstats
	// runtime.memsats 是一个结构体类型，用于存储内存分配的统计信息
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)
	fmt.Print("Alloc: ", memstats.Alloc/1024/1024, "MB")
	fmt.Print("\t TotalAlloc: ", memstats.TotalAlloc/1024/1024, "MB")
	fmt.Print("\t Sys: ", memstats.Sys/1024/1024, "MB")
	fmt.Printf("\t NumGC: %v \n", memstats.NumGC)

	return runtime.GOOS
}

// 测试函数
func test1() {
	// iota 常量生成器
	const (
		a = iota
		b
		c
		d = 1 << iota
		e
	)
	fmt.Println(a, b, c, d, e)
	// if 条件语句
	if f := d * e; f <= 100 {
		fmt.Println("f < 10 ")
	} else {
		fmt.Println("f = ", f)
	}
	// for 循环语句
	imax := 10
	for i := 0; i < imax; i++ {
		if i == 0 {
			fmt.Printf("i = %d ", i)
			continue
		}
		fmt.Printf(" %d ", i)
		if i == imax-1 {
			fmt.Println()
		}
	}
	// switch 语句
	switch os := operatingSystem(); os {
	case "darwin":
		fmt.Println("Mac OS")
	case "linux":
		fmt.Println("Linux")
	case "windows":
		fmt.Println("Windows")
	default:
		fmt.Println("Other OS")
	}
}

var value int = 1

func main() {
	a, b := 1, 2
	value = add(a, b)
	fmt.Println("a + b", "= ", value)
	bitvalue := 1
	fmt.Println("位运算左移 ", bitvalue<<1)
	fmt.Println("位运算右移 ", bitvalue>>1)
	test1()
}
