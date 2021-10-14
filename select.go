package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		time.Sleep(4 * time.Second)
		close(c)
	}()

	go func() {
		// time.Sleep(3 * time.Second)
		time.Sleep(5 * time.Second)
		ch1 <- 3
	}()

	go func() {
		// time.Sleep(3 * time.Second)
		time.Sleep(5 * time.Second)
		ch2 <- 5
	}()

	fmt.Println("Blocking on read")

	// 随机选择就绪的通道执行
	select {
	case <- c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	case <- ch1:
		fmt.Printf("ch1 case \n")
	case <- ch2:
		fmt.Printf("ch2 case \n")

	/*default:
		fmt.Printf("default go...\n")*/
	}

}
