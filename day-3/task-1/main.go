package main

import "fmt"

func main() {
	fmt.Println(ArrayMerge([]string{"king", "devil jin", "akuma"}, []string{"eddie", "steve", "geese"}))
	fmt.Println(ArrayMerge([]string{"sergei", "jin"}, []string{"jin", "steve", "bryan"}))
	fmt.Println(ArrayMerge([]string{"alisa", "yoshimitsu"}, []string{"devil jin", "yoshimitsu", "alisa", "law"}))
	fmt.Println(ArrayMerge([]string{}, []string{"devil jin", "sergei"}))
	fmt.Println(ArrayMerge([]string{"hwoarang"}, []string{}))
	fmt.Println(ArrayMerge([]string{}, []string{}))
}

func ArrayMerge(arrayA, arrayB []string) []string {
	if len(arrayA) == 0 && len(arrayB) == 0 {
		return []string{}
	}

	if len(arrayA) == 0 {
		return arrayB
	}

	if len(arrayB) == 0 {
		return arrayA
	}

	mapTemp := make(map[string]bool)

	for _, item := range arrayA {
		mapTemp[item] = true
	}

	for _, item := range arrayB {
		if !mapTemp[item] {
			arrayA = append(arrayA, item)
		}
	}

	return arrayA
}
