package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func swap1(a, b int) (int, int) {
	return b, a
}

func main() {
	a, b := 1, 2
	swap(&a, &b)
	fmt.Println(a, b)

	a1, b1 := swap1(a, b)
	fmt.Println(a1, b1)
}
