package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	equationsFile := OpenFile("/2024/7.1/input1.2.txt")
	equations := ExtractLines(equationsFile)

	t1 := time.Now()
	correctResults := CalculateCorrectEquations(equations)
	sum := Sum(correctResults)
	t2 := time.Now()

	diff := t2.Sub(t1)
	fmt.Println(diff)
	fmt.Println(sum)
}

func Sum(slice []int) int {
	cnt := 0
	for _, i := range slice {
		cnt += i
	}
	return cnt
}

func CalculateCorrectEquations(equations []Equation) []int {
	var results []int
	for i := 0; i < len(equations); i++ {
		if IsEquationPossible(equations[i]) {
			results = append(results, equations[i].result)
		}
	}

	return results
}

func IsEquationPossible(equation Equation) bool {
	possibleSignsCombinations := CalculatePossibleSignsCombinations(len(equation.products) - 1)

	for i := 0; i < len(possibleSignsCombinations); i++ {
		result := equation.products[0]
		for j := 0; j < len(possibleSignsCombinations[i]); j++ {

			if possibleSignsCombinations[i][j] == '+' {
				result += equation.products[j+1]
			} else if possibleSignsCombinations[i][j] == '*' {
				result *= equation.products[j+1]
			} else if possibleSignsCombinations[i][j] == 'd' {
				first := strconv.Itoa(result)
				second := strconv.Itoa(equation.products[j+1])
				concat, _ := strconv.Atoi(first + second)
				result = concat
			} else {
				panic("Nope")
			}
		}

		if equation.result == result {
			return true
		}
	}

	return false
}

func CalculatePossibleSignsCombinations(k int) []string {
	signs := []string{"+", "*", "d"}
	result := []string{}

	CalculatePossibleSignsCombinationsRecursively(signs, "", len(signs), k, &result)

	return result
}

func CalculatePossibleSignsCombinationsRecursively(signs []string, prefix string, n int, k int, result *[]string) {
	if k == 0 {
		*result = append(*result, prefix)
		return
	}

	for j := 0; j < n; j++ {
		newPrefix := prefix + signs[j]
		CalculatePossibleSignsCombinationsRecursively(signs, newPrefix, n, k-1, result)
	}
}

type Equation struct {
	result   int
	products []int
}

func ExtractLines(file *os.File) []Equation {
	scanner := bufio.NewScanner(file)
	var result []Equation
	for scanner.Scan() {
		sanitizedLines := strings.Replace(scanner.Text(), ":", "", -1)
		numbersAsStrings := strings.Split(sanitizedLines, " ")
		numbersAsInts := ConvertStringsArrayToIntsArray(numbersAsStrings)

		result = append(result, Equation{result: numbersAsInts[0], products: numbersAsInts[1:]})
	}

	return result
}

func ConvertStringsArrayToIntsArray(stringsArray []string) []int {
	var result []int
	for _, str := range stringsArray {
		number, _ := strconv.Atoi(str)
		result = append(result, number)
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
