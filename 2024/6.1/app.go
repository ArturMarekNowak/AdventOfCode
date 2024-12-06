package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	guardsMapFile := OpenFile("/2024/6.1/input1.2.txt")
	guardsMap := ExtractLines(guardsMapFile)

	x, y := FindInitPosition(guardsMap)
	stepsCount := CountNumberOfSteps(guardsMap, x, y)

	fmt.Println(stepsCount)
}

const (
	right = iota
	left
	up
	down
)

func CountNumberOfSteps(guardsMap [][]string, x int, y int) int {

	direction := up
	cnt, i, j := 0, x, y

	for {
		if i+1 == len(guardsMap) || j+1 == len(guardsMap) {
			break
		}

		if direction == down && guardsMap[i+1][j] == "#" ||
			direction == up && guardsMap[i-1][j] == "#" ||
			direction == left && guardsMap[i][j-1] == "#" ||
			direction == right && guardsMap[i][j+1] == "#" {
			switch direction {
			case right:
				direction = down
				break
			case left:
				direction = up
				break
			case up:
				direction = right
				break
			case down:
				direction = left
				break
			default:
				panic("Nope")
			}
		}

		switch direction {
		case right:
			j++
			break
		case left:
			j--
			break
		case up:
			i--
			break
		case down:
			i++
			break
		default:
			panic("Nope")
		}

		if !(guardsMap[i][j] == "X") {
			cnt++
			guardsMap[i][j] = "X"
		}
	}

	return cnt
}

func FindInitPosition(guardsMap [][]string) (int, int) {
	for i := 0; i < len(guardsMap); i++ {
		for j := 0; j < len(guardsMap[i]); j++ {
			if guardsMap[i][j] == "^" {
				return i, j
			}
		}
	}

	panic("Nope")
}

func ExtractLines(file *os.File) [][]string {
	scanner := bufio.NewScanner(file)
	var result [][]string
	for scanner.Scan() {
		letters := strings.Split(scanner.Text(), "")
		result = append(result, letters)
	}
	return result
}

func OpenFile(filename string) *os.File {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
