package leetcodeTests

import (
	"testing"
	"fmt"
)

func TestChannel(t *testing.T) {
	c := make(chan string)
	go a("Hello", c)
	go b("world", c)
	fmt.Println("Main end!")
}

func a(text string, c chan string) {
	fmt.Println("start function a, and wait to receive a data from channel")
	t := <- c
	fmt.Println("function a receive a data from channel")
	fmt.Println(t)
	fmt.Println(text)
}

func b(text string, c chan string) {
	fmt.Println("start function b, and send a data to channel")
	c <- "hi"
	fmt.Println(text)
}
