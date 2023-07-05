package main

import "fmt"

func Praktek(A, B string) string {
	lenA := len(A)
	lenB := len(B)
	var commonSubstring string

	for i := 0; i < lenA; i++ {
		for j := 0; j < lenB; j++ {
			if A[i] == B[j] {
				substring := string(A[i])
				k := 1
				for (i+k) < lenA && (j+k) < lenB && A[i+k] == B[j+k] {
					substring += string(A[i+k])
					k++
				}
				if len(substring) > len(commonSubstring) {
					commonSubstring = substring
				}
			}
		}
	}

	return commonSubstring
}

func main() {
	A := "AKA"
	B := "AKASHI"
	C := "KANG"
	D := "KANGGORO"
	result := Praktek(A, B)
	result1 := Praktek(C, D)
	fmt.Println(result)
	fmt.Println(result1)
}
