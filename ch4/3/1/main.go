package main

import "fmt"

func main() {
	doWork := func(strings <-chan string) <-chan interface{} {
		completed := make(chan interface{})
		go func() {
			defer fmt.Println("dowork exited.")
			defer close(completed)
			for s := range strings {
				// なにか面白い処理
				fmt.Println(s)
			}
		}()
		return completed
	}
	doWork(nil)
	// もう少し何らかの処理がここで行われる
	fmt.Println("Done.")
}
