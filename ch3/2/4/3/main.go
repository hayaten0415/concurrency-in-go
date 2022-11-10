package main

import (
	"sync"
)

func main() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceB.Do(initA) }
	onceA.Do(initA)
}
