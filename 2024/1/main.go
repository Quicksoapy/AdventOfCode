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

	list1, list2 := sortNumbersInTwoLists(content)

	sort.Ints(list1[:])
	sort.Ints(list2[:])

	sum := sumOfAllDistances(list1, list2)

	fmt.Println(sum)
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		// ... handle error
		log.Fatal(err)
	}

	return i
}

func sortNumbersInTwoLists(b []byte) (l1 []int, l2 []int) {
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
