package main

import (
	"fmt"
)

func main() {

	var a = []int{1, 2, 44, 443}

	var b = make([]int, 4, 10)

	_ = copy(b, a)

	fmt.Println(a)
	fmt.Println(b)

	a[0] = 21

	fmt.Println(a)
	fmt.Println(b)
}
