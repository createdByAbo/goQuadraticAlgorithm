package main

import "fmt"

func quadraticEquation(a float32, b float32, c float32) {
	if a == 0 {
		linearEquation(b, c)
	} else {
		fmt.Println(a, b, c)
	}
}
