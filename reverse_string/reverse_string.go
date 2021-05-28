package main

import "fmt"

func reverse(s string) string {
	x := ""
	for i := len(s) - 1; i >= 0; i-- {
		x = x + string(s[i])
	}
	return x
}

func main() {
	x1 := reverse("abbbcc kjfle kefjsej")
	fmt.Println(x1)
}
