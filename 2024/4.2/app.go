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
	for i := 0; i < len(result[0])-2; i++ {
		for j := 0; j < len(result[0])-2; j++ {
			if IsVerticalXmasPresent(result, i, j) {
				xmasCounter++
			}
		}
	}

	return xmasCounter
}

func IsVerticalXmasPresent(result [][]string, i int, j int) bool {
	return result[i][j] == "M" &&
		result[i+2][j] == "M" &&
		result[i+1][j+1] == "A" &&
		result[i][j+2] == "S" &&
		result[i+2][j+2] == "S" ||
		result[i][j] == "S" &&
			result[i+2][j] == "M" &&
			result[i+1][j+1] == "A" &&
			result[i][j+2] == "S" &&
			result[i+2][j+2] == "M" ||
		result[i][j] == "S" &&
			result[i+2][j] == "S" &&
			result[i+1][j+1] == "A" &&
			result[i][j+2] == "M" &&
			result[i+2][j+2] == "M" ||
		result[i][j] == "M" &&
			result[i+2][j] == "S" &&
			result[i+1][j+1] == "A" &&
			result[i][j+2] == "M" &&
			result[i+2][j+2] == "S"
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
	file, err := os.Open(pwd + "/2024/4.2/input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	return file
}
