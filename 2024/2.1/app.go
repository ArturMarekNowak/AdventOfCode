package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + "/2024/2.1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	safeLevels := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		numbers := strings.Fields(scanner.Text())
		if (AreAllNumbersAscending(numbers) ||
			AreAllNumbersDescending(numbers)) &&
			IsDifferenceOfAllNumbersSmallerThanThreeAndBiggerThanOne(numbers) {
			fmt.Println(numbers)
			safeLevels++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(safeLevels)
}

func AreAllNumbersAscending(numbers []string) bool {
	for i := 0; i < len(numbers)-1; i++ {
		a, _ := strconv.Atoi(numbers[i])
		b, _ := strconv.Atoi(numbers[i+1])
		if a > b {
			return false
		}
	}

	return true
}

func AreAllNumbersDescending(numbers []string) bool {
	for i := 0; i < len(numbers)-1; i++ {
		a, _ := strconv.Atoi(numbers[i])
		b, _ := strconv.Atoi(numbers[i+1])
		if a < b {
			return false
		}
	}

	return true
}

func IsDifferenceOfAllNumbersSmallerThanThreeAndBiggerThanOne(numbers []string) bool {
	for i := 0; i < len(numbers)-1; i++ {
		a, _ := strconv.Atoi(numbers[i])
		b, _ := strconv.Atoi(numbers[i+1])
		absDiff := Abs(a, b)
		if absDiff < 1 || absDiff > 3 {
			return false
		}
	}

	return true
}

func Abs(a int, b int) int {
	if a-b < 0 {
		return b - a
	}
	return a - b
}
