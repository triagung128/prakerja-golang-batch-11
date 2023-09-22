package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(AppearedOnce("1234123"))
	fmt.Println(AppearedOnce("76523752"))
	fmt.Println(AppearedOnce("12345"))
	fmt.Println(AppearedOnce("1122334455"))
	fmt.Println(AppearedOnce("0872504"))
}

func AppearedOnce(numString string) []int {
	arrTemp := []int{}
	result := []int{}

	for i := 0; i < len(numString); i++ {
		numInt, _ := strconv.Atoi(string(numString[i]))
		arrTemp = append(arrTemp, numInt)
	}

	for i := 0; i < len(arrTemp); i++ {
		isTwin := false
		for j := 0; j < len(arrTemp); j++ {
			if i == j {
				continue
			}
			if arrTemp[i] == arrTemp[j] {
				isTwin = true
				break
			}
		}
		if !isTwin {
			result = append(result, arrTemp[i])
		}
	}

	return result
}
