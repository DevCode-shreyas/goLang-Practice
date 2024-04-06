package main

import "fmt"

func main() {
	// fmt.Println("Hello, World!")

	var name string = "shreyas"

	fmt.Println("Hello, " + name + "!")

	version := "version 1.0"

	fmt.Println(version)

	var dimension float64 = 22.22

	fmt.Println(dimension)

	var decided bool = true

	fmt.Println(decided)

	const pi = 3.14
	// pi = 3.15 // cannot assign to pi
	fmt.Println(pi)

	// var persion string = "shreyas bole"
	persion := "shreyas bole"
	fmt.Println(persion)

	// capital letter is public is accessible outside the package and small letter is private and not accessible outside the package
	Public := "data is important"
	private := "data is private"

	fmt.Println(Public)
	fmt.Println(private)

	// print

	age := 22
	name = "shreyas"

	fmt.Println("Hello, ", name, "!", "You are ", age, " years old.")

	// print using format specifier

	fmt.Printf("Hello, %v! You are %v years old.\n", name, age)

	// %v is used to print the value of the variable
	// %T is used to print the type of the variable
	// %f is used to print the float value - %.3f is used to print the float value upto 3 decimal places
	// %d is used to print the integer value
	// %s is used to print the string value
}
