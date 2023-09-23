package main

import "fmt"

func main() {
	names := []string{}
	scores := []int{}

	for i := 1; i <= 5; i++ {
		var name string
		var score int

		fmt.Printf("Input %d Student's Name : ", i)
		fmt.Scanln(&name)
		fmt.Printf("Input %d Student's Score : ", i)
		fmt.Scanln(&score)

		names = append(names, name)
		scores = append(scores, score)
	}

	student := Student{
		Name:  names,
		Score: scores,
	}

	student.calculate()
}

type Student struct {
	Name  []string
	Score []int
}

func (s Student) calculate() {
	mapTemp := map[string]int{}

	for i, v := range s.Name {
		mapTemp[v] = s.Score[i]
	}

	totalScore := 0
	for _, v := range mapTemp {
		totalScore += v
	}
	average := totalScore / len(mapTemp)

	nameMinScore := s.Name[0]
	minScore := mapTemp[s.Name[0]]
	for k, v := range mapTemp {
		if v < minScore {
			nameMinScore = k
			minScore = mapTemp[k]
		}
	}

	nameMaxScore := s.Name[0]
	maxScore := mapTemp[s.Name[0]]
	for k, v := range mapTemp {
		if v > maxScore {
			nameMaxScore = k
			maxScore = mapTemp[k]
		}
	}

	fmt.Println()
	fmt.Println("Average Score : ", average)
	fmt.Printf("Min Score of Students : %s (%d)", nameMinScore, minScore)
	fmt.Println()
	fmt.Printf("Max Score of Students : %s (%d)", nameMaxScore, maxScore)
}
