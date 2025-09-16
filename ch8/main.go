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
	fmt.Println("---------------")
	l := &List[int]{}
	l.Add(5)
	l.Add(10)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))

	l.Insert(100, 0)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	l.Insert(200, 1)
	fmt.Println(l.Index(5))
	fmt.Println(l.Index(10))
	fmt.Println(l.Index(200))
	fmt.Println(l.Index(20))
	fmt.Println(l.Index(100))

	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(300, 10)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Add(400)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}

	l.Insert(500, 6)
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		fmt.Println(curNode.Value)
	}
}
