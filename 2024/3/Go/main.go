package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var total int
	text := string(content)

	// Regular expression to match mul(<number>,<number>)
	re := regexp.MustCompile(`mul\s*\((\d+)\s*,\s*(\d+)\)`)
	// re := regexp.MustCompile("mul\\((\\d{1,3}),(\\d{1,3})\\)")

	matches3 := re.FindAllStringSubmatch(text, -1)

	for _, match := range matches3 {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		total += num1 * num2
	}

	fmt.Println(total)

}
