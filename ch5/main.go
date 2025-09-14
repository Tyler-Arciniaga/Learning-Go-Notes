package main

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func main() {
	q1()
	//fileLen("2")
	helloPrefix := prefixer("Hey")
	fmt.Println(helloPrefix("Tyler"))
	fmt.Println(helloPrefix("Mom"))
}

func q1() {
	addFunc := func(x, y int) (int, error) { return x + y, nil }
	subFunc := func(x, y int) (int, error) { return x - y, nil }
	multFunc := func(x, y int) (int, error) { return x * y, nil }
	divFunc := func(x, y int) (int, error) {
		if y == 0 {
			return 0, errors.New("division by zero")
		}
		return x / y, nil
	}

	opMap := map[string]func(int, int) (int, error){
		"+": addFunc,
		"-": subFunc,
		"*": multFunc,
		"/": divFunc,
	}

	add := opMap["+"]
	div := opMap["/"]
	fmt.Println(add(1, 2))
	fmt.Println(div(1, 2))
	fmt.Println(div(1, 0))

}

func fileLen(n string) (int, error) {
	f, err := os.Open(n)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	data := make([]byte, 2048)
	var numBytes int
	for {
		count, err := f.Read(data)
		numBytes += count
		if err != nil {
			if err != io.EOF {
				return 0, err
			}
			break
		}
	}
	return numBytes, nil
}

func prefixer(p string) func(string) string {
	return func(s string) string {
		return p + " " + s
	}
}
