package main

import (
	"fmt"
)

func main() {
	emp := NewEnhancedMessagePrinter("foo")
	emp.message2 = "bar"
	emp.printMessage()
}

func NewEnhancedMessagePrinter(msg string) *enhancedMessagePrinter {
	return &enhancedMessagePrinter{messagePrinter{message: msg}}
}

type messagePrinter struct {
	message  string
	message2 string
}

func (mp *messagePrinter) printMessage() {
	fmt.Println(mp.message, mp.message2)
}

type enhancedMessagePrinter struct {
	messagePrinter
}
