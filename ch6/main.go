package main

import "fmt"

type Person struct {
	FirstName string
	Lastname  string
	Age       int
}

func MakePerson(firstName, lastName string, age int) Person {
	return Person{FirstName: firstName, Lastname: lastName, Age: age}
}

func MakePersonPointer(firstName, lastName string, age int) *Person {
	return &Person{FirstName: firstName, Lastname: lastName, Age: age}
}

func UpdateSlice(s []string, x string) {
	s[len(s)-1] = x
	fmt.Println(s)
}

func GrowSlice(s []string, x string) {
	s = append(s, x)
	fmt.Println(s)
}
func main() {
	p1 := MakePerson("tyler", "arciniaga", 20)
	fmt.Println(p1)
	p2 := MakePersonPointer("tyler", "arciniaga", 20)
	fmt.Println(p2)

	s := []string{"1", "2", "3", "4"}
	fmt.Println(s)
	UpdateSlice(s, "5")
	fmt.Println(s)
	fmt.Println("-------")
	s2 := []string{"1", "2", "3", "4"}
	fmt.Println(s2)
	GrowSlice(s2, "5")
	fmt.Println(s2)
}
