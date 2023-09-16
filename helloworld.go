package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	var name string = "I am a string"
	fmt.Println(name)
	var a int
	var b int = 2
	var c int = 3
	myNums(a, b, c)
}

func myNums(a int, b int, c int) {
	a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
