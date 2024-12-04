package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

	var characters [][]rune
	//X = 88
	//M = 77
	//A = 65
	//S = 83

	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		characters = append(characters, runes)
	}

	fmt.Print(scanForX(characters))
}

func scanForX(characters [][]int32) int {
	var xmasCount int
	for i, e := range characters {
		for i2, e2 := range e {
			if e2 == 88 {
				xmasCount += scanForXmas(characters, i, i2)
			}
		}
	}

	return xmasCount
}

func scanForXmas(characters [][]int32, row int, col int) (xmasCount int) {

	//horizontal, right
	if len(characters[row]) > 3+col {
		if characters[row][col+1] == 77 {
			if characters[row][col+2] == 65 {
				if characters[row][col+3] == 83 {
					xmasCount++
					fmt.Println("HR")
				}
			}
		}
	}

	//horizontal, left
	if col > 2 {
		if characters[row][col-1] == 77 {
			if characters[row][col-2] == 65 {
				if characters[row][col-3] == 83 {
					xmasCount++
					fmt.Println("HL")
				}
			}
		}
	}

	//vertical, up
	if row > 2 {
		if characters[row-1][col] == 77 {
			if characters[row-2][col] == 65 {
				if characters[row-3][col] == 83 {
					xmasCount++
					fmt.Println("VU")
				}
			}
		}
	}

	//vertical, down
	if len(characters) > row+3 {
		if characters[row+1][col] == 77 {
			if characters[row+2][col] == 65 {
				if characters[row+3][col] == 83 {
					xmasCount++
					fmt.Println("VD")
				}
			}
		}
	}

	//diagonal, down-right
	if len(characters[row]) > 3+col && len(characters) > row+3 {
		if characters[row+1][col+1] == 77 {
			if characters[row+2][col+2] == 65 {
				if characters[row+3][col+3] == 83 {
					xmasCount++
					fmt.Println("DDR")
				}
			}
		}
	}

	//diagonal, down-left (not working)
	if len(characters) > row+3 && col > 2 {
		if characters[row+1][col-1] == 77 {
			if characters[row+2][col-2] == 65 {
				if characters[row+3][col-3] == 83 {
					xmasCount++
					fmt.Println("DDL")
				}
			}
		}
	}

	//diagonal, up-left
	if col > 2 && row > 2 {
		if characters[row-1][col-1] == 77 {
			if characters[row-2][col-2] == 65 {
				if characters[row-3][col-3] == 83 {
					xmasCount++
					fmt.Println("DUL")
				}
			}
		}
	}

	//diagonal, up-right
	if len(characters[row]) > 3+col && row > 2 {
		if characters[row-1][col+1] == 77 {
			if characters[row-2][col+2] == 65 {
				if characters[row-3][col+3] == 83 {
					xmasCount++
					fmt.Println("DUR")
				}
			}
		}
	}

	return xmasCount
}
