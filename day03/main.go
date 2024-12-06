package main

import (
	"fmt"
	"log"
	"os"
	"regexp"

	util "github.com/yattoni/advent-of-code-2024"
)

func main() {
	bytes, err := os.ReadFile(util.INPUT)
	if err != nil {
		log.Fatal("Error reading file", err)
	}
	input := string(bytes)
	fmt.Println(input)
	regex := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := regex.FindAllString(input, -1)
	fmt.Println(matches)

	sum := 0
	digitRegex := regexp.MustCompile(`\d+`)
	for i, m := range matches {
		digitStrs := digitRegex.FindAllString(m, -1)
		if len(digitStrs) != 2 {
			log.Fatal("not enough digits", i, m, digitStrs)
		}
		sum += util.MustAtoi(digitStrs[0]) * util.MustAtoi(digitStrs[1])
	}
	fmt.Println("part1 sum", sum)

	// bytes2, err := os.ReadFile(util.PROMPT_INPUT + "-2")
	bytes2, err := os.ReadFile(util.INPUT)
	if err != nil {
		log.Fatal("Error reading file", err)
	}
	input2 := string(bytes2)
	fmt.Println(input2)
	regex2 := regexp.MustCompile(`(mul\(\d+,\d+\)|do\(\)|don't\(\))`)
	matches2 := regex2.FindAllString(input2, -1)
	fmt.Println(matches2)

	sum2 := 0
	currentlyEnabled := true
	for i, m := range matches2 {
		if m == "do()" {
			currentlyEnabled = true
		} else if m == "don't()" {
			currentlyEnabled = false
		} else if currentlyEnabled {
			digitStrs := digitRegex.FindAllString(m, -1)
			if len(digitStrs) != 2 {
				log.Fatal("not enough digits", i, m, digitStrs)
			}
			sum2 += util.MustAtoi(digitStrs[0]) * util.MustAtoi(digitStrs[1])
		}
	}
	fmt.Println("part2 sum", sum2)
}
