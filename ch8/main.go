package main

import "fmt"

type Val interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func DoubleVal[T Val](v T) T {
	return v * 2
}

func main() {
	var x int
	var y float32
	var z float64
	x = 1
	y = 2.0
	z = 3.0
	fmt.Println(DoubleVal(x))
	fmt.Println(DoubleVal(y))
	fmt.Println(DoubleVal(z))
}
