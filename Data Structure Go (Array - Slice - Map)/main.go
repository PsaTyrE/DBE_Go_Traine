package main

import (
	"fmt"
	"reflect"
)

func arrayA() {
	prim := [5]int{}
	countries := [5]string{}

	fmt.Println(reflect.ValueOf(prim).Kind())
	fmt.Println(reflect.ValueOf(countries).Kind())
}

func arrayB() {
	x := [5]int{1, 2, 5, 7, 3}
	y := [5]int{1, 6, 9, 7, 3}

	fmt.Println(x)
	fmt.Println(y)
}
func main() {
	arrayA()
	arrayB()
}
