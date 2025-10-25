package main

import "fmt"

func main() {
	a := 0
	add(&a)
	fmt.Println(a)

	b := []int{1, 2, 3}
	mutip(b)
	fmt.Println(b)
}

func add(a *int) {
	*a += 10
}

func mutip(a []int) {
	for i := range a {
		a[i] *= 2
	}
}
