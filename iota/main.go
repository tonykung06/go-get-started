package main

import (
// "fmt"
)

const (
	first = iota
	second
)

const (
	third = iota //get reset
)

const (
	forth = 1 << (10 * iota)
	fifth
	sixth
)

func main() {
	println(first, second, third, forth, fifth, sixth)
}
