package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	equationsFile := OpenFile("/2024/8.1/input1.2.txt")
	antennas, n := ExtractLines(equationsFile)
	antinodes := CalculateAntinodes(antennas, n)

	fmt.Println(antinodes)
	fmt.Println(len(antinodes))
}

func CalculateAntinodes(antennas map[uint8][][2]int, n int) [][2]int {
	antinodes := [][2]int{}
	for _, v := range antennas {
		for i := 0; i < len(v); i++ {
			for j := i + 1; j < len(v); j++ {
				firstAntenna := v[i]
				secondAntenna := v[j]

				diffX := secondAntenna[0] - firstAntenna[0]
				diffY := secondAntenna[1] - firstAntenna[1]

				firstAntinode := [2]int{firstAntenna[0] - diffX, firstAntenna[1] - diffY}
				secondAntinode := [2]int{secondAntenna[0] + diffX, secondAntenna[1] + diffY}

				if IsAntinodeInGrid(firstAntinode, n) && !slices.Contains(antinodes, firstAntinode) {
					antinodes = append(antinodes, firstAntinode)
				}

				if IsAntinodeInGrid(secondAntinode, n) && !slices.Contains(antinodes, secondAntinode) {
					antinodes = append(antinodes, secondAntinode)
				}
			}
		}
	}

	return antinodes
}

func IsAntinodeInGrid(antinode [2]int, n int) bool {
	return antinode[0] >= 0 && antinode[0] < n && antinode[1] >= 0 && antinode[1] < n
}

func ExtractLines(file *os.File) (map[uint8][][2]int, int) {
	scanner := bufio.NewScanner(file)
	result := make(map[uint8][][2]int)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		for i := 0; i < len(line); i++ {
			if line[i] != '.' {
				result[line[i]] = append(result[line[i]], [2]int{lineNumber, i})
			}
		}
		lineNumber++
	}

	return result, lineNumber
}

func OpenFile(filename string) *os.File {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
