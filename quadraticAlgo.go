package main

import (
	"fmt"
	"math"
)

func quadraticEquation(a float64, b float64, c float64) {
	if a == 0 {
		linearEquation(b, c)
	} else {
		var delta = b*b - 4*a*c
		fmt.Println("delta = ", delta)
		if delta < 0 {
			fmt.Println("No solutions")
		} else if delta == 0 {
			var x = -b / (2 * a)
			fmt.Println("Only one solutionEnglishing - x = ", x)
		} else {
			var x1 = (-b - math.Sqrt(delta)) / (2 * a)
			var x2 = (-b + math.Sqrt(delta)) / (2 * a)
			fmt.Println("Two solutions - x1 = ", x1, " x2 = ", x2)
		}
	}
}
