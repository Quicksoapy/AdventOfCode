package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
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

	// list instances of 'don't' and 'do', rip out everything inbetween, then send it to the regex?
	// don't(.*?)do(?!n't) does not work in go because lookbehind (!?) isn't supported in go wehhh
	cmd := exec.Command("python3", "regex.py", text)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}

	// Regular expression to match mul(<number>,<number>)
	re := regexp.MustCompile(`mul\s*\((\d+)\s*,\s*(\d+)\)`)

	matches3 := re.FindAllStringSubmatch(string(output), -1)

	for _, match := range matches3 {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		total += num1 * num2
	}

	fmt.Println(total)
}
