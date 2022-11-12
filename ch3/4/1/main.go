package main

func main() {
	var c1, c2 <-chan interface{}
	var c3 chan<- interface{}
	select {
	case <-c1:
		//なにか
	case <-c2:
		//なにか
	case c3 <- struct{}{}:
		//なにか
	}
}
