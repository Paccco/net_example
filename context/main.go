/**
@time : 2021/3/4 17:52
@auth : paco
*/
package main

import (
	context "context"
	"fmt"
	"sync"

	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
	go worker2(ctx)
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
	wg.Done()
}

func worker2(ctx context.Context) {
LOOP:
	for {
		fmt.Println("worker2")
		time.Sleep(time.Second)
		select {
		case <-ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}
}

// context.WithCancel()
func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}

func main() {
	//ctx, cancel := context.WithCancel(context.Background())
	//wg.Add(1)
	//go worker(ctx)
	//time.Sleep(time.Second * 3)
	//cancel() // 通知子goroutine结束
	//wg.Wait()
	//fmt.Println("over")

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	//for n := range gen(ctx) {
	//	fmt.Println(n)
	//	if n == 5 {
	//		break
	//	}
	//}

	d := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	defer cancel()
	select {
	case <-time.After(time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println(ctx.Err())

	}
}
