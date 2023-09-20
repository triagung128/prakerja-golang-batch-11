package main

import (
	"fmt"
	"math"
)

func main() {
	var n int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&n)

	if isPrima(n) {
		fmt.Printf("%d adalah bilangan prima", n)
	} else {
		fmt.Printf("%d adalah bukan bilangan prima", n)
	}
}

func isPrima(n int) bool {
	if n <= 1 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
