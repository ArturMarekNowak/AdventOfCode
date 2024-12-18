package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	const N = 71
	file := OpenFile("/2024/18.1/input1.2.txt")
	coords := ExtractCoords(file)
	grid := GenerateGrid(coords, N)
	pathCost := CalculatePathCost(grid, N)

	fmt.Println(pathCost)
}

type Path struct {
	Cost     int
	Position Position
}

type Position struct {
	x int
	y int
}

func CalculatePathCost(grid [][]string, n int) int {
	var queue []Path
	var seen []Position
	queue = append(queue, Path{0, Position{0, 0}})
	foo := 0
	for len(queue) > 0 {
		elem := queue[0]
		queue = queue[1:]
		foo++
		if elem.Position.x == n-1 && elem.Position.y == n-1 {
			return elem.Cost
		}
		if slices.Contains(seen, elem.Position) {
			continue
		}
		neighbouringCoordinates := []Position{
			{-1, 0},
			{0, 1},
			{1, 0},
			{0, -1}}
		seen = append(seen, elem.Position)
		for _, neighbour := range neighbouringCoordinates {
			row := elem.Position.x + neighbour.x
			col := elem.Position.y + neighbour.y
			if 0 <= row && row < n && 0 <= col && col < n && grid[row][col] != "#" {
				queue = append(queue, Path{elem.Cost + 1, Position{row, col}})
			}
		}
	}

	fmt.Println(foo)
	return -1
}

func GenerateGrid(coords [][2]int, n int) [][]string {
	grid := make([][]string, n)
	for i := 0; i < n; i++ {
		grid[i] = make([]string, n)
		for j := 0; j < n; j++ {
			grid[i][j] = "."
		}
	}

	fmt.Println(len(coords))

	//for i := 0; i < len(coords); i++ {
	for i := 0; i < 1024; i++ {
		grid[coords[i][0]][coords[i][1]] = "#"
		//grid[coords[i][1]][coords[i][0]] = "#"
	}

	return grid
}

func ExtractCoords(file *os.File) [][2]int {
	var coords [][2]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbersAsStrings := strings.Split(line, ",")
		x, _ := strconv.Atoi(numbersAsStrings[0])
		y, _ := strconv.Atoi(numbersAsStrings[1])
		coords = append(coords, [2]int{x, y})
	}

	return coords
}

func OpenFile(filename string) *os.File {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
