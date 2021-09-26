package main

import "fmt"

func main() {
	fmt.Println("Enter a value:")
	var a float64
	fmt.Scanln(&a)
	fmt.Println("Enter b value:")
	var b float64
	fmt.Scanln(&b)
	fmt.Println("Enter c value:")
	var c float64
	fmt.Scanln(&c)

	quadraticEquation(a, b, c)
}
