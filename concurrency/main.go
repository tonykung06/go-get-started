package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println(runtime.GOMAXPROCS(2)) //printing previous setting
	go abcGen()
	fmt.Println("This comes first")
	time.Sleep(10 * time.Millisecond)
}

func abcGen() {
	for l := byte('a'); l <= byte('z'); l++ {
		go fmt.Println(string(l))
	}
}
