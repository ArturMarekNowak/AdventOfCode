package main

import (
	"fmt"
	"strconv"
)

func main() {
	stones := []int{92, 0, 286041, 8034, 34394, 795, 8, 2051489}
	//stones := []int{125, 17}
	//foo := BlinkNTimes(stones, 25)
	bar := BlinkNTimes(stones, 75)

	//fmt.Println(Sum(foo))
	fmt.Println(Sum(bar))
}

func Sum(numbers []int) int {
	var cnt int
	for _, n := range numbers {
		cnt += n
	}
	return cnt
}

func BlinkNTimes(stones []int, n int) []int {
	var cnt []int
	var lookup = make(map[[2]int]int)
	for i := 0; i < len(stones); i++ {
		cnt = append(cnt, Blink(stones[i], n, lookup))
		fmt.Println(i)
	}

	return cnt
}

func Blink(number int, n int, lookup map[[2]int]int) int {

	_, ok := lookup[[2]int{number, n}]
	if ok {
		return lookup[[2]int{number, n}]
	}

	ret := 0
	if n == 0 {
		ret = 1
	} else if number == 0 {
		ret = Blink(1, n-1, lookup)
	} else if len(strconv.Itoa(number))%2 == 0 {
		numberAsString := strconv.Itoa(number)
		leftNumber, _ := strconv.Atoi(numberAsString[:len(numberAsString)/2])
		rightNumber, _ := strconv.Atoi(numberAsString[len(numberAsString)/2:])
		ret = Blink(leftNumber, n-1, lookup) + Blink(rightNumber, n-1, lookup)
	} else {
		ret = Blink(number*2024, n-1, lookup)
	}

	lookup[[2]int{number, n}] = ret
	return ret
}
