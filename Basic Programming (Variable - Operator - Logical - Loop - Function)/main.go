package main

import (
	"fmt"
)

func Asterik(n int) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("* ")
		}
		fmt.Println()
	}
}

func main() {
	// variable
	firstnama := "dimas"
	lastnama := "bayu"
	fmt.Println("full nama : ", firstnama+" "+lastnama)

	// operator
	x := 5
	b := 9

	hasil := x + b

	fmt.Println(hasil)

	// luas segitiga

	a := 10
	t := 20
	luasSegitiga := (a * t) / 2

	fmt.Println(luasSegitiga)

	// branch dan looping if, else, else if

	jam := 20

	if jam < 12 {
		fmt.Println("selamat pagi")
	} else if jam < 15 {
		fmt.Println("Selamat sore")
	} else {
		fmt.Println("selamat malam")
	}

	// branch dan looping swicth break

	today := 5

	switch today {
	case 1:
		fmt.Println("today is monday")
	case 2:
		fmt.Println("today is tuesday")
	default:
		fmt.Println("other day")
	}

	// looping
	for i := 0; i < 5; i++ {
		if i == 0 {
			continue
		}
		if i > 3 {
			break
		}
		fmt.Println(i + 1)
	}

	sentence := "hello world"

	for i := 0; i < len(sentence); i++ {
		fmt.Printf(string(sentence[i]) + "-")
	}
	fmt.Println("\n")

	for pos, char := range sentence {
		fmt.Printf("%d %c\n ", pos+1, char)
	}

	// Praktikum
	nilai := 80

	if nilai < 100 && nilai > 90 {
		fmt.Println("A")
	} else if nilai < 89 && nilai > 70 {
		fmt.Println("B")
	} else if nilai < 69 && nilai > 50 {
		fmt.Println("C")
	} else if nilai < 49 && nilai > 30 {
		fmt.Println("D")
	} else {
		fmt.Println("E")
	}

	row := 5
	Asterik(row)
}
