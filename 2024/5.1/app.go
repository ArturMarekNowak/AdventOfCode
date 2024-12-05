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

	updatesFile := OpenFile("/2024/5.1/input2.2.txt")
	pagesFile := OpenFile("/2024/5.1/input2.1.txt")

	updates := ExtractUpdates(updatesFile)
	pages := ExtractPages(pagesFile)

	correctPagesUpdates := CalculateListOfCorrectPagesUpdates(updates)
	foundPagesUpdates := SearchForCorrectPagesUpdates(updates, pages, correctPagesUpdates)
	sumOfMiddleNumbers := SumOfMiddleNumbers(foundPagesUpdates)

	fmt.Println(updates)
	fmt.Println(pages)
	fmt.Println(correctPagesUpdates)
	fmt.Println(foundPagesUpdates)
	fmt.Println(sumOfMiddleNumbers)
}

func SumOfMiddleNumbers(updates [][]string) interface{} {
	sum := 0
	for i := 0; i < len(updates); i++ {

		if len(updates[i])%2 == 0 {
			panic("Nope")
		}

		middleIdx := len(updates[i]) / 2
		num, _ := strconv.Atoi(updates[i][middleIdx])
		sum += num
	}

	return sum
}

func SearchForCorrectPagesUpdates(updates [][]string, pages []string, correctPagesUpdates [][]string) [][]string {

	var ret [][]string
	for i := 0; i < len(updates); i++ {
		if IsPagesUpdateCorrect(correctPagesUpdates[i], pages) {
			ret = append(ret, updates[i])
		}
	}

	return ret
}

func IsPagesUpdateCorrect(correctPagesUpdates []string, pages []string) bool {
	for i := 0; i < len(correctPagesUpdates); i++ {
		if !slices.Contains(pages, correctPagesUpdates[i]) {
			return false
		}
	}

	return true
}

func CalculateListOfCorrectPagesUpdates(updates [][]string) [][]string {
	var updatesRules [][]string
	for i := 0; i < len(updates); i++ {
		var tmp []string
		for j := 0; j < len(updates[i]); j++ {
			for k := j + 1; k < len(updates[i]); k++ {
				rule := fmt.Sprintf("%s|%s", updates[i][j], updates[i][k])
				tmp = append(tmp, rule)
			}
		}
		updatesRules = append(updatesRules, tmp)
	}
	return updatesRules
}

func ExtractPages(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var result []string
	for scanner.Scan() {
		letters := scanner.Text()
		result = append(result, letters)
	}
	return result
}

func ExtractUpdates(file *os.File) [][]string {
	scanner := bufio.NewScanner(file)
	var result [][]string
	for scanner.Scan() {
		letters := strings.Split(scanner.Text(), ",")
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
