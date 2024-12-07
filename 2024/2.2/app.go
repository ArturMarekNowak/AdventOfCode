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
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/2024/2.2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	safeLevels := 0
	scanner := bufio.NewScanner(file)
	var result [][]int
	var orig [][]int
	for scanner.Scan() {
		numbers := ConvertToSliceOfNumbers(strings.Fields(scanner.Text()))
		numbersDiff := CalculateSliceOfDiffs(numbers)
		orig = append(orig, numbers)
		result = append(result, numbersDiff)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(result); i++ {
		countOfWrongDiffNumbers, idx := CountOfWrongDiffNumbers(result[i])
		if countOfWrongDiffNumbers == 0 {
			safeLevels++
		}

		if countOfWrongDiffNumbers != len(idx) {
			fmt.Println(i)
		}

		if countOfWrongDiffNumbers == 1 && len(idx) == 1 {
			tmp := slices.Delete(orig[i], idx[0], idx[0]+1)
			countOfWrongDiffNumbersAgain, _ := CountOfWrongDiffNumbers(CalculateSliceOfDiffs(tmp))
			if countOfWrongDiffNumbersAgain == 0 {
				//fmt.Println(orig[i])
				//fmt.Println(tmp)
				//fmt.Println(i)
				safeLevels++
			}
		}
	}

	fmt.Println(safeLevels)
}

func CountOfWrongDiffNumbers(numbers []int) (int, []int) {
	positiveNumbers := 0
	negativeNumbers := 0
	wrongNumbers := 0
	var positiveNumbersIdx []int
	var negativeNumbersIdx []int
	var wrongNumbersIdx []int

	for i := 0; i < len(numbers); i++ {
		if numbers[i] <= 3 && numbers[i] >= 1 {
			positiveNumbers++
			positiveNumbersIdx = append(positiveNumbersIdx, i)
		} else if numbers[i] >= -3 && numbers[i] <= -1 {
			negativeNumbers++
			negativeNumbersIdx = append(wrongNumbersIdx, i)
		} else {
			wrongNumbers++
			wrongNumbersIdx = append(wrongNumbersIdx, i)
		}
	}

	if positiveNumbers > negativeNumbers {
		return negativeNumbers + wrongNumbers, append(wrongNumbersIdx, negativeNumbersIdx...)
	}

	return positiveNumbers + wrongNumbers, append(wrongNumbersIdx, positiveNumbersIdx...)
}

func CalculateSliceOfDiffs(numbers []int) []int {
	var numbersDiff []int
	for i := 0; i < len(numbers)-1; i++ {
		numbersDiff = append(numbersDiff, numbers[i]-numbers[i+1])
	}

	return numbersDiff
}

func ConvertToSliceOfNumbers(strings []string) []int {
	result := make([]int, len(strings))

	for i, str := range strings {
		result[i], _ = strconv.Atoi(str)
	}

	return result
}

func Abs(a int, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
