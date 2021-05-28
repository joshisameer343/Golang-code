package main

import (
	"fmt"
)

func main() {

	var a = [][]int{{1, 2, 44, 443}, {3, 4, 5, 88}, {1, 2, 44, 443}}

	for _, v := range a {

		for _, b := range v {
			fmt.Println(b)
		}
	}
}
