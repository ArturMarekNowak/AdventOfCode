package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
)

func main() {
	file := OpenFile("/2024/12.1/input1.4.txt")
	grid := ExtractGrid(file)
	cost := CalculateFenceCost(grid)

	fmt.Println(cost)
}

type Position struct {
	i int
	j int
}

func CalculateFenceCost(grid [][]string) int {
	var result int
	var visited []Position
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if slices.Contains(visited, Position{i, j}) {
				continue
			}
			queue := []Position{{i, j}}
			area := 0
			perimeter := 0
			for len(queue) > 0 {
				r, c := queue[0].i, queue[0].j
				queue = queue[1:]
				if slices.Contains(visited, Position{r, c}) {
					continue
				}
				visited = append(visited, Position{r, c})
				area++
				neighbouringCoordinates := []Position{
					{r, c + 1},
					{r, c - 1},
					{r + 1, c},
					{r - 1, c}}
				for _, v := range neighbouringCoordinates {
					if 0 <= v.i && v.i < len(grid) && 0 <= v.j && v.j < len(grid[0]) && grid[v.i][v.j] == grid[r][c] {
						queue = append(queue, Position{v.i, v.j})
					} else {
						perimeter++
					}
				}
			}
			result += area * perimeter
		}
	}

	return result
}

func ExtractGrid(file *os.File) [][]string {
	var grid [][]string
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		var row []string
		for i := 0; i < len(line); i++ {
			row = append(row, string(line[i]))
		}
		grid = append(grid, row)
		lineNumber++
	}

	return grid
}

func OpenFile(filename string) *os.File {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
