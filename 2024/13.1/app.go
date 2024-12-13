package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := OpenFile("/2024/13.1/input1.2.txt")
	machines := ExtractMachines(file)
	costs := CalculateCosts(machines)

	fmt.Println(costs)
}

func CalculateCosts(machines []Machine) int {
	costs := 0
	for i := 0; i < len(machines); i++ {
		cost := machines[i].CalculateCost(100)
		if cost != 0 {
			fmt.Println(cost)

		}
		costs += cost
	}

	return costs
}

func ExtractMachines(file *os.File) []Machine {
	var machines []Machine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := []int{}
		for _, v := range strings.Split(line, ",") {
			numberAsInt, _ := strconv.Atoi(v)
			numbers = append(numbers, numberAsInt)
		}
		machines = append(machines, Machine{numbers[0],
			numbers[1],
			numbers[2],
			numbers[3],
			numbers[4],
			numbers[5]})
	}

	return machines
}

func OpenFile(filename string) *os.File {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

type Machine struct {
	Ax int
	Ay int
	Bx int
	By int
	X  int
	Y  int
}

func (m Machine) CalculateCost(maxPresses int) int {

	// https://en.wikipedia.org/wiki/Cramer's_rule#Explicit_formulas_for_small_systems
	denominator := m.Ax*m.By - m.Ay*m.Bx

	if denominator == 0 {
		return 0
	}

	x := float64(m.By*m.X-m.Bx*m.Y) / float64(denominator)
	y := float64(m.Ax*m.Y-m.X*m.Ay) / float64(denominator)

	if x == math.Trunc(x) && y == math.Trunc(y) && 0 <= x && int(x) <= maxPresses && 0 <= int(y) && int(y) <= maxPresses {
		return 3*int(x) + int(y)
	}

	return 0
}
