// 定义包名
package main

//引入fmt包
import "fmt"

//声明函数
//main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）
func main() {
	//声明变量
	var hello = "Hello, 小姐姐!"

	//语句 可以将字符串输出到控制台
	fmt.Println(hello)
}

//单行注释
/*
多
行
注
释
以上就是go语言的基本结构
*/
