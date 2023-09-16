package main

import "fmt"

func main() {
	fmt.Println("Hello world")
	var name string = "I am a string"
	fmt.Println(name)
	var a int
	const b int = 2
	const c int = 3
	myNums(a, b, c)
	myArray()
	mySlices()
}

func mySlices() {
	slice := []int{}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	powerslice := []string{"Slices", "Are", "Powerful"}
	fmt.Println(len(powerslice))
	fmt.Println(cap(powerslice))
	fmt.Println(powerslice)
}

func myArray() {
	var array1 = [3]int{1, 2, 3}
	array2 := [5]int{4, 5, 6, 7, 8}
	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(array1)
	fmt.Println((array2))
	fmt.Println(cars[1])
	// declaring uninitialized arrays
	var UninitArray = [5]int{}
	var PartialUninitArray = [5]int{1, 2}
	fmt.Println(UninitArray)
	fmt.Println(PartialUninitArray)
	//init only specific indexes of an array
	var InitArrayIndex = [5]int{1: 10, 3: 34}
	fmt.Println(InitArrayIndex)
	//printing length
	fmt.Println(len(UninitArray))
}

func myNums(a int, b int, c int) {
	a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
