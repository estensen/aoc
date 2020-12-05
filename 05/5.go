package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("5.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	maxID := 0
	bitmap := [819]bool{} // found in a): 818+1

	for scanner.Scan() {
		line := scanner.Text()

		minRow := 0
		maxRow := 127
		minCol := 0
		maxCol := 7

		for _, c := range line {
			switch c {
			case 'F':
				maxRow -= (maxRow - minRow + 1) / 2
			case 'B':
				minRow += (maxRow - minRow + 1) / 2
			case 'L':
				maxCol -= (maxCol - minCol + 1) / 2
			case 'R':
				minCol += (maxCol - minCol + 1) / 2
			}
		}
		id := minRow*8 + minCol
		bitmap[id] = true
		if id > maxID {
			maxID = id
		}
	}
	fmt.Println(maxID)

	prev := false
	for i, b := range bitmap {
		if !b && prev {
			println(i)
		}
		prev = b
	}
}
