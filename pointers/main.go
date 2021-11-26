package main

import "fmt"

func add(a *int, b *int, sum *int) {
	*sum = *a + *b
}

func printList(l ...int) {
	for _, i := range l {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}


func seq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	a := 10
	b := 20
	sum := 0

	add(&a, &b, &sum)

	fmt.Printf("Sum = %d\n\n", sum)

	l := []int{1, 2, 3, 3, 5}
	printList(l...)

	nextInt := seq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
