package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	stones := []string{"92", "0", "286041", "8034", "34394", "795", "8", "2051489"}
	//stones := []string{"125", "17"}
	stones = Blink25Times(stones)

	fmt.Println(len(stones))
}

func Blink25Times(stones []string) []string {
	for i := 0; i < 25; i++ {
		stones = Blink(stones)
		fmt.Println(i)
	}

	return stones
}

func Blink(stones []string) []string {
	var ret []string
	for i := 0; i < len(stones); i++ {
		if stones[i] == "0" {
			ret = append(ret, "1")
		} else if len(stones[i])%2 == 0 {
			ret = append(ret, TrimTrailingZeroes(stones[i][:len(stones[i])/2]))
			ret = append(ret, TrimTrailingZeroes(stones[i][len(stones[i])/2:]))
		} else {
			stoneNumber, _ := strconv.Atoi(stones[i])
			ret = append(ret, strconv.Itoa(stoneNumber*2024))
		}
	}

	return ret
}

func TrimTrailingZeroes(s string) string {
	trimmedZeroes := strings.TrimLeft(s, "0")

	if trimmedZeroes == "" {
		return "0"
	}

	return trimmedZeroes
}
