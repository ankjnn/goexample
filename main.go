package main

import "fmt"

func main() {
	fmt.Println("Hello, this is main opening")
	foo()
	fmt.Println("this is exit line")
}

func foo() {
	fmt.Println("this is from foo")
}
