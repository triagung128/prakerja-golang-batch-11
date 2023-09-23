package main

import "fmt"

func main() {
	var a1, a2, a3, a4, a5, a6, min, max int
	fmt.Scan(&a1)
	fmt.Scan(&a2)
	fmt.Scan(&a3)
	fmt.Scan(&a4)
	fmt.Scan(&a5)
	fmt.Scan(&a6)
	min, max = getMinMax(&a1, &a2, &a3, &a4, &a5, &a6)
	fmt.Println("Nilai min : ", min)
	fmt.Println("Nilai max : ", max)
}

func getMinMax(numbers ...*int) (min int, max int) {
	minNum := *numbers[0]
	for _, v := range numbers {
		if *v < minNum {
			minNum = *v
		}
	}

	maxNum := *numbers[0]
	for _, v := range numbers {
		if *v > maxNum {
			maxNum = *v
		}
	}

	return minNum, maxNum
}
