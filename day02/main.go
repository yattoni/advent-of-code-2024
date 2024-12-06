package main

import (
	"fmt"
	"math"

	util "github.com/yattoni/advent-of-code-2024"
)

func main() {
	reports := util.GetNumbersFromFileSeparatedBySpaces(util.INPUT)

	fmt.Println("len(reports)=", len(reports))
	fmt.Println(reports[0:5])

	allIncreasing := func(report []int) bool {
		prev := int(math.Inf(-1))
		for _, val := range report {
			if val <= prev {
				return false
			}
			prev = val
		}
		return true
	}

	allDecreasing := func(report []int) bool {
		prev := int(math.Inf(1))
		for _, val := range report {
			if val >= prev {
				return false
			}
			prev = val
		}
		return true
	}

	getDeltas := func(report []int) []int {
		deltas := make([]int, len(report)-1)
		for i := 0; i < len(deltas); i++ {
			deltas[i] = int(math.Abs(float64(report[i+1]) - float64(report[i])))
		}
		return deltas
	}

	allWithinRange := func(xs []int, lower, upper int) bool {
		for _, x := range xs {
			if x < lower || x > upper {
				fmt.Println(x)
				return false
			}
		}
		return true
	}

	isSafe := func(report []int) bool {
		if !allIncreasing(report) && !allDecreasing(report) {
			return false
		}
		deltas := getDeltas(report)
		fmt.Println(deltas)
		return allWithinRange(deltas, 1, 3)
	}

	countOfSafe := 0
	for i, r := range reports {
		fmt.Println("evaluating report", i)
		isCurrSafe := isSafe(r)
		fmt.Println("report", i, "isSafe", isCurrSafe)
		if isCurrSafe {
			countOfSafe++
		}
	}

	fmt.Println("part1", countOfSafe)

	countOfSafe2 := 0
	for i, r := range reports {
		fmt.Println("evaluating report", i, r)
		possibilities := make([][]int, len(r)+1)
		possibilities[0] = r
		for j := 0; j < len(r); j++ {
			rClone := make([]int, len(r))
			copy(rClone, r)
			possibilities[j+1] = append(rClone[:j], rClone[j+1:]...)
		}
		fmt.Println(possibilities)
		isCurrSafe := false
		for _, p := range possibilities {
			isCurrSafe = isCurrSafe || isSafe(p)
		}
		fmt.Println("report", i, "isSafe", isCurrSafe)
		if isCurrSafe {
			countOfSafe2++
		}
	}
	fmt.Println("part2", countOfSafe2)
}
