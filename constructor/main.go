package main

import (
	"fmt"
)

func main() {
	testing := newMyCustomType()
	testing.myMap["testing"] = "testing2"
	fmt.Println(testing.myMap["testing"])

}

func newMyCustomType() *myCustomType {
	instance := &myCustomType{}
	instance.myMap = make(map[string]string)
	return instance
}

type myCustomType struct {
	myMap map[string]string
}
