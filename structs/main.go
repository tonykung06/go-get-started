package main

import (
	"fmt"
)

func main() {
	i1 := myCustomType{"testing"} //creating instance in local stack memory
	i2 := new(myCustomType)       //creating the instance in heap memory
	i2.field1 = "testing"

	fmt.Println(i1, i2)
}

type myCustomType struct {
	field1 string
}
