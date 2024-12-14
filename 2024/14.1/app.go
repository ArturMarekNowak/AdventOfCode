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
	file := OpenFile("/2024/14.1/input1.2.txt")
	const length = 101
	const height = 103
	const seconds = 100
	robots := ExtractRobots(file)
	robots = CalculateRobotsPositions(robots, length, height, seconds)
	numberOfRobotsNotInTheMidRowsAndCols := CalculateNumberOfRobotsNotInTheMidRowsAndCols(robots, length, height)

	fmt.Println(numberOfRobotsNotInTheMidRowsAndCols)

	PrintPossibleAnswers(robots, length, height, 10000)
}

// I am not proud of this function
func PrintPossibleAnswers(robots []Robot, length int, height int, seconds int) {
	for i := 100; i < seconds; i++ {
		tmpRobots := make([]Robot, len(robots))
		for j := 0; j < len(robots); j++ {
			tmpRobots = append(tmpRobots, robots[j].CalculateRobotPosition(i, length, height))
		}

		numberOfRobotsInMid := CalculateNumberOfRobotsInMid(tmpRobots, length, height)
		if numberOfRobotsInMid > 200 {
			fmt.Println(i + 100)
		}
	}
}

func CalculateNumberOfRobotsInMid(robots []Robot, length int, height int) int {
	cnt := 0
	coeff := 0.3
	x1 := int(coeff * float64(length))
	y1 := int(coeff * float64(height))
	x2 := int(2 * coeff * float64(height))
	y2 := int(2 * coeff * float64(height))
	for _, robot := range robots {
		if robot.X > x1 && robot.Y > y1 && robot.X < x2 && robot.Y < y2 {
			cnt++
		}
	}

	return cnt
}

func CalculateNumberOfRobotsNotInTheMidRowsAndCols(robots []Robot, length int, height int) int {
	cntQ1 := 0
	cntQ2 := 0
	cntQ3 := 0
	cntQ4 := 0
	midRow := length / 2
	midCol := height / 2
	for _, robot := range robots {
		if robot.X < midRow && robot.Y < midCol {
			cntQ1++
		}

		if robot.X < midRow && robot.Y > midCol {
			cntQ2++
		}

		if robot.X > midRow && robot.Y < midCol {
			cntQ3++
		}

		if robot.X > midRow && robot.Y > midCol {
			cntQ4++
		}
	}

	return cntQ1 * cntQ2 * cntQ3 * cntQ4
}

func CalculateRobotsPositions(robots []Robot, length int, height int, seconds int) []Robot {
	var retRobots []Robot
	for i := 0; i < len(robots); i++ {
		retRobots = append(retRobots, robots[i].CalculateRobotPosition(seconds, length, height))
	}

	return retRobots
}

type Robot struct {
	X  int
	Y  int
	Vx int
	Vy int
}

func (robot Robot) CalculateRobotPosition(seconds int, length int, height int) Robot {
	robot.X = (robot.X + (seconds * robot.Vx)) % length
	robot.Y = (robot.Y + (seconds * robot.Vy)) % height

	if robot.X < 0 {
		robot.X = length + robot.X
	}

	if robot.Y < 0 {
		robot.Y = height + robot.Y
	}

	return robot
}

func ExtractRobots(file *os.File) []Robot {
	var robots []Robot
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		numbers := []int{}
		for _, v := range strings.Split(line, ",") {
			numberAsInt, _ := strconv.Atoi(v)
			numbers = append(numbers, numberAsInt)
		}
		robots = append(robots, Robot{numbers[0],
			numbers[1],
			numbers[2],
			numbers[3]})
	}

	return robots
}

func OpenFile(filename string) *os.File {
	pwd, _ := os.Getwd()
	file, err := os.Open(pwd + filename)
	if err != nil {
		log.Fatal(err)
	}

	return file
}
