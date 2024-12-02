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
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var numbers []string
	//true = going up, false = going down
	var posneg bool
	var difference int
	var safeLineCount int

	scanner := bufio.NewScanner(strings.NewReader(string(content)))

outerLoop:
	for scanner.Scan() {
		// The levels are either all increasing or all decreasing.
		// Any two adjacent levels differ by at least one and at most three.
		numbers = strings.Fields(scanner.Text())

		if stringToInt(numbers[0]) < stringToInt(numbers[1]) {
			posneg = true
		} else {
			posneg = false
		}

		for i, _ := range numbers {
			if len(numbers) < i+2 {
				continue
			}

			if posneg {
				if stringToInt(numbers[i])-stringToInt(numbers[i+1]) > 0 {
					continue outerLoop
				}
			} else {
				if stringToInt(numbers[i])-stringToInt(numbers[i+1]) < 0 {
					continue outerLoop
				}
			}

			difference = int(math.Abs(float64(stringToInt(numbers[i]) - stringToInt(numbers[i+1]))))
			if difference > 3 || difference < 1 {
				continue outerLoop
			}
		}
		safeLineCount++
	}

	fmt.Println(safeLineCount)
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		log.Fatal(err)
	}

	return i
}
