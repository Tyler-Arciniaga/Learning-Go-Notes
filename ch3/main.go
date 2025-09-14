package main

import "fmt"

/*
Write a program that defines a variable named greetings
of type slice of strings with the following values:
"Hello", "Hola", "नमस्कार", "こんにちは", and "Привіт".
Create a sub-slice containing the first two values,
a second subslice with the second, third, and fourth values,
and a third subslice with the fourth and fifth values. Print out all four slices.
*/

func main() {
	q1()
	q2()
	q3()
}

func q1() {
	s := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	s1 := s[:2]
	s2 := s[1:4]
	s3 := s[3:]
	fmt.Println(s)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}

func q2() {
	msg := "Hi 👩 and 👨"
	rune_msg := []rune(msg)
	fmt.Println(string(rune_msg[3]))
}

func q3() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}

	tyler := Employee{"Tyler", "Arciniaga", 1}
	p1 := Employee{firstName: "p1", lastName: "x", id: 2}
	var p2 Employee
	p2.firstName = "p2"
	p2.lastName = "y"
	p2.id = 3

	fmt.Println(tyler)
	fmt.Println(p1)
	fmt.Println(p2)
}
