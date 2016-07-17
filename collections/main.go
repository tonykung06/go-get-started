package main

import (
	"fmt"
	"reflect"
)

func main() {
	myArray := [...]int{1, 2, 3}
	mySlice := myArray[:]
	mySlice[0] = 5555
	fmt.Println("sharing share array", myArray)
	fmt.Println("sharing share array", mySlice)
	mySlice = append(mySlice, 100)
	mySlice[1] = 666
	fmt.Println("not sharing share array", myArray)
	fmt.Println("not sharing share array", mySlice)

	fmt.Println("\n==================make slice with size and capacity===============")
	sizedSlice := make([]int, 50, 100)
	sizedSlice[0] = 21
	sizedSlice[2] = 23
	fmt.Println(sizedSlice)
	fmt.Println(len(sizedSlice))

	fmt.Println("\n==================uninitialized map and make map===============")
	var myMap map[string]string
	myMap2 := make(map[string]string)
	var uninitSlice []int
	var uninitArray [2]int
	fmt.Println(uninitSlice == nil)
	fmt.Println(reflect.TypeOf(uninitSlice))
	fmt.Println(reflect.TypeOf(uninitArray))
	fmt.Println(uninitSlice)
	fmt.Println(uninitArray)
	fmt.Println(myMap == nil)
	fmt.Println(myMap2 == nil)
	fmt.Println(myMap2["test"] == "")

	fmt.Println(reflect.TypeOf(14.))
}
