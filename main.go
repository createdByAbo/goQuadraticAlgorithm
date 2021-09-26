package main

import "fmt"

func main() {
	fmt.Println("Enter a value:")
	var a int
	fmt.Scanln(&a)
	fmt.Println("Enter b value:")
	var b int
	fmt.Scanln(&b)
	fmt.Println("Enter c value:")
	var c int
	fmt.Scanln(&c)

	quadraticEquation(float32(a), float32(b), float32(c))
}
