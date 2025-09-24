package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed english_rights.txt
var EnglishTrans string

//go:embed french_rights.txt
var FrenchTrans string

//go:embed dutch_rights.txt
var DutchTrans string

func main() {
	if len(os.Args) > 1{
		switch os.Args[1] {
		case "e":
			fmt.Println(EnglishTrans)
		case "f":
			fmt.Println(FrenchTrans)
		case "d":
			fmt.Println(DutchTrans)
		default:
			fmt.Println(EnglishTrans)
		}
	}
}
