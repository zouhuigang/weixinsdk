package main

import "fmt"

func defer_call() {
	defer func() { fmt.Println("a") }()
	defer func() { fmt.Println("b") }()
	defer func() { fmt.Println("c") }()
	panic("异常")
}
func main() {
	defer_call()
}
