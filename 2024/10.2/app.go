package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	equationsFile := OpenFile("/2024/10.1/input1.2.txt")
	grid := ExtractGrid(equationsFile)
	possibleTrailheads := CalculatePossibleTrailheads(grid)
	cnt := CountSumOfTrailheadsScores(possibleTrailheads, grid)

	fmt.Println(cnt)
}

func CountSumOfTrailheadsScores(possibleTrailheads []Position, grid [][]int) int {
	cnt := 0
	for i := 0; i < len(possibleTrailheads); i++ {
		cnt += CalculateTrailheadsScores(grid, possibleTrailheads[i].i, possibleTrailheads[i].j)
	}
	return cnt
}

func CalculateTrailheadsScores(grid [][]int, n int, m int) int {
	var queue []Position
	queue = append(queue, Position{n, m})
	trails := 0
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		neighbouringCoordinates := []Position{
			{elem.i, elem.j + 1},
			{elem.i, elem.j - 1},
			{elem.i + 1, elem.j},
			{elem.i - 1, elem.j}}
		for _, neighbour := range neighbouringCoordinates {
			if neighbour.i < 0 || neighbour.j < 0 || neighbour.i >= len(grid) || neighbour.j >= len(grid[0]) {
				continue
			}
			if grid[neighbour.i][neighbour.j] != grid[elem.i][elem.j]+1 {
				continue
			}
			if grid[neighbour.i][neighbour.j] == 9 {
				trails++
			} else {
				queue = append(queue, neighbour)
			}
		}
	}

	return trails
}

func CalculatePossibleTrailheads(grid [][]int) []Position {
	var possibleTrailheads []Position
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 {
				possibleTrailheads = append(possibleTrailheads, Position{i, j})
			}
		}
	}

	return possibleTrailheads
}

type Position struct {
	i int
	j int
}

func ExtractGrid(file *os.File) [][]int {
	var grid [][]int
	scanner := bufio.NewScanner(file)
	lineNumber := 0
	for scanner.Scan() {
		line := scanner.Text()
		var row []int
		for i := 0; i < len(line); i++ {
			number, _ := strconv.Atoi(string(line[i]))
			row = append(row, number)
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
