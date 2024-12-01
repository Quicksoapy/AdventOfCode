package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	list1, list2 := SortNumbersInTwoLists(content)

	sort.Ints(list1[:])
	sort.Ints(list2[:])

	sum := sumOfAllDistances(list1, list2)

	fmt.Println("Sum of all differences: " + strconv.Itoa(sum))

	var totalSimilarityScore int

	for _, element := range list1 {
		simScore := getSimilarityScoreOfNumber(element, list2)
		totalSimilarityScore += simScore
	}

	fmt.Println("Total similarity score: " + strconv.Itoa(totalSimilarityScore))
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		log.Fatal(err)
	}

	return i
}

func SortNumbersInTwoLists(b []byte) (l1 []int, l2 []int) {
	var list1 []int
	var list2 []int

	var numbers []string

	scanner := bufio.NewScanner(strings.NewReader(string(b)))
	for scanner.Scan() {
		numbers = strings.Fields(scanner.Text())

		list1 = append(list1, stringToInt(numbers[0]))
		list2 = append(list2, stringToInt(numbers[1]))
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("error occurred: %v\n", err)
	}

	return list1, list2
}

func sumOfAllDistances(l1 []int, l2 []int) int {
	var distance int
	var sum int
	for i := range l1 {
		distance = l1[i] - l2[i]
		sum += int(math.Abs(float64(distance)))
	}

	return sum
}

func getSimilarityScoreOfNumber(number int, list []int) int {
	var count int
	for _, e := range list {
		if e == number {
			count++
		}
	}

	return count * number
}
