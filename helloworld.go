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
	var array [4]string = myArray()
	sliceFromArray(array)
	mySlices()
	copyArray()
}

func copyArray() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("numbers = %v\n", numbers)
	fmt.Printf("length = %d\n", len(numbers))
	fmt.Printf("capacity = %d\n", cap(numbers))

	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
	fmt.Printf("length = %d\n", len(numbersCopy))
	fmt.Printf("capacity = %d\n", cap(numbersCopy))
}

func sliceFromArray(array [4]string) {
	mySlice := array[1:3]
	fmt.Printf("mySlice = %v\n", mySlice)
	fmt.Printf("length = %d\n", len(mySlice))
	fmt.Printf("capacity = %d\n", cap(mySlice))
}

func mySlices() {
	slice := []int{}
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	powerslice := []string{"Slices", "Are", "Powerful"}
	fmt.Println(len(powerslice))
	fmt.Println(cap(powerslice))
	fmt.Println(powerslice)

	makeSlice := make([]int, 5, 10)

	fmt.Printf("mySlice = %v\n", makeSlice)
	fmt.Printf("length = %d\n", len(makeSlice))
	fmt.Printf("capacity = %d\n", cap(makeSlice))

	makeSlice = append(makeSlice, 29, 100)

	fmt.Printf("mySlice = %v\n", makeSlice)
	fmt.Printf("length = %d\n", len(makeSlice))
	fmt.Printf("capacity = %d\n", cap(makeSlice))

	halfSlice := []int{1, 2, 3}
	otherHalfSlice := []int{4, 5, 6}
	myFullSlice := append(halfSlice, otherHalfSlice...)

	fmt.Printf("mySlice = %v\n", myFullSlice)
	fmt.Printf("length = %d\n", len(myFullSlice))
	fmt.Printf("capacity = %d\n", cap(myFullSlice))

}

func myArray() [4]string {
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

	return cars
}

func myNums(a int, b int, c int) {
	a = 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
