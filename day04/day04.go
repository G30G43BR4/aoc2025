package main

import (
	"fmt"
	"log"
	"os"
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

func getNeighbours(x int, y int, lines []string) string {
	neighbours := ""
	if x == 0 {
		if y == 0 {
			neighbours += string(lines[y][x+1])
			neighbours += string(lines[y+1][x])
			neighbours += string(lines[y+1][x+1])
			return neighbours
		} else if y == len(lines)-1 {
			neighbours += string(lines[y][x+1])
			neighbours += string(lines[y-1][x])
			neighbours += string(lines[y-1][x+1])
			return neighbours
		} else {
			neighbours += string(lines[y-1][x])
			neighbours += string(lines[y-1][x+1])
			neighbours += string(lines[y][x+1])
			neighbours += string(lines[y+1][x])
			neighbours += string(lines[y+1][x+1])
			return neighbours
		}
	} else if x == len(lines[0])-1 {
		if y == 0 {
			neighbours += string(lines[y][x-1])
			neighbours += string(lines[y+1][x])
			neighbours += string(lines[y+1][x-1])
			return neighbours
		} else if y == len(lines)-1 {
			neighbours += string(lines[y][x-1])
			neighbours += string(lines[y-1][x])
			neighbours += string(lines[y-1][x-1])
			return neighbours
		} else {
			neighbours += string(lines[y-1][x])
			neighbours += string(lines[y-1][x-1])
			neighbours += string(lines[y][x-1])
			neighbours += string(lines[y+1][x])
			neighbours += string(lines[y+1][x-1])
			return neighbours
		}
	} else if y == 0 {
		if x > 0 && x < len(lines[0])-1 {
			neighbours += string(lines[y][x-1])
			neighbours += string(lines[y][x+1])
			neighbours += string(lines[y+1][x-1])
			neighbours += string(lines[y+1][x])
			neighbours += string(lines[y+1][x+1])
			return neighbours
		}
	} else if y == len(lines)-1 {
		if x > 0 && x < len(lines[0])-1 {
			neighbours += string(lines[y][x-1])
			neighbours += string(lines[y][x+1])
			neighbours += string(lines[y-1][x-1])
			neighbours += string(lines[y-1][x])
			neighbours += string(lines[y-1][x+1])
			return neighbours
		}
	}
	neighbours += string(lines[y-1][x-1])
	neighbours += string(lines[y-1][x])
	neighbours += string(lines[y-1][x+1])
	neighbours += string(lines[y][x-1])
	neighbours += string(lines[y][x+1])
	neighbours += string(lines[y+1][x-1])
	neighbours += string(lines[y+1][x])
	neighbours += string(lines[y+1][x+1])
	return neighbours
}

func checkNeighbours(neighbours string) int {
	paperRolls := 0
	for i := 0; i < len(neighbours); i++ {
		if string(neighbours[i]) == "@" {
			paperRolls++
		}
	}
	return paperRolls
}

func clearRolls(lines []string) (int, []string) {
	count := 0
	for i := range lines {
		b := []byte(lines[i])
		for j := 0; j < len(b); j++ {
			if b[j] == '@' {
				neighbours := getNeighbours(j, i, lines)
				paperRolls := checkNeighbours(neighbours)
				if paperRolls < 4 {
					b[j] = '.'
					count++
				}
			}
		}
		lines[i] = string(b)
	}
	return count, lines
}

func part01(lines []string) int {
	accessibleRolls := 0
	for i, line := range lines {
		for j := 0; j < len(line); j++ {
			if string(line[j]) == "@" {
				neighbours := getNeighbours(j, i, lines)
				paperRolls := checkNeighbours(neighbours)
				if paperRolls < 4 {
					accessibleRolls++
				}
			}
		}
	}
	return accessibleRolls
}

func part02(lines []string) int {
	clearedRolls := 0
	for {
		count, lines_new := clearRolls(lines)
		if count == 0 {
			break
		}
		clearedRolls += count
		lines = lines_new
	}
	return clearedRolls
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Part 1:", part01(lines))
	fmt.Println("Part 2:", part02(lines))
}
