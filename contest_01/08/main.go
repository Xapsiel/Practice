package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	if a < b {
		a, b = b, a
	}
	if a < c {
		a, c = c, a
	}
	fmt.Println(a)
}
