package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	q1And2()
	q3()
}

func q1And2() {
	s := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		s = append(s, rand.IntN(101))
	}

	for _, v := range s {
		switch {
		case v%2 == 0 && v%3 == 0:
			fmt.Println("Six!")
		case v%2 == 0:
			fmt.Println("Two!")
		case v%3 == 0:
			fmt.Println("Three!")
		default:
			fmt.Println("Never mind")
		}

	}
}

func q3() {
	var total int
	for i := 0; i < 10; i++ {
		total = total + i
		fmt.Println(total)
	}
	fmt.Println(total)
}
