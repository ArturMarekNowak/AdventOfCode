package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file := OpenFile("/2024/22.1/input1.2.txt")
	initNumbers := ExtractInitNumbers(file)
	secretNumbers := CalculateSecretNumbers(initNumbers)
	fmt.Println(Sum(secretNumbers))
}

func Sum(numbers []int) any {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func CalculateSecretNumbers(numbers []int) []int {
	secretNumbers := make([]int, len(numbers))
	for i := 0; i < len(numbers); i++ {
		secretNumbers[i] = CalculateSecretNumber(numbers[i])
	}
	return secretNumbers
}

func CalculateSecretNumber(n int) int {
	for i := 0; i < 2000; i++ {
		n = Prune(Mix(n, 64*n))
		n = Prune(Mix(n, n/32))
		n = Prune(Mix(n, n*2048))
	}
	return n
}

func Mix(n int, m int) int {
	return n ^ m
}

func Prune(n int) int {
	return n % 16777216
}

func ExtractInitNumbers(file *os.File) []int {
	var initNumbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, _ := strconv.Atoi(line)
		initNumbers = append(initNumbers, number)
	}

	return initNumbers
}

func OpenFile(filename string) *os.File {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
