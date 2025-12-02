package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part01() {

	dial := 50
	pw := 0

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		rotation := strings.SplitAfterN(line, "", 2)
		direction := rotation[0]
		distance, err := strconv.Atoi(rotation[1])

		if err != nil {
			fmt.Println("Error converting string:", err)
			return
		}

		if direction == "L" {
			dial = (dial + 100 - (distance % 100)) % 100
		} else if direction == "R" {
			dial = (dial + distance) % 100
		} else {
			fmt.Print("Unsupported rotation sequence.")
		}
		if dial == 0 {
			pw++
		}
	}
	fmt.Println("Part1:")
	fmt.Println(pw)
}

func part02() {
	dial := 50
	zeroes := 0

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	lines := strings.Split(string(data), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		rotation := strings.SplitAfterN(line, "", 2)
		direction := rotation[0]
		distance, err := strconv.Atoi(rotation[1])

		if err != nil {
			fmt.Println("Error converting string:", err)
			return
		}

		if direction == "L" {
			zeroes = zeroes + (distance / 100)
			dial_old := dial
			dial = (dial - (distance % 100))
			if dial <= 0 && dial_old != 0 {
				zeroes++
			}
			if dial < 0 {
				dial = 100 + dial
			}
		} else if direction == "R" {
			zeroes = zeroes + ((dial + distance) / 100)
			dial = (dial + distance) % 100
		} else {
			fmt.Print("Unsupported rotation sequence.")
		}
	}
	fmt.Println("Part2:")
	fmt.Println(zeroes)
}

func main() {
	part01()
	part02()
}
