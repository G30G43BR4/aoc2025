package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func split_string(input string, segment_size int) []string {

	var parts []string

	if len(input)%segment_size != 0 {
		fmt.Println("Unsupported string-segment combination.")
		return parts
	}
	number_parts := len(input) / segment_size

	for i := 0; i < number_parts; i++ {
		start := (i * segment_size)
		end := start + segment_size

		parts = append(parts, input[start:end])
	}
	return parts
}

func check_parts(input []string) int {
	prev_part := ""
	count := 0

	for i, part := range input {
		if i == 0 {
			prev_part = part
			continue
		}
		if part != prev_part {
			return 0
		}
		count++
		prev_part = part
	}
	return count
}

func part01() {

	count := 0

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	segments := strings.Split(string(data), ",")

	for _, segment := range segments {
		split_segment := strings.Split(segment, "-")

		start, _ := strconv.Atoi(split_segment[0])
		end, _ := strconv.Atoi(split_segment[1])

		for i := start; i <= end; i++ {

			as_string := strconv.Itoa(i)

			if (len(as_string) % 2) != 0 {
				continue
			}

			split := len(as_string) / 2

			part1 := as_string[:split]
			part2 := as_string[split:]

			if part1 == part2 {
				count += i
			}
		}
	}
	fmt.Println("Part1:")
	fmt.Println(count)
}

func part02() {

	count := 0

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	segments := strings.Split(string(data), ",")

	for _, segment := range segments {
		split_segment := strings.Split(segment, "-")

		start, _ := strconv.Atoi(split_segment[0])
		end, _ := strconv.Atoi(split_segment[1])

		for i := start; i <= end; i++ {

			as_string := strconv.Itoa(i)

			for j := 1; j <= len(as_string); j++ {
				if len(as_string)%j != 0 {
					continue
				}
				parts := split_string(as_string, j)
				check := check_parts(parts)
				if check != 0 {
					count += i
					break
				}
			}
		}
	}
	fmt.Println("Part2:")
	fmt.Println(count)

}

func main() {
	part01()
	part02()
}
