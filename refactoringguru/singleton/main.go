package main

import "fmt"

func main() {

	for i := 0; i < 30; i++ {
		// 5. 检查客户端代码， 将对单例的构造函数的调用替换为对其静态构建方法的调用。
		go getInstance()
	}

	// Scanln is similar to Scan, but stops scanning at a newline and
	// after the final item there must be a newline or EOF.
	fmt.Scanln()
}
