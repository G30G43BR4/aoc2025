package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	rawLines := strings.Split(string(data), "\n")
	lines := make([]string, 0, len(rawLines))
	for _, l := range rawLines {
		l = strings.TrimSpace(l)
		if l != "" {
			lines = append(lines, l)
		}
	}
	return lines, nil
}

func getMaxJoltage1(input string) int {
	max := 0
	for i := 1; i < len(input); i++ {
		value, err := strconv.Atoi(string(input[0]) + string(input[i]))
		if err != nil {
			log.Fatal(err)
		}
		if value > max {
			max = value
		}
	}
	return max
}

func getMaxJoltage(input string, batteries int) string {
	if batteries > len(input) {
		fmt.Println("Selected more batteries than available.")
		return ""
	}
	max := ""
	checkString := input
	for i := batteries; i > 0; i-- {
		cutoff := len(checkString) - i + 1
		checkCut := checkString[:cutoff]
		maxValue := getMax(checkCut)
		max += maxValue
		_, after, _ := strings.Cut(checkString, maxValue)
		checkString = after
	}
	return max
}

func getMax(input string) string {
	max := 0
	for i := 0; i < len(input); i++ {
		value, _ := strconv.Atoi(string(input[i]))
		if value > max {
			max = value
		}
	}
	return strconv.Itoa(max)
}

func part01(lines []string) int {
	totalJoltage := 0
	for _, line := range lines {
		maxJoltage := 0
		for i := 0; i <= (len(line) - 2); i++ {
			check := string(line[i:])
			joltage := getMaxJoltage1(check)
			if joltage > maxJoltage {
				maxJoltage = joltage
			}
		}
		totalJoltage += maxJoltage
	}

	return totalJoltage
}

func part02(lines []string) int {
	totalJoltage := 0

	for _, line := range lines {
		joltage, err := strconv.Atoi(getMaxJoltage(line, 12))
		if err != nil {
			log.Fatal(err)
		}
		totalJoltage += joltage
	}

	return totalJoltage
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1:", part01(lines))
	fmt.Println("Part 2:", part02(lines))
}
