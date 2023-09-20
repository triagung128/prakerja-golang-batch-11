package main

import "fmt"

func main() {
	var a float32
	var b float32
	var h float32

	fmt.Print("Masukkan sisi atas : ")
	fmt.Scan(&a)

	fmt.Print("Masukkan sisi bawah : ")
	fmt.Scan(&b)

	fmt.Print("Masukkan tinggi : ")
	fmt.Scan(&h)

	result := calculateAreaTrapezoid(a, b, h)

	fmt.Printf("Hasil luas trapesium adalah %.1f", result)
}

func calculateAreaTrapezoid(a float32, b float32, h float32) float32 {
	return 0.5 * (a + b) * h
}
