package main

import "fmt"

func linearEquation(b float64, c float64) {
	if b == 0 {
		if c == 0 {
			fmt.Println("identity equation")
		} else {
			fmt.Println("contradictory equation")
		}
	} else {
		var x = -c / b
		fmt.Println("x = ", x)
	}
}
