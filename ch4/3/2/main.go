package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{},
		strings <-chan string,
	) <-chan interface{} {
		terminated := make(chan interface{})
		go func() {
			defer fmt.Println("dowork exited.")
			defer close(terminated)
			for {
				select {
				case s := <-strings:
					// なにか面白い処理
					fmt.Println(s)
				case <-done:
					return
				}
			}
		}()
		return terminated
	}
	done := make(chan interface{})
	terminated := doWork(done, nil)

	go func() { // 3
		// 1秒後に操作をキャンセルする
		time.Sleep(1 * time.Second)
		fmt.Println("Canceling doWork goroutine...")
		close(done)
	}()
	<-terminated // 4
	fmt.Println("Done.")
}
