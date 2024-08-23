package main

import (
	"fmt"
	"sync/atomic"
)

/**
 * 实验：使用通道实现轮训10次打印
 */
func main() {
	var count = new(atomic.Int32)
	count.Store(1)
	con := make(chan struct{}, 1)
	cat := make(chan struct{}, 1)
	dog := make(chan struct{}, 1)
	fish := make(chan struct{}, 1)
	cat <- struct{}{} // 初始化
	go catP(cat, dog, con, count)
	go dogP(dog, fish, con, count)
	go fishP(fish, cat, con, count)
	defer func() {
		// close(con)
		close(cat)
		close(dog)
		close(fish)
	}()
	<-con
	// // 创建一个通道，用于接收信号
	// sigChan := make(chan os.Signal, 1)
	// // 将中断信号（SIGINT）发送到通道中
	// signal.Notify(sigChan, syscall.SIGINT)
	// // 从通道中接收信号
	// <-sigChan
}

func catP(cat, dog, con chan struct{}, count *atomic.Int32) {
	for {
		<-cat
		fmt.Println("cat")
		if count.Load() < 10 {
			count.Add(1)
		} else {
			con <- struct{}{}
			return
		}
		dog <- struct{}{}
	}
}

func dogP(dog, fish, con chan struct{}, count *atomic.Int32) {
	for {
		<-dog
		fmt.Println("dog")
		if count.Load() < 10 {
			count.Add(1)
		} else {
			con <- struct{}{}
			return
		}
		fish <- struct{}{}
	}
}

func fishP(fish, cat, con chan struct{}, count *atomic.Int32) {
	for {
		<-fish
		fmt.Println("fish")
		if count.Load() < 10 {
			count.Add(1)
		} else {
			con <- struct{}{}
			return
		}
		cat <- struct{}{}
	}
}
