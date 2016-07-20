package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.GOMAXPROCS(4)) //printing previous setting
	ch := make(chan string, 2)
	doneCh := make(chan bool)
	go abcGen(ch)
	go printer(ch, doneCh)
	fmt.Println("This comes first")
	<-doneCh
}

func abcGen(ch chan string) {
	for l := byte('a'); l <= byte('z'); l++ {
		ch <- string(l)
		fmt.Println(l)
	}
	close(ch)
}

func printer(ch chan string, doneCh chan bool) {
	for l := range ch {
		fmt.Println(l)
	}
	doneCh <- true
}
