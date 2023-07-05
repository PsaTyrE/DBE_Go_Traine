package main

import (
	"fmt"
)

func generatePrimes(n int) []bool {
	// Membuat sebuah slice dengan semua elemen diinisialisasi ke true
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	// Menerapkan algoritma Sieve of Eratosthenes
	for p := 2; p*p <= n; p++ {
		if primes[p] == true {
			// Mengubah semua kelipatan dari p menjadi false
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	return primes
}

func main() {
	n := 13
	primes := generatePrimes(n)

	fmt.Println("Bilangan prima dari 2 hingga", n, "adalah:")
	for i := 2; i <= n; i++ {
		if primes[i] == true {
			fmt.Println(primes)
		}
	}
}
