package main

import (
	"fmt"
	"math"
)

func main() {
	ex1()
	fmt.Println("----")
	ex2()
	fmt.Println("----")
	ex3()
}

func ex1() {
	i := 20
	f := float64(i)
	fmt.Println(i)
	fmt.Println(f)
}

func ex2() {
	const value = 10
	i := value
	var f float64 = value

	fmt.Println(i)
	fmt.Println(f)
}

func ex3() {
	var b byte = math.MaxUint8
	var smallI int32 = math.MaxInt32
	var bigI uint64 = math.MaxUint64

	b += 1
	smallI += 1
	bigI += 1

	fmt.Println(b)
	fmt.Println(smallI)
	fmt.Println(bigI)

}
