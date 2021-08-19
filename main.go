package main

import (
	"fmt"
	"math"
)

func main() {

	var a int
	var b int
	var c int

	fmt.Println("input a")
	fmt.Scan(&a)
	fmt.Println("input b")
	fmt.Scan(&b)
	fmt.Println("input c")
	fmt.Scan(&c)

	if a == 0 {
		if b == 0 {
			if c == 0 {
				fmt.Println("równanie tożsamościowe")
			} else {
				fmt.Println("równanie sprzeczne")
			}
		} else {
			var x float64 = float64(-c / b)
			fmt.Println(x)
		}
	} else {
		var Step1 = b * b
		var Step2 int = 4 * a
		var Step3 int = int(Step2) * c
		var delta int = int(Step1) - int(Step3)

		fmt.Println("delta = ", delta)

		if delta < 0 {
			fmt.Println("Brak rozwiązań")
		} else {
			if delta == 0 {
				fmt.Println("jedno rozwiązanie - wynik = ", delta)
			} else {
				var ax2 int = 2 * a

				var x1step1 int = -b
				var x1step2 float64 = math.Sqrt(float64(delta))
				var x1step3 float64 = float64(x1step1) - x1step2
				var x1 float64 = float64(x1step3) / float64(ax2)
				fmt.Println("x1 = ", x1)

				var x2step1 float64 = float64(-b) + x1step2
				var x2step2 = x2step1
				var x2 float64 = x2step2 / float64(ax2)
				fmt.Println("x2 = ", x2)
			}
		}
	}

}
