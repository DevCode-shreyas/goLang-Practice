package main

import "fmt"

var pl = fmt.Println

func main() {
	var vName string = "World"
	pl("Hello, " + vName)

	v1, v2, v3 := 1, 2, 3
	pl(v1, v2, v3)
}
