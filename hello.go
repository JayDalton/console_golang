package main

import "fmt"

const PI float64 = 3.14

func main() {
	fmt.Println("Hello, World!")

	var Student1 string = "John"
	var Student2 = "Jane"
	number := 42

	fmt.Println(Student1)
	fmt.Println(Student2)
	fmt.Println(number)

	var a, b, c, d int = 1, 3, 5, 7
	fmt.Println(a, b, c, d)

	fmt.Println(PI)

	var arr1 = [3]int{3,5,7}
	arr2 := [5]int{2,4,6,8,0}
	arr3 := [5]int{1:10,2:20}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(len(arr3))
	fmt.Println(arr2[0])

	var cars = [4]string{"Volvo", "BMW", "Audi", "Ford"}
	fmt.Println(cars)


}
