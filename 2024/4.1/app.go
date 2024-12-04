package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file := OpenFile()
	defer file.Close()

	result := ExtractMatrixOfLetters(file)

	fmt.Println(NumberOfXmasInAWindow(result))
}

func NumberOfXmasInAWindow(result [][]string) int {
	xmasCounter := 0
	for i := 0; i < len(result[0]); i++ {
		for j := 0; j < len(result[0])-3; j++ {
			if IsVerticalXmasPresent(result, i, j) {
				xmasCounter++
			}
		}
	}

	for i := 0; i < len(result[0])-3; i++ {
		for j := 0; j < len(result[0]); j++ {
			if IsHorizontalXmasPresent(result, i, j) {
				xmasCounter++
			}
		}
	}

	for i := 0; i < len(result[0])-3; i++ {
		for j := 0; j < len(result[0])-3; j++ {
			if IsReverseDiagonalXmasPresent(result, i, j) {
				xmasCounter++
			}
		}
	}

	for i := 0; i < len(result[0])-3; i++ {
		for j := 0; j < len(result[0])-3; j++ {
			if IsDiagonalXmasPresent(result, i, j) {
				xmasCounter++
			}
		}
	}

	return xmasCounter
}

func IsVerticalXmasPresent(result [][]string, i int, j int) bool {
	return (result[i][j] == "X" &&
		result[i][j+1] == "M" &&
		result[i][j+2] == "A" &&
		result[i][j+3] == "S") ||
		result[i][j] == "S" &&
			result[i][j+1] == "A" &&
			result[i][j+2] == "M" &&
			result[i][j+3] == "X"
}

func IsHorizontalXmasPresent(result [][]string, i int, j int) bool {
	return result[i][j] == "X" &&
		result[i+1][j] == "M" &&
		result[i+2][j] == "A" &&
		result[i+3][j] == "S" ||
		result[i][j] == "S" &&
			result[i+1][j] == "A" &&
			result[i+2][j] == "M" &&
			result[i+3][j] == "X"
}

func IsDiagonalXmasPresent(result [][]string, i int, j int) bool {
	return result[i][j] == "X" &&
		result[i+1][j+1] == "M" &&
		result[i+2][j+2] == "A" &&
		result[i+3][j+3] == "S" ||
		result[i][j] == "S" &&
			result[i+1][j+1] == "A" &&
			result[i+2][j+2] == "M" &&
			result[i+3][j+3] == "X"
}

func IsReverseDiagonalXmasPresent(result [][]string, i int, j int) bool {
	return result[i][j+3] == "X" &&
		result[i+1][j+2] == "M" &&
		result[i+2][j+1] == "A" &&
		result[i+3][j] == "S" ||
		result[i][j+3] == "S" &&
			result[i+1][j+2] == "A" &&
			result[i+2][j+1] == "M" &&
			result[i+3][j] == "X"
}

func ExtractMatrixOfLetters(file *os.File) [][]string {
	scanner := bufio.NewScanner(file)
	var result [][]string
	for scanner.Scan() {
		letters := strings.Split(scanner.Text(), "")
		result = append(result, letters)
	}
	return result
}

func OpenFile() *os.File {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/2024/4.1/input2.txt")
	if err != nil {
		log.Fatal(err)
	}

	return file
}
