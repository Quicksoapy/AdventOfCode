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

	// var lines [][]string
	var numbers []string
	//true = going up, false = going down
	var posneg bool
	var difference int
	var safeLineCount int

	scanner := bufio.NewScanner(strings.NewReader(string(content)))

	file, err := os.Create("output.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

scannerLoop:
	for scanner.Scan() {
		// The levels are either all increasing or all decreasing.
		// Any two adjacent levels differ by at least one and at most three.
		numbers = strings.Fields(scanner.Text())

	numberLoop:
		for index := range numbers {

			numbers2 := make([]string, len(numbers))
			copy(numbers2, numbers)
			numbers2 = remove(numbers2, index)

			if stringToInt(numbers2[0]) < stringToInt(numbers2[1]) {
				posneg = true
			} else {
				posneg = false
			}

			for a := range numbers2 {
				if len(numbers2) < a+2 {
					continue
				}

				if posneg {
					if stringToInt(numbers2[a])-stringToInt(numbers2[a+1]) > 0 {
						continue numberLoop
					}
				} else {
					if stringToInt(numbers2[a])-stringToInt(numbers2[a+1]) < 0 {
						continue numberLoop
					}
				}

				difference = int(math.Abs(float64(stringToInt(numbers2[a]) - stringToInt(numbers2[a+1]))))
				if difference > 3 || difference < 1 {
					continue numberLoop
				}
			}

			safeLineCount++
			continue scannerLoop
		}
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

func remove(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}
