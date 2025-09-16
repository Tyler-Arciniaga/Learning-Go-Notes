package main

import "fmt"

type Val interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func DoubleVal[T Val](v T) T {
	return v * 2
}

type Printable interface {
	fmt.Stringer
	~int | ~float64
}

type PrintableInt int

func (p PrintableInt) String() string {
	return fmt.Sprintf("Printable int type: %d\n", p)
}

type PrintableFloat float64

func (p PrintableFloat) String() string {
	return fmt.Sprintf("Printable float type: %f\n", p)
}

func PrintPrintable[T Printable](p T) {
	fmt.Println(p)
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
	fmt.Println("---------------")
	var i PrintableInt = 7
	var i2 PrintableFloat = 14.0
	PrintPrintable(i)
	PrintPrintable(i2)
}
