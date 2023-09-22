package main

import "fmt"

func main() {
	fmt.Println(Mapping([]string{"asd", "qwe", "adi", "qwe", "qwe"}))
	fmt.Println(Mapping([]string{"asd", "qwe", "asd"}))
	fmt.Println(Mapping([]string{}))
}

func Mapping(array []string) map[string]int {
	arrTemp := []string{}
	mapTemp := map[string]int{}

	for i := 0; i < len(array); i++ {
		_, ok := mapTemp[array[i]]
		if !ok {
			mapTemp[array[i]] = 0
			arrTemp = append(arrTemp, array[i])
		}
	}

	for i := 0; i < len(arrTemp); i++ {
		total := 0
		for j := 0; j < len(array); j++ {
			if arrTemp[i] == array[j] {
				total++
			}
		}
		mapTemp[arrTemp[i]] = total
	}

	return mapTemp
}
