package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func getInput(day int) []string {
	file, err := os.Open(fmt.Sprintf("./inputs/day_%d.txt", day))
	if err != nil {
		log.Fatal("Error opening input file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error reading file:", err)
	}

	return lines
}
