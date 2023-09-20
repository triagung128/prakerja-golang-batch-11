package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&n)

	if isMultipleSeven(n) {
		fmt.Printf("%d adalah bilangan kelipatan 7", n)
	} else {
		fmt.Printf("%d adalah bukan bilangan kelipatan 7", n)
	}

}

func isMultipleSeven(n int) bool {
	return n%7 == 0
}
