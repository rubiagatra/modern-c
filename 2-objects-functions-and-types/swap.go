package main

import "fmt"

func swap(a, b *int) {
	ta := *a
	tb := *b
	*a = tb
	*b = ta
}

func main() {
	a := 21
	b := 17

	swap(&a, &b)
	fmt.Printf("main: a = %d, b = %d\n", a, b)
}
