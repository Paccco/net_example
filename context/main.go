/**
@time : 2021/3/4 17:52
@auth : paco
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var exit bool

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Nanosecond)
		if exit {
			break
		}
	}
	// 接收外部命令退出
	wg.Done()
}

func main() {
	wg.Add(1)
	go worker()
	time.Sleep(time.Second * 3)
	exit = true
	// 如何结束子goroutine
	wg.Wait()
	fmt.Println("over")
}
