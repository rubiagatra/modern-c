package main

import "fmt"

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {
	a := 21
	b := 17

	swap(&a, &b)
	fmt.Printf("main: a = %d, b = %d\n", a, b)
}
