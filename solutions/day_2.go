package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func solution2_1() {
	input := getInput(2)
	reports := make([][]int, len(input))
	safeReports := 0

	for idx, line := range input {
		pieces := strings.Split(line, " ")

		for _, piece := range pieces {
			reports[idx] = append(reports[idx], func() int {
				x, _ := strconv.Atoi(piece)
				return x
			}())
		}
	}

	for _, report := range reports {
		var increasing bool
		if report[0] < report[1] {
			increasing = true
		}

		var safe bool
		for i, v := range report {
			if i == 0 {
				safe = true
				continue
			}

			// incr/decr check
			if increasing {
				// check if curr < prev
				if v <= report[i-1] {
					safe = false
					break
				}
			}

			if !increasing {
				// check if curr > prev
				if v >= report[i-1] {
					safe = false
					break
				}
			}

			if -3 > v-report[i-1] || v-report[i-1] > 3 {
				safe = false
				break
			}
		}

		if safe {
			safeReports++
		}

	}

	fmt.Printf("D2P1 solution: %v\n", safeReports)
}

func solution2_2() {
	input := getInput(2)
	reports := make([][]int, len(input))
	safeReports := 0

	for idx, line := range input {
		pieces := strings.Split(line, " ")

		for _, piece := range pieces {
			reports[idx] = append(reports[idx], func() int {
				x, _ := strconv.Atoi(piece)
				return x
			}())
		}
	}

	for _, report := range reports {
		var increasing bool
		if report[0] < report[1] {
			increasing = true
		}

		var safe bool
		for i, v := range report {
			if i == 0 {
				safe = true
				continue
			}

			// incr/decr check
			if increasing {
				// check if curr < prev
				if v <= report[i-1] {
					safe = false
					break
				}
			}

			if !increasing {
				// check if curr > prev
				if v >= report[i-1] {
					safe = false
					break
				}
			}

			if -3 > v-report[i-1] || v-report[i-1] > 3 {
				safe = false
				break
			}
		}

		if safe {
			safeReports++
		}

	}

	fmt.Printf("D2P1 solution: %v\n", safeReports)
}

func init() {
	solution2_1()
	// TODO: finish pt2
}
