package solutions

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func solution1_1() {
	input := getInput(1)
	left := []int{}
	right := []int{}
	distance := 0

	for _, line := range input {
		pieces := strings.Split(line, "   ")
		left = append(left, func() int {
			x, _ := strconv.Atoi(pieces[0])
			return x
		}())

		right = append(right, func() int {
			x, _ := strconv.Atoi(pieces[1])
			return x
		}())
	}

	sort.Ints(left)
	sort.Ints(right)

	for idx, val := range left {
		if x := val - right[idx]; x < 0 {
			distance += (x * -1)
		} else {
			distance += x
		}
	}

	fmt.Printf("D1P1 solution: %v\n", distance)
}

func solution1_2() {
	input := getInput(1)
	left := []int{}
	right := []int{}

	for _, line := range input {
		pieces := strings.Split(line, "   ")
		left = append(left, func() int {
			x, _ := strconv.Atoi(pieces[0])
			return x
		}())

		right = append(right, func() int {
			x, _ := strconv.Atoi(pieces[1])
			return x
		}())
	}

	countMap := make(map[int]int)

	for _, val := range right {
		if _, ok := countMap[val]; !ok {
			countMap[val] = 1
		} else {
			countMap[val]++
		}
	}

	simScore := 0
	for _, val := range left {
		if _, ok := countMap[val]; !ok {
			continue
		}

		simScore += val * countMap[val]
	}

	fmt.Printf("D1P2 solution: %v\n", simScore)
}

func init() {
	solution1_1()
	solution1_2()
}
