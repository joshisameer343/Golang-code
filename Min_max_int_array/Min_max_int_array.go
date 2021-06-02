package main

import "fmt"

func main() {

	a := []int{22, 45, 43, -44, 32, 128, -22, -55}
	fmt.Println(a)

	MinInt := a[0]
	MaxInt := a[0]

	for _, v := range a {

		if v > MaxInt {
			MaxInt = v
		}
		if (v) < MinInt {
			MinInt = (v)
		}

	}

	fmt.Println("Max:", MaxInt)
	fmt.Println("Min:", MinInt)
}
