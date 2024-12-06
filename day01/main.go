package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	util "github.com/yattoni/advent-of-code-2024"
)

func main() {
	lines := util.ReadFileWithSpacesToLines(util.INPUT)

	list1 := make([]int, len(lines))
	list2 := make([]int, len(lines))

	for i, l := range lines {
		numStrs := strings.Fields(l)
		list1[i] = util.MustAtoi(numStrs[0])
		list2[i] = util.MustAtoi(numStrs[1])
	}
	fmt.Println("list1", list1)
	fmt.Println("list2", list2)

	slices.Sort(list1)
	slices.Sort(list2)

	fmt.Println("list1 sorted", list1)
	fmt.Println("list2 sorted", list2)

	totalDistance := 0
	for i := 0; i < len(list1); i++ {
		distance := int(math.Abs(float64(list1[i]) - float64(list2[i])))
		totalDistance += distance
	}
	fmt.Println("total distance part 1", totalDistance)

	list2Map := make(map[int]int, len(list2))
	for i := 0; i < len(list2); i++ {
		list2Map[list2[i]]++
	}
	totalSimilarityScore := 0
	for i := 0; i < len(list1); i++ {
		similarityScore := list1[i] * list2Map[list1[i]]
		totalSimilarityScore += similarityScore
	}
	fmt.Println("total similarity score part 2", totalSimilarityScore)
}
