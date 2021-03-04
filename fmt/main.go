package main

import (
	"fmt"
	"os"
)

func main() {

	//Print系列函数会将内容输出到系统的标准输出，区别在于Print函数直接输出内容，Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符。
	fmt.Print("在终端打印该信息。")
	name := "paco"
	fmt.Printf("我是：%s\n", name)
	fmt.Println("在终端打印单独一行显示")

	// Fprint 系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
	// 向标准输出写入内容
	_, _ = fmt.Fprintln(os.Stdout, "向标准输出写入内容")
	fileObj, err := os.OpenFile("./tmp/xx.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("打开文件出错，err:", err)
		return
	}
	// 向打开的文件句柄中写入内容
	_, _ = fmt.Fprintf(fileObj, "往文件中写如信息：%s", name)

	// Sprint系列函数会把传入的数据生成并返回一个字符串。

	s1 := fmt.Sprint("paco")
	age := 18
	s2 := fmt.Sprintf("name:%s,age:%d", name, age)
	s3 := fmt.Sprintln("paco")
	fmt.Println(s1, s2, s3)
}
