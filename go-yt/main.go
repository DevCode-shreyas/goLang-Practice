package main

import "fmt"

// "bufio"

// "os"

// func main() {
// fmt.Println("Hello, World!")

// var name string = "shreyas"

// fmt.Println("Hello, " + name + "!")

// version := "version 1.0"

// fmt.Println(version)

// var dimension float64 = 22.22

// fmt.Println(dimension)

// var decided bool = true

// fmt.Println(decided)

// const pi = 3.14
// pi = 3.15 // cannot assign to pi
// fmt.Println(pi)

// var persion string = "shreyas bole"
// persion := "shreyas bole"
// fmt.Println(persion)

// capital letter is public is accessible outside the package and small letter is private and not accessible outside the package
// Public := "data is important"
// private := "data is private"

// fmt.Println(Public)
// fmt.Println(private)

// print

// age := 22
// name = "shreyas"

// fmt.Println("Hello, ", name, "!", "You are ", age, " years old.")

// print using format specifier

// fmt.Printf("Hello, %v! You are %v years old.\n", name, age)

// %v is used to print the value of the variable
// %T is used to print the type of the variable
// %f is used to print the float value - %.3f is used to print the float value upto 3 decimal places
// %d is used to print the integer value
// %s is used to print the string value

// how to take input from the user

// fmt.Println("Enter your name: ")

// reader := bufio.NewReader(os.Stdin)

// name, _ := reader.ReadString('\n')

// fmt.Println("Hello, " + name + "!")

// Functions in Go

// fmt.Println("Hello, World!")

// simpleFunction()

// ans := add(2, 3)
// fmt.Println(ans)
// }

// func simpleFunction() {
// 	fmt.Println("This is a simple function")
// }

// func add(a int, b int) (result int) {
// 	return a + b
// }

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		return 0, fmt.Errorf("cannot divide by zero")
// 	}
// 	return a / b, nil
// }

// func main() {
// 	fmt.Println("Hello, World!")

// 	ans, _ := divide(10, 4)
// 	fmt.Println(ans)
// }

// Arrays in Go

// func main() {
// 	fmt.Println("Hello, World!")

// var name [5]string

// name[0] = "shreyas"
// name[2] = "bole"

// fmt.Println(name)

// fmt.Println("Length of the array is: ", len(name))

// fmt.Println("First element of the array is: ", name[2])

// Default value of int,float is 0 , for string is " " , for bool is false for pointers or complex type is nil

// var price [5]int

// fmt.Println("price array is: ", price)
// }

// Slices in Go

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	numbers = append(numbers, 6)
	fmt.Println(numbers)
	fmt.Println("Length of the slice is: ", len(numbers))
}
